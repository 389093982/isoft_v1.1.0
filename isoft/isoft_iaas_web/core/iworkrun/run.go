package iworkrun

import (
	"database/sql"
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkcomponent"
	"isoft/isoft_iaas_web/core/iworkcomponent/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func Run(work iwork.Work, steps []iwork.WorkStep) {
	trackingId := stringutil.RandomUUID()
	// 插入 RunLogRecord
	iwork.InsertRunLogRecord(&iwork.RunLogRecord{
		TrackingId:trackingId,
		WorkName:work.WorkName,
		CreatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedBy:"SYSTEM",
		LastUpdatedTime:time.Now(),
	})

	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
	}()

	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute work:%s",work.WorkName))
	for _, step := range steps {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep:%s",step.WorkStepName))
		RunStep(work, step)
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep:%s",step.WorkStepName))
	}
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute work:%s",work.WorkName))
}

func RunStep(work iwork.Work, step iwork.WorkStep) {
	switch strings.ToUpper(step.WorkStepType) {
	case "SQL_QUERY":
		SQLQueryRun(work, step)
	}
}

func SQLQueryRun(work iwork.Work, step iwork.WorkStep) {
	db, err := sqlutil.GetConnForMysql("mysql", iworkcomponent.GetParamValue(step, "db_conn"))
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(iworkcomponent.GetParamValue(step, "sql"))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	colNames, _ := rows.Columns()
	for rows.Next() {
		colValues := make([]sql.RawBytes, len(colNames))
		scanArgs := make([]interface{}, len(colValues))
		for i := range colValues {
			scanArgs[i] = &colValues[i]
		}
		rows.Scan(scanArgs...)
		for _, colValue := range colValues {
			fmt.Print(colValue)
		}
	}
}
