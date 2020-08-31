package chain

import (
	"fmt"
	"github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/codebase-operator/v2/pkg/controller/gitserver"
	"github.com/epmd-edp/codebase-operator/v2/pkg/util"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CleanTempDirectory struct {
	client client.Client
	git    gitserver.Git
}

func (h CleanTempDirectory) ServeRequest(cb *v1alpha1.CodebaseBranch) error {
	rl := log.WithValues("namespace", cb.Namespace, "codebase branch", cb.Name)
	rl.Info("start CleanTempDirectory method...")

	c, err := util.GetCodebase(h.client, cb.Spec.CodebaseName, cb.Namespace)
	if err != nil {
		return err
	}

	wd := fmt.Sprintf("/home/codebase-operator/edp/%v/%v/%v/%v", c.Namespace, c.Name, "templates", c.Name)
	if err := deleteWorkDirectory(wd); err != nil {
		return err
	}

	rl.Info("end CleanTempDirectory method...")
	return nil
}

func deleteWorkDirectory(dir string) error {
	if err := util.RemoveDirectory(dir); err != nil {
		return errors.Wrapf(err, "couldn't delete directory %v", dir)
	}
	log.Info("directory was cleaned", "path", dir)
	return nil
}
