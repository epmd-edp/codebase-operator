Codebase is the representation of any repository that holds code and can be maintained with
CI/CD pipelines.

The main purpose of codebase controller is to watch changes in the respective Kubernetes
 Custom resource (Codebase CR) and to ensure that state in that resource is applied in EDP.
 
There are several steps that are performed in the reconcile loop:

![arch](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/epmd-edp/codebase-operator/master/documentation/puml/codebase_chain.puml)

One of the field in Codebase CR is strategy. It can be import, create and clone. The main difference is:
- clone and create strategy require creation of the new repository in VCS, that uses for particular installation
of EDP. The source code is pulling from the repository, that was specified in Codebase CR (clone strategy), or from the
one of predefined ones (create strategy) 
- import strategy uses the existing repository and integrate it in EDP.

That's why you can see 2 different branches in the diagram.

Import consists of the following steps:
- clone git repository. The existence of the repository from Codebase CR is checked and repository is pulled
in the temporary workspace.
- ensure deploy config in git. Instructions how to deploy this codebase in Kubernetes are added at this step.
Represented as Helm charts
- ensure s2i image stream. Deprecated step. Should be deleted in https://jiraeu.epam.com/browse/EPMDEDP-4116
- ensure Jenkins Folder CR. Custom resource for Jenkins folder to hold ci/cd pipelines related to this codebase
is added
- cleaner. Technical step. Ensures that all workspaces are wiped out.

Clone and create branch includes the following steps:
- ensure gerrit project. Ensures that the corresponding Gerrit project is created for this codebase. Clone and push
the source code from the specified repository is performed at this step.
- ensure gerrit replication. Configuration of the replication of newly created Gerrit project is set up at this step.
Enabled if field vcs_integration_enabled in the config map edp-config is set to true.
- ensure Perf integration. Deprecated step.
- ensure deploy config in git. Instructions how to deploy this codebase in Kubernetes are added at this step.
- ensure s2i image stream. Deprecated step. Should be deleted in https://jiraeu.epam.com/browse/EPMDEDP-4116
- ensure Jenkins Folder CR. Custom resource for Jenkins folder to hold ci/cd pipelines related to this codebase
is added
- cleaner. Technical step. Ensures that all workspaces are wiped out.