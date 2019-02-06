package iworkrun

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
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

func GetParamValue(step iwork.WorkStep, paramName string) string {
	var paramInputSchema iworkdata.ParamInputSchema
	if err := xml.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err != nil{
		panic(err)
	}
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		if item.ParamName == paramName{
			// 非必须参数不得为空
			if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == ""{
				panic(errors.New(fmt.Sprint("it is a mast parameter for %s", item.ParamName)))
			}
			return item.ParamValue
		}
	}
	return ""
}

func SQLQueryRun(work iwork.Work, step iwork.WorkStep) {
	db, err := sqlutil.GetConnForMysql("mysql", sqlutil.GetDataSourceName(GetParamValue(step,"db_conn")))
	if err != nil{
		panic(err)
	}
	rows, err := db.Query(GetParamValue(step,"sql"))
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


