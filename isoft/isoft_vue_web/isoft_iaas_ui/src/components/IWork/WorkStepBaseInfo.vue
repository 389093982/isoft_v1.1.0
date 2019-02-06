<template>
  <Modal
    v-model="showFormModal"
    width="800"
    title="查看/编辑 workstep"
    :footer-hide="true"
    :mask-closable="false">
    <div>
      <!-- 表单信息 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
        <FormItem label="work_id" prop="work_id">
          <Input v-model.trim="formValidate.work_id" readonly placeholder="请输入 work_id"></Input>
        </FormItem>
        <FormItem label="work_step_id" prop="work_step_id">
          <Input v-model.trim="formValidate.work_step_id" readonly placeholder="请输入 work_step_id"></Input>
        </FormItem>
        <FormItem label="work_step_name" prop="work_step_name">
          <Input v-model.trim="formValidate.work_step_name" placeholder="请输入 work_step_name"></Input>
        </FormItem>
        <FormItem label="work_step_type" prop="work_step_type">
          <Select v-model="formValidate.work_step_type" placeholder="请选择 work_step_type">
            <Option :value="default_work_step_type" v-for="default_work_step_type in default_work_step_types">{{default_work_step_type}}</Option>
          </Select>
        </FormItem>
        <FormItem>
          <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
          <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
        </FormItem>
      </Form>
    </div>
  </Modal>
</template>

<script>
  import {EditWorkStepBaseInfo} from "../../api"
  import {LoadWorkStepInfo} from "../../api"

  export default {
    name: "WorkStepBaseInfo",
    data(){
      return {
        showFormModal:false,
        default_work_step_types:["work_start","work_end","sql_query","sql_insert"],
        formValidate: {
          work_id: -1,
          work_step_id: -1,
          work_step_name: '',
          work_step_type: '',
        },
        ruleValidate: {
          work_step_name: [
            { required: true, message: 'work_step_name 不能为空!', trigger: 'blur' }
          ],
          work_step_type: [
            { required: true, message: 'work_step_type 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      loadWorkStepInfo:async function(){
        const result = await LoadWorkStepInfo(this.formValidate.work_id,this.formValidate.work_step_id);
        if(result.status == "SUCCESS"){
          this.formValidate.work_step_name = result.step.work_step_name;
          this.formValidate.work_step_type = result.step.work_step_type;
        }
      },
      showWorkStepBaseInfo:function (work_id, work_step_id) {
        this.formValidate.work_id = work_id;
        this.formValidate.work_step_id = work_step_id;
        this.loadWorkStepInfo();
        this.showFormModal = true;
      },
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditWorkStepBaseInfo(this.formValidate.work_id,
              this.formValidate.work_step_id,this.formValidate.work_step_name,this.formValidate.work_step_type);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              this.showFormModal = false;
              // 通知父组件添加成功
              this.$emit('handleSuccess');
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    },
  }
</script>

<style scoped>

</style>
