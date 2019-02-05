package iworkdata

import "encoding/xml"

type ParamResolver struct {
	ParamStr string
} 

func (this *ParamResolver) ParseParamStrToMap() *map[string]interface{}{
	return &map[string]interface{}{}
}

type ParamDefinition struct {
	XMLName     		xml.Name 				`xml:"paramDefinition" json:"-"`
	ParamDefinitionItems []ParamDefinitionItem	`xml:"paramDefinitionItem"`
}

func (this *ParamDefinition) RenderToXml() string {
	if bytes, err := xml.MarshalIndent(this,"", "\t"); err == nil{
		return string(bytes)
	}
	return ""
}

type ParamDefinitionItem struct {
	XMLName    	xml.Name 	`xml:"paramDefinitionItem" json:"-"`
	ParamName 	string		`xml:"paramName"`
	ParamValue 	string		`xml:"paramValue"`
}