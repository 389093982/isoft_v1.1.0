<template>
  <span>
     <Tree :data="data1"></Tree>
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
        const result = {
          title: this.paramOutputSchemaTreeNode.NodeName,
          expand: true,
        };
        if(this.paramOutputSchemaTreeNode.NodeChildrens.length > 0){
          const arr = [];
          for(var i=0; i<this.paramOutputSchemaTreeNode.NodeChildrens.length; i++){
            arr.push({
              title: this.paramOutputSchemaTreeNode.NodeChildrens[i].NodeName,
            });
          }
          result.children = arr;
        }
        return [
          result
        ]
      }
    }
  }
</script>

<style scoped>

</style>
