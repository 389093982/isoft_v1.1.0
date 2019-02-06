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

// 输出参数转换成 TreeNode 用于树形结构展示
type TreeNode struct {
	NodeName string
	NodeLink string
	NodeChildrens []*TreeNode
} 

func (this *ParamOutputSchema) RenderToTreeNodes() *TreeNode {
	topTreeNode := &TreeNode{
		NodeName:"$NODE_NAME_OUTPUT",
		NodeLink:"$NODE_NAME_OUTPUT",
	}
	for _, item := range this.ParamOutputSchemaItems{
		topTreeNode.NodeChildrens = append(topTreeNode.NodeChildrens, &TreeNode{
			NodeName:item.ParamName,
			NodeLink:item.ParamName,
		})
	}
	return topTreeNode
}