# Jira Fix Version Controller

**Jira Fix Version** is the representation of Fix Version in Jira. It is used to automate routine work with
managing Fix Versions

The main purpose of a Jira Fix Version controller is to watch changes in the respective Kubernetes Custom Resource
 (JiraFixVersion CR) and to ensure that the state in that resource is applied in corresponding Jira Server.
 
Inspect the main steps performed in the reconcile loop on the diagram below:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/codebase-operator/master/documentation/puml/jira_fix_version_chain.puml)

The diagram above displays the following steps:

- *Check Parent Jira Server CR*. Each Jira Fix Version has a parent Jira Server. Firstly, the parent codebase is retrieved
by the `*spec.codebaseName"` field. Then the Jira Server CR is retrieved by `codebave.spec.jiraServer` field. If one of the
aforementioned components is not available the loop ends up with an error.
- *Ensure Fix Version*. The purpose of this step is to create Fix version in Jira Server by `metadata.name` field. Jira 
Project Id is get by the first item in `spec.tickets` field. The URL and required credentials is retrieved from the parent
Jira Server CR.
- *Ensure Jira tickets with fix version*. Jira fix version is added for all Jira tickets specified in `spec.tickets` field.
- *Delete Jira Fix Version CR*. After all successful previous steps, Controller should delete the respective Jira Fix Version
Custom Resource. By design, it is disposable and no changes are planned to be made in it. To avoid redundant calls to Jira
Server, this CR should be deleted.