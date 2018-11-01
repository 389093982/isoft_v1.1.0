<template>

  <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
    <Row>
      <Col span="16" style="padding:20px;border-right: 1px solid #e6e6e6;">
        <div>
          <!-- 表单正文 -->
          <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
            <FormItem label="分享类型" prop="share_type">
              <Row>
                <Col span="22">
                  <Input v-model="formValidate.share_type" placeholder="请输入分享类型"></Input>
                </Col>
                <Col span="2" style="text-align: right;">
                  <Poptip v-model="visible" placement="left-start" width="420">
                    <a href="javascript:;">热门分类</a>
                    <div slot="content">
                  <span v-for="type in hot_share_type" style="margin: 5px;float: left;">
                    <Tag><a @click="closePoptip(type.name)">{{type.name}}</a></Tag>
                  </span>
                    </div>
                  </Poptip>
                </Col>
              </Row>
            </FormItem>
            <FormItem label="简短描述" prop="share_desc">
              <Input v-model="formValidate.share_desc" placeholder="请输入简短描述"></Input>
            </FormItem>
            <FormItem label="文章内容" prop="content">
              <mavon-editor v-model="formValidate.content" :toolbars="toolbars" :ishljs = "true" style="z-index: 1;"/>
            </FormItem>
            <FormItem label="分享链接" prop="link_href">
              <Input v-model="formValidate.link_href" placeholder="请输入分享链接"></Input>
            </FormItem>
            <FormItem>
              <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
              <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
            </FormItem>
          </Form>
        </div>
      </Col>
      <Col span="8" style="padding: 20px;">

      </Col>
    </Row>
  </div>
</template>

<script>
  import {AddNewShareLink} from "../../api"

  export default {
    name: "ShareLinkAdd",
    data(){
      return {
        toolbars: {
          bold: true, // 粗体
          italic: true, // 斜体
          header: true, // 标题
          underline: true, // 下划线
          // mark: true, // 标记
          superscript: true, // 上角标
          quote: true, // 引用
          ol: true, // 有序列表
          link: true, // 链接
          imagelink: true, // 图片链接
          help: true, // 帮助
          code: true, // code
          subfield: true, // 是否需要分栏
          fullscreen: true, // 全屏编辑
          readmodel: true, // 沉浸式阅读
          undo: true, // 上一步
          trash: true, // 清空
          save: true, // 保存（触发events中的save事件）
          navigation: true // 导航目录
        },
        visible:false,
        hot_share_type: this.GLOBAL.hot_share_type,
        formValidate: {
          share_type: '',
          share_desc: '',
          link_href: '',
          content:"",
        },
        ruleValidate: {
          share_type: [
            { required: true, message: '分享类型不能为空', trigger: 'change' }
          ],
          share_desc: [
            { required: true, message: '简短描述不能为空', trigger: 'blur' }
          ],
          content: [
            { required: true, message: '文章内容不能为空' }
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
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddNewShareLink(_this.formValidate.share_type, _this.formValidate.share_desc, _this.formValidate.link_href,_this.formValidate.content);
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
