package iworkrun

import (
	"database/sql"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkcomponent/sqlutil"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

func Run(work iwork.Work, steps []iwork.WorkStep)  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Print(err)
		}
	}()
	for _, step := range steps{
		RunStep(work, step)
	}
}

func RunStep(work iwork.Work, step iwork.WorkStep)  {
	switch strings.ToUpper(step.WorkStepType) {
	case "SQL_QUERY":
		SQLQueryRun(work, step)
	}
}


func SQLQueryRun(work iwork.Work, step iwork.WorkStep) {
	db, err := sqlutil.GetConnForMysql("mysql", sqlutil.GetDataSourceName(iworkdata.GetParamValue(step,"db_conn")))
	if err != nil{
		panic(err)
	}
	rows, err := db.Query(iworkdata.GetParamValue(step,"sql"))
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	colNames,_ := rows.Columns()
	for rows.Next(){
		colValues := make([]sql.RawBytes, len(colNames))
		scanArgs := make([]interface{}, len(colValues))
		for i := range colValues {
			scanArgs[i] = &colValues[i]
		}
		rows.Scan(scanArgs...)
		for _, colValue := range colValues{
			fmt.Print(colValue)
		}
	}
}


