package iworkdata

import "encoding/xml"

type ParamResolver struct {
	ParamStr string
} 

func (this *ParamResolver) ParseParamStrToMap() *map[string]interface{}{
	return &map[string]interface{}{}
}

type ParamInputSchema struct {
	XMLName     		xml.Name 				`xml:"paramInputSchema" json:"-"`
	ParamInputSchemaItems []ParamInputSchemaItem			`xml:"paramInputSchemaItem"`
}

func (this *ParamInputSchema) RenderToXml() string {
	if bytes, err := xml.MarshalIndent(this,"", "\t"); err == nil{
		return string(bytes)
	}
	return ""
}

type ParamInputSchemaItem struct {
	XMLName    	xml.Name 	`xml:"paramInputSchemaItem" json:"-"`
	ParamName 	string		`xml:"paramName"`
	ParamValue 	string		`xml:"paramValue"`
}