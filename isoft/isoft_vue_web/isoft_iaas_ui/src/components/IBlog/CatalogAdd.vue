<template>
<span>
  <div>
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="分类名称" prop="catalog_name">
        <Row>
        <Col span="18">
          <Input v-model="formValidate.catalog_name" placeholder="Enter catalog name..."/>
        </Col>
        <Col span="6">
          <span style="margin-left: 5px;"><Icon type="md-book" />选择推荐分类</span>
        </Col>
      </Row>
      </FormItem>
      <FormItem label="分类简介" prop="catalog_desc">
        <Input v-model="formValidate.catalog_desc" placeholder="Enter catalog desc..."></Input>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
        <Button style="margin-left: 8px" @click="handleReset('formValidate')">Cancel</Button>
      </FormItem>
    </Form>
  </div>
</span>
</template>

<script>
  import {CatalogEdit} from "../../api"

  export default {
    name: "CatalogAdd",
    data () {
      return {
        formValidate: {
          catalog_name: '',
          catalog_desc: '',
        },
        ruleValidate: {
          catalog_name: [
            { required: true, message: '分类名称不能为空', trigger: 'blur' }
          ],
          catalog_desc: [
            { required: true, message: '分类描述不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await CatalogEdit(_this.formValidate.catalog_name, _this.formValidate.catalog_desc);
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
    }
  }
</script>

<style scoped>

</style>
