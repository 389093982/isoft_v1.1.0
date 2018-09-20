<template>
  <div>
    <IFileUpload/>

    <Table :columns="columns1" :data="metadatas" size="small" height="450"></Table>
  </div>
</template>

<script>
  import IFileUpload from "./IFileUpload.vue"
  import {MetaDataList} from '../../api'

  export default {
    name: "IFile",
    components: {IFileUpload},
    data(){
      return {
        metadatas: [],
        columns1 : [
          {
            title: 'Name',
            key: 'Name',
            width:300
          },
          {
            title: 'Version',
            key: 'Version',
            width:100
          },
          {
            title: 'Size',
            key: 'Size',
            width:100
          },
          {
            title: 'Hash',
            key: 'Hash',
          }
        ]
      }
    },
    methods:{
      async refreshMetaDataList(){
        const data = await MetaDataList();
        if(data.status == "SUCCESS"){
          this.metadatas = data.metadatas;
        }
      }
    },
    mounted:function(){
      this.refreshMetaDataList();
    },
  }
</script>

<style scoped>

</style>
