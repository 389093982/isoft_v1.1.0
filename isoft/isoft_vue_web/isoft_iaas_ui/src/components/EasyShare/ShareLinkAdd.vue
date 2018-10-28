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
            <Row>
              <Col span="20">
                <Input v-model="formValidate.share_type" placeholder="请输入分享类型"></Input>
              </Col>
              <Col span="4" style="text-align: right;">
                <Poptip v-model="visible" placement="left-start" width="420">
                  <a href="javascript:;">热门分类</a>
                  <div slot="content">
                    <span v-for="type in hot_share_type" style="margin: 5px;float: left;">
                      <Button @click="closePoptip(type)">{{type}}</Button>
                    </span>
                  </div>
                </Poptip>
              </Col>
            </Row>
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
  import {AddNewShareLink} from "../../api"

  export default {
    name: "ShareLinkAdd",
    data(){
      return {
        visible:false,
        showShareLinkAddFlag:false,
        hot_share_type:["Java","Python","Golang","Java","Python","Golang","Java","Python","Golang"],
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
      closePoptip (type) {
        this.formValidate.share_type=type;
        this.visible = false;
      },
      handleSubmit (name) {
        let data = {
          share_type:this.formValidate.share_type,
          link_href:this.formValidate.link_href,
        };

        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddNewShareLink(_this.formValidate.share_type, _this.formValidate.link_href);
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
