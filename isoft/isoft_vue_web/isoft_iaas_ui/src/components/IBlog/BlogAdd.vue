<template>
  <div>
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="文章标题" prop="blog_title">
        <Input v-model="formValidate.blog_title" placeholder="Enter blog title..."/>
      </FormItem>
      <FormItem label="检索词条" prop="key_words">
        <Input v-model="formValidate.key_words" placeholder="Enter key_words..."></Input>
      </FormItem>
      <FormItem label="文章分类" prop="catalog_id">
        <Select v-model="formValidate.catalog_id" filterable>
          <Option v-for="mycatalog in mycatalogs" :value="mycatalog.id" :key="mycatalog.id">
            {{ mycatalog.catalog_name }}
          </Option>
        </Select>
      </FormItem>
      <FormItem label="文章内容" prop="content">
        <mavon-editor v-model="formValidate.content" :ishljs = "true"/>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
        <Button style="margin-left: 8px" @click="handleReset('formValidate')">Cancel</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {GetMyCatalogs} from "../../api"
  import {BlogEdit} from "../../api"

  export default {
    name: "BlogAdd",
    data () {
      return {
        // 我的所有文章分类
        mycatalogs:[],
        formValidate: {
          blog_title: '',
          key_words: '',
          catalog_id: -1,
          content:"",
        },
        ruleValidate: {
          blog_title: [
            { required: true, message: '文章标题不能为空', trigger: 'blur' }
          ],
          key_words: [
            { required: true, message: '检索词条不能为空', trigger: 'blur' }
          ],
          catalog_id: [
            { required: true, type: 'number', message: '文章分类不能为空', trigger: 'blur' }
          ],
          content: [
            { required: true, message: '文章内容不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BlogEdit(_this.formValidate.blog_title,
              _this.formValidate.key_words, _this.formValidate.catalog_id, _this.formValidate.content);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              _this.$router.go(0);     // 页面刷新,等价于 location.reload()
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      }
    },
    mounted:async function () {
      const result = await GetMyCatalogs();
      if(result.status=="SUCCESS"){
        this.mycatalogs = result.catalogs;
      }
    }
  }
</script>

<style scoped>

</style>
