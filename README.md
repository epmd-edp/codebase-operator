# Codebase operator

## Overview

Codebase operator is an EDP operator that is responsible for provisioning codebase entities. Operator installation can be applied on two container orchestration platforms: Openshift and Kubernetes.

_**NOTE:** Operator is platform-independent, that is why there is an unified instruction for deploying._

## Prerequisites
1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following one of the instructions: [edp-install-openshift](https://github.com/epmd-edp/edp-install/blob/master/documentation/openshift_install_edp.md#edp-project) or [edp-install-kubernetes](https://github.com/epmd-edp/edp-install/blob/master/documentation/kubernetes_install_edp.md#edp-namespace).

## Installation
In order to install the Codebase operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":
     ```bash
     helm repo add epamedp https://chartmuseum.demo.edp-epam.com/
     ```
2. Choose available Helm chart version:
     ```bash
     helm search repo epamedp/codebase-operator
     NAME                           CHART VERSION   APP VERSION     DESCRIPTION
     epamedp/codebase-operator      v2.4.0                          Helm chart for Golang application/service deplo...
     ```
  
    _**NOTE:** It is highly recommended to use the latest released version._

3. Deploy operator:

   Full available chart parameters list:
 ```
    - chart_version                                   # a version of CD Pipeline operator Helm chart;
    - global.edpName                                  # a namespace or a project name (in case of OpenShift);
    - global.platform                                 # openShift or kubernetes;
    - global.database.host                            # Host to DB (<db-name>.<namespace>);
    - global.database.port                            # Port to DB;
    - global.database.name                            # Name of DB;
    - image.name                                      # EDP image. The released image can be found on [Dockerhub](https://hub.docker.com/repository/docker/epamedp/codebase-operator);
    - image.version                                   # EDP tag. The released image can be found on [Dockerhub](https://hub.docker.com/repository/docker/epamedp/codebase-operator/tags);
    - jira.integration                                # Flag to enable/disable Jira integration;
    - jira.name                                       # JiraServer CR name;
    - jira.apiUrl                                     # API URL for development;
    - jira.rootUrl                                    # URL to Jira server;
    - jira.credentialName                             # Name of secret with credentials to Jira server;
 ```
 
Install operator in the <edp_cicd_project> namespace with helm command. Bellow is the installation command example:
```bash
helm install codebase-operator epamedp/codebase-operator --version <chart_version> --namespace <edp_cicd_project> --set name=codebase-operator --set global.edpName=<edp_cicd_project> --set global.platform=<platform_type> --set jira.integration=false
```

* Check the <edp_cicd_project> namespace that should contain operator deployment with your operator in a running status

### Related Articles

- [Codebase Controller Overview](documentation/codebase_controller.md)
- [Codebase Branch Controller](documentation/codebase_branch_controller.md)
- [Jira Server Controller](documentation/jira_server_controller.md)
- [Jira Fix Version Controller](documentation/jira_fix_version_controller.md)
- [Git Server Controller](documentation/git_server_controller.md)