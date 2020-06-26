# Jira Server Overview

**Jira Server** is the representation of Jira. It is used to communicate with Jira API to manage different entities
like tickets, fix version etx  

The main purpose of a Jira Server controller is to watch changes in the respective Kubernetes Custom Resource (Codebase CR)
 and to ensure that the state in that resource is applied in EDP.
 
Explore the main steps performed in the reconcile loop on the diagram below:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/codebase-operator/master/documentation/puml/jira_server_chain.puml)

There are several steps in the chain^

- *Ensure connection to Jira Server*. Using spec.ApiUrl and spec.credentialName controller tries to login to the 
specified URL. If connection is not successful loop ends with error. 
- *Put EDP Component*. Registration of new Component in EDP.
- *Update status*. Updates status in the respective Jira Server CR