package common

import (
	"github.com/JayneJacobs/FullStackWebDev/kickstart/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
