package iworkdata

import "encoding/xml"

type ParamInputSchemaItem struct {
	XMLName    	xml.Name 	`xml:"paramInputSchemaItem" json:"-"`
	ParamName 	string		`xml:"paramName"`
	ParamValue 	string		`xml:"paramValue"`
}

type ParamInputSchema struct {
	XMLName     		xml.Name 							`xml:"paramInputSchema" json:"-"`
	ParamInputSchemaItems []ParamInputSchemaItem			`xml:"paramInputSchemaItem"`
}

func (this *ParamInputSchema) RenderToXml() string {
	if bytes, err := xml.MarshalIndent(this,"", "\t"); err == nil{
		return string(bytes)
	}
	return ""
}

type ParamOutputSchemaItem struct {
	XMLName    	xml.Name 	`xml:"paramOutputSchemaItem" json:"-"`
	ParamName 	string		`xml:"paramName"`
	ParamValue 	string		`xml:"paramValue"`
}

type ParamOutputSchema struct {
	XMLName     		xml.Name 							`xml:"paramOutputSchema" json:"-"`
	ParamOutputSchemaItems []ParamOutputSchemaItem			`xml:"paramOutputSchemaItem"`
}

func (this *ParamOutputSchema) RenderToXml() string {
	if bytes, err := xml.MarshalIndent(this,"", "\t"); err == nil{
		return string(bytes)
	}
	return ""
}