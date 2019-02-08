<template>
  <span>
     <Tree :data="data1" show-checkbox></Tree>
  </span>
</template>

<script>
  export default {
    name: "WorkStepParamOutputDisplay",
    props:{
      paramOutputSchemaTreeNode:{
        type: Object,
        default: () => {},
      }
    },
    computed:{
      data1:function () {
        var appendChildrens = function (paramOutputSchemaTreeNode, node) {       // 父级节点对象、父级节点树元素
          if(paramOutputSchemaTreeNode.NodeChildrens != null && paramOutputSchemaTreeNode.NodeChildrens.length > 0){
            const arr = [];
            for(var i=0; i<paramOutputSchemaTreeNode.NodeChildrens.length; i++) {
              var childParamOutputSchemaTreeNode = paramOutputSchemaTreeNode.NodeChildrens[i];
              var childNode = {title: childParamOutputSchemaTreeNode.NodeName,expand: true,};
              // 递归操作
              appendChildrens(childParamOutputSchemaTreeNode, childNode);
              arr.push(childNode);
            }
            node.children = arr;
          }
        };
        const topTreeNode = {
          title: this.paramOutputSchemaTreeNode.NodeName,
          expand: true,
        };
        appendChildrens(this.paramOutputSchemaTreeNode, topTreeNode);
        return [
          topTreeNode
        ]
      }
    }
  }
</script>

<style scoped>

</style>
