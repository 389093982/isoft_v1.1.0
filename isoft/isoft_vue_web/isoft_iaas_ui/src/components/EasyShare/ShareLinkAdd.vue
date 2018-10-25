<template>
  <span>
    <a href="javascript:;" @click="showShareLinkAddFlag=true">新增</a>

    <Modal
      v-model="showShareLinkAddFlag"
      width="500"
      title="我要分享"
      :footer-hide="true"
      :mask-closable="false">
        <!-- 表单正文 -->
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
          <FormItem label="分享类型" prop="share_type">
            <Input v-model="formValidate.share_type" placeholder="请输入分享类型"></Input>
          </FormItem>
          <FormItem label="分享链接" prop="link_href">
            <Input v-model="formValidate.link_href" placeholder="请输入分享链接"></Input>
          </FormItem>
          <FormItem>
            <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
            <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
          </FormItem>
        </Form>
    </Modal>
  </span>
</template>

<script>
  export default {
    name: "ShareLinkAdd",
    data(){
      return {
        showShareLinkAddFlag:false,
        formValidate: {
          share_type: '',
          link_href: '',
        },
        ruleValidate: {
          share_type: [
            { required: true, message: '分享类型不能为空', trigger: 'blur' }
          ],
          link_href: [
            { required: true, message: '分享链接不能为空', trigger: 'blur' }
          ],
        }
      }
    },
    methods: {
      handleSubmit (name) {
        let data = {
          share_type:this.formValidate.share_type,
          link_href:this.formValidate.link_href,
        };

        var _this = this;
        this.$refs[name].validate((valid) => {
          if (valid) {
            alert(111);
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      }
    }
  }
</script>

<style scoped>

</style>
