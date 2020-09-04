package controller

import (
	"github.com/epmd-edp/codebase-operator/v2/pkg/controller/gittag"
)

func init() {
	/*AddToManagerFuncs = append(AddToManagerFuncs, codebase.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, codebasebranch.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, gitserver.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, jirafixversion.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, jiraserver.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, imagestreamtag.Add)*/
	AddToManagerFuncs = append(AddToManagerFuncs, gittag.Add)
}
