package iworkdata

import "encoding/xml"

type ParamResolver struct {
	ParamStr string
} 

func (this *ParamResolver) ParseParamStrToMap() *map[string]interface{}{
	return &map[string]interface{}{}
}

type ParamSchema struct {
	XMLName     		xml.Name 				`xml:"paramSchema" json:"-"`
	ParamSchemaItems []ParamSchemaItem	`xml:"paramSchemaItem"`
}

func (this *ParamSchema) RenderToXml() string {
	if bytes, err := xml.MarshalIndent(this,"", "\t"); err == nil{
		return string(bytes)
	}
	return ""
}

type ParamSchemaItem struct {
	XMLName    	xml.Name 	`xml:"paramSchemaItem" json:"-"`
	ParamName 	string		`xml:"paramName"`
	ParamValue 	string		`xml:"paramValue"`
}