package iworknode

import (
	"isoft/isoft_iaas_web/models/iwork"
)

type EmptyNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EmptyNode) Execute(trackingId string) {

}
