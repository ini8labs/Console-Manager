package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/console-manager/src/console"
	"github.com/ini8labs/console-manager/src/middlewares"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"

	//"github.com/google/martian/log"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

type Server struct {
	console.Console
	logger *logrus.Logger
}

func initServer(logger *logrus.Logger) Server {
	config, err := rest.InClusterConfig()
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()

	// use the current context in kubeconfig
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	s := Server{
		Console: console.Console{DynamicClient: clientSet,
			Logger: logger,
		},
		logger: logger,
	}

	return s
}

func NewServer(addr string, logger *logrus.Logger) error {

	s := initServer(logger)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middlewares.LoggingMiddleware())

	r.GET("/createconsole/k8s", s.createEKSConsole)
	r.GET("/getconsole/k8s", s.getEKSConsole)
	r.DELETE("/deleteconsole/k8s", s.deleteEKSConsole)

	logger.Infof("Starting Serve on Port %s", addr)

	return r.Run(addr)
}

func (s Server) createEKSConsole(c *gin.Context) {
	appName := c.Query("app")
	namespace := c.Query("ns")
	labels := c.Query("labels")

	obj := console.BuildEKSConsoleObj(appName, labels)
	if err := s.Create(context.TODO(), namespace, obj); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		s.Errorf(err.Error())
		return
	}

	c.JSON(http.StatusOK, "object created Successfully")
}

func (s Server) getEKSConsole(c *gin.Context) {
	appName := c.Query("app")
	namespace := c.Query("ns")

	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      appName,
	}

	if err := s.Get(context.TODO(), namespacedName); err != nil {
		s.Errorf(err.Error())
		c.JSON(http.StatusInternalServerError, "Some internal error")
		return
	}

	c.JSON(http.StatusOK, "Object Retrieved Successfully")
}

func (s Server) deleteEKSConsole(c *gin.Context) {
	appName := c.Query("app")
	namespace := c.Query("ns")

	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      appName,
	}

	if err := s.Delete(context.TODO(), namespacedName); err != nil {
		s.Errorf(err.Error())
		c.JSON(http.StatusInternalServerError, "Some internal error")
		return
	}
	c.Handler()
	c.JSON(http.StatusAccepted, "Object Deleted Successfully")
}
