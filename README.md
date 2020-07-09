# How to Install Operator

EDP installation can be applied on two container orchestration platforms: OpenShift and Kubernetes.

_**NOTE:** Installation of operators is platform-independent, that is why there is a unified instruction for deploying._

## Prerequisites
1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following one of the instructions: [edp-install-openshift](https://github.com/epmd-edp/edp-install/blob/master/documentation/openshift_install_edp.md#edp-project) or [edp-install-kubernetes](https://github.com/epmd-edp/edp-install/blob/master/documentation/kubernetes_install_edp.md#edp-namespace).

## Installation
In order to install the EDP codebase-operator, follow the steps below:

1. To add the Helm EPAMEDP Charts for local client, run "helm repo add":
     ```bash
     helm repo add epamedp https://chartmuseum.demo.edp-epam.com/
     ```
2. Choose available Helm chart version:
    ```bash
     helm search repo epamedp/codebase-operator
    ```
   Example response:
   ```
   NAME                            CHART VERSION   APP VERSION     DESCRIPTION
   epamedp/codebase-operator       v2.4.0                          Helm chart for Golang application/service deplo...
   ```

    _**NOTE:** It is highly recommended to use the latest released version._
3. Deploy operator:
    Full available chart parameters list:
    ```
        - <chart_version>                        # Helm chart version;
        - global.edpName                         # a namespace or a project name (in case of OpenShift);
        - global.platform                        # a platform type that can be "kubernetes" or "openshift";
        - name                                   # component name;
        - database.required                      # database deployment request can be "true" or "false";
        - database.host                          # database host;
        - database.name                          # database name;
        - database.port                          # database port;
        - image.name                             # EDP codebase-operator Docker image name. The released image can be found on [Dockerhub](https://hub.docker.com/repository/docker/epamedp/codebase-operator);
        - image.version                          # EDP codebase-oprator Docker image tag. The released image can be found on [Dockerhub](https://hub.docker.com/repository/docker/epamedp/codebase-operator/tags);
        - jira.integration                       # enable jira integration, can be "true" or "false";
        - jira.name                              # jira server name
        - jira.apiUrl                            # link to jira api
        - jira.rootUrl                           # jira URL
        - jira.credentialName                    # name of credentials for jira integration
    ```
    Set your parameters and launching a Helm chart deployment. Example command:
    ```bash
    helm install codebase-operator epamedp/codebase-operator --version <chart_version> --namespace <edp_cicd_project> --set name=codebase-operator --set global.edpName=<edp_cicd_project> --set global.platform=<platform_type> --set image.name=epamedp/codebase-operator --set image.version=<operator_version>
    ```
    
4. Check the <edp_cicd_project> namespace that should contain Deployment with your operator in a running status.

### Related Articles

- [Codebase Controller Overview](documentation/codebase_controller.md)
- [Codebase Branch Controller](documentation/codebase_branch_controller.md)
- [Jira Server Controller](documentation/jira_server_controller.md)
- [Jira Fix Version Controller](documentation/jira_fix_version_controller.md)
- [Git Server Controller](documentation/git_server_controller.md)