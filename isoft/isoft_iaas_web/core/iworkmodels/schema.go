package iworkmodels

import (
	"encoding/json"
	"encoding/xml"
)

type ParamInputSchemaItem struct {
	XMLName      xml.Name `xml:"paramInputSchemaItem" json:"-"`
	ParamName    string   `xml:"paramName"`
	ParamDesc    string   `xml:"paramDesc"` // 使用说明信息
	Repeatable   bool     `xml:"repeatable"`
	RepeatRefer  string   `xml:"repeatRefer"`
	ParamChoices []string `xml:"paramChoices"`
	PureText     bool     `xml:"pureText"`
	ParamValue   string   `xml:"paramValue"`
}

type ParamInputSchema struct {
	XMLName               xml.Name               `xml:"paramInputSchema" json:"-"`
	ParamInputSchemaItems []ParamInputSchemaItem `xml:"paramInputSchemaItem"`
}

func (this *ParamInputSchema) RenderToJson() string {
	if bytes, err := json.MarshalIndent(this, "", "\t"); err == nil {
		return string(bytes)
	}
	return ""
}

type ParamOutputSchemaItem struct {
	XMLName    xml.Name `xml:"paramOutputSchemaItem" json:"-"`
	ParentPath string   `xml:"parentPath"`
	ParamName  string   `xml:"paramName"`
	ParamValue string   `xml:"paramValue"`
}

type ParamOutputSchema struct {
	XMLName                xml.Name                `xml:"paramOutputSchema" json:"-"`
	ParamOutputSchemaItems []ParamOutputSchemaItem `xml:"paramOutputSchemaItem"`
}

func (this *ParamOutputSchema) RenderToJson() string {
	if bytes, err := json.MarshalIndent(this, "", "\t"); err == nil {
		return string(bytes)
	}
	return ""
}

// 输出参数转换成 TreeNode 用于树形结构展示
type TreeNode struct {
	NodeName      string
	NodeLink      string
	NodeChildrens []*TreeNode
}

func (this *ParamOutputSchema) RenderToTreeNodes(rootName string) *TreeNode {
	// 渲染顶级节点
	topTreeNode := &TreeNode{
		NodeName: rootName,
		NodeLink: rootName,
	}
	for _, item := range this.ParamOutputSchemaItems {
		this.appendToTreeNodes(topTreeNode, item)
	}
	return topTreeNode
}

// 元素追加到树上面
func (this *ParamOutputSchema) appendToTreeNodes(treeNode *TreeNode, item ParamOutputSchemaItem) {
	pTreeNode := this.createAndGetParentTreeNode(treeNode, item)
	pTreeNode.NodeChildrens = append(pTreeNode.NodeChildrens, &TreeNode{NodeName: item.ParamName, NodeLink: item.ParamName})
}

func (this *ParamOutputSchema) createAndGetParentTreeNode(treeNode *TreeNode, item ParamOutputSchemaItem) *TreeNode {
	// 父级节点是根节点
	if item.ParentPath == "" {
		return treeNode
	}
	// 父级节点不是根节点
	for _, children := range treeNode.NodeChildrens {
		if children.NodeName == item.ParentPath {
			return children
		}
	}
	// 父级节点未曾创建过则重新创建
	pTreeNode := &TreeNode{NodeName: item.ParentPath, NodeLink: item.ParentPath}
	treeNode.NodeChildrens = append(treeNode.NodeChildrens, pTreeNode)
	return pTreeNode
}
