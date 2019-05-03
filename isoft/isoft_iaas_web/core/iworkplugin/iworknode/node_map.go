package iworknode

import (
	"isoft/isoft_iaas_web/models/iwork"
)

type MapNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}
