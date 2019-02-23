package task

import (
	"isoft/isoft_iaas_web/imodules"
)

func RegisterCronTask()  {
	if imodules.CheckModule("ilearning"){
		startILearningCronTask()
	}

	if imodules.CheckModule("iwork"){
		startIWorkCronTask()
	}
}



