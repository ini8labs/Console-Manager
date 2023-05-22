[![Go Report Card](https://goreportcard.com/badge/github.com/ini8labs/console-manager)](https://goreportcard.com/report/github.com/ini8labs/console-manager)
# Console-Manager

- Console Manager is a Kubernetes controller that manages the CRD `EksConsoleShell`

- Below are the list of the operations that can be performed va this controller:
  - Create object EksConsoleShell
  - Read object EksConsoleShell
  - Delete object EksConsoleShell

### Create EksConsoleShell

- **Endpoint** : GET `/createconsole/k8s?app=<app-name>&ns=<namespace>&labels=<labels>`
  - app-name: name of the pod that will be created by this API.
  - namespace: Namespace where you want to create the app.
  - labels: labels in format `key=value` to be applied on the pod.

- Response of above request is 200 status Code with message: `object created Successfully`


### Get EksConsoleShell

- **Endpoint** : GET `/getconsole/k8s?app=<app-name>&ns=<namespace>`
    - app-name: name of the App to be fetched.
    - namespace: Namespace where the app is created.

- Response: 200 Status code with message `object Retrieved Successfully`

### Delete EksConsoleShell

- **Endpoint** : DELETE `/deleteconsole/k8s?app=<app-name>&ns=<namespace>`
  - app-name: name of the App to be deleted.
  - namespace: Namespace where the app should be deleted from.

- Response: 202 Status code with message `Object Deleted Successfully`
