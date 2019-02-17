<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" modal-title="新增/编辑 work" :modal-width="600">
    <!-- 表单信息 -->
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
      <FormItem label="work_name" prop="work_name">
        <Input v-model.trim="formValidate.work_name" placeholder="请输入 work_name"></Input>
      </FormItem>
      <FormItem label="work_desc" prop="work_desc">
        <Input v-model.trim="formValidate.work_desc" type="textarea" :rows="4" placeholder="请输入 work_desc"></Input>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
      </FormItem>
    </Form>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {EditWork} from "../../api"
  import {validateCommonPatternForString} from "../../tools"

  export default {
    name: "WorkEdit",
    components:{ISimpleBtnTriggerModal},
    data(){
      const _validateWorkName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('字段值不能为空!'));
        } else if (!validateCommonPatternForString(value)) {
          callback(new Error('存在非法字符，只能包含字母，数字，下划线!'));
        } else {
          callback();
        }
      };
      return {
        formValidate: {
          work_id:-1,
          work_name: '',
          work_desc: '',
        },
        ruleValidate: {
          work_name: [
            { validator: _validateWorkName, trigger: 'blur' },
          ],
          work_desc: [
            { required: true, message: 'work_desc 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      triggerWorkEdit(work){
        // 回显内容
        this.formValidate.work_id = work.id;
        this.formValidate.work_name = work.work_name;
        this.formValidate.work_desc = work.work_desc;
        // 显示模态对话框
        this.$refs.triggerModal.triggerClick();
      },
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditWork(this.formValidate.work_id, this.formValidate.work_name, this.formValidate.work_desc);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 调用子组件隐藏 modal (this.refs.xxx.子组件定义的方法())
              this.$refs.triggerModal.hideModal();
              // 通知父组件添加成功
              this.$emit('handleSuccess');
              // 表单重置,以取消缓存
              this.$refs[name].resetFields();
              this.formValidate.work_id = -1;
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    }
  }
</script>

<style scoped>

</style>
