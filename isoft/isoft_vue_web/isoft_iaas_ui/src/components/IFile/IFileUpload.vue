<template>
<div>
  <Button type="success" @click="fileUploadModal = true">文件上传</Button>

  <Modal
    v-model="fileUploadModal"
    width="500"
    title="文件上传"
    :mask-closable="false">
    <div>
      <Upload
        ref="upload"
        multiple
        :on-success="uploadComplete"
        action="/api/ifile/fileUpload/">
        <Button icon="ios-cloud-upload-outline">文件上传</Button>
      </Upload>
    </div>
  </Modal>
</div>
</template>

<script>
    export default {
      name: "IFileUpload",
      data () {
        return {
          // 文件上传 modal
          fileUploadModal: false,
        }
      },
      methods:{
        uploadComplete(res, file) {
          if(res.status=="SUCCESS"){
            // 父子组件通信
            this.$emit('refreshTable');
            this.$Notice.success({
              title: '文件上传成功',
              desc: '文件 ' + file.name + ' 上传成功!'
            });
          }else{
            this.$Notice.error({
              title: '文件上传失败',
              desc: '文件 ' + file.name + ' 上传失败!'
            });
          }
        },
      }
    }
</script>

<style scoped>

</style>
