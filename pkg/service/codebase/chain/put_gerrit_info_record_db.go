package chain

import (
	"database/sql"
	"github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/codebase-operator/v2/pkg/openshift"
	"github.com/epmd-edp/codebase-operator/v2/pkg/service/codebase/chain/handler"
)

type PutGerritInfoRecordDB struct {
	next      handler.CodebaseHandler
	clientSet openshift.ClientSet
	DB        sql.DB
}

func (h PutGerritInfoRecordDB) ServeRequest(c *v1alpha1.Codebase) error {
	rLog := log.WithValues("codebase name", c.Name)
	rLog.V(2).Info("start updating codebase db...")
	/*if err := h.tryToUpdateCodebaseTable(*c); err != nil {
		return err
	}*/
	rLog.Info("codebase DB has been updated")
	return nextServeOrNil(h.next, c)
}

func (h PutGerritInfoRecordDB) tryToUpdateCodebaseTable(name string) error {
	//h.DB.Query()
	return nil
}
