<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="查看/编辑" modal-title="查看/编辑 workstep" :modal-width="1000">
    <Scroll height="500">
      <!-- 表单信息 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
        <FormItem label="work_id" prop="work_id">
          <Input v-model.trim="formValidate.work_id" disabled placeholder="请输入 work_id"></Input>
        </FormItem>
        <FormItem label="work_step_id" prop="work_step_id">
          <Row :gutter="1">
            <Col span="16">
              <Select v-model="formValidate.work_step_id" placeholder="请选择 work_step_id">
                <Option v-for="_step in all_steps" :value="_step.work_step_id">{{_step.work_step_id}}</Option>
              </Select>
            </Col>
            <Col span="6">
              <Button type="success" @click="loadWorkStepInfo">加载并编辑指定步骤 ID</Button>
            </Col>
          </Row>
        </FormItem>
        <FormItem label="work_step_name" prop="work_step_name">
          <Input v-model.trim="formValidate.work_step_name" placeholder="请输入 work_step_name"></Input>
        </FormItem>
        <FormItem label="work_step_type" prop="work_step_type">
          <Select v-model="formValidate.work_step_type" placeholder="请选择 work_step_type">
            <Option :value="default_work_step_type" v-for="default_work_step_type in default_work_step_types">{{default_work_step_type}}</Option>
          </Select>
        </FormItem>
        <Row>
          <Col span="12">
            <FormItem label="work_step_input" prop="work_step_input">
              <Tabs type="card" :animated="false">
                <TabPane label="Xml">
                  <Input v-model.trim="formValidate.work_step_input" type="textarea" :autosize="{minRows: 10,maxRows: 20}" placeholder="请输入 work_step_input"></Input>
                </TabPane>
                <TabPane label="edit">
                  <WorkStepInputEdit :paramDefinitionItems="paramDefinition.ParamDefinitionItems"/>
                </TabPane>
              </Tabs>
            </FormItem>
          </Col>
          <Col span="12">
            <FormItem label="work_step_output" prop="work_step_output">
              <Tabs type="card" :animated="false">
                <TabPane label="Xml">
                  <Input v-model.trim="formValidate.work_step_output" type="textarea" :autosize="{minRows: 10,maxRows: 20}" placeholder="请输入 work_step_output"></Input>
                </TabPane>
                <TabPane label="edit">
                  <span>AAAAAAAAAAA</span>
                </TabPane>
              </Tabs>
            </FormItem>
          </Col>
        </Row>
        <FormItem>
          <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
          <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
        </FormItem>
      </Form>
    </Scroll>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import  WorkStepInputEdit from "./WorkStepInputEdit"
  import {EditWorkStep} from "../../api"
  import {LoadWorkStepInfo} from "../../api"
  import {GetAllWorkStepInfo} from "../../api"

  export default {
    name: "WorkStepEdit",
    components:{ISimpleBtnTriggerModal,WorkStepInputEdit},
    props: {
      workId: {
        type: Number,
        default: -1
      },
    },
    data(){
      return {
        paramDefinition:"",
        paramDefinitionXml:"",
        // 所有的步骤信息,主要用于下拉列表使用
        all_steps:[],
        default_work_step_types:["work_start","work_end","sql_query","sql_insert"],
        formValidate: {
          work_id: this.workId,
          work_step_id: 0,
          work_step_name: '',
          work_step_type: '',
          work_step_input: '',
          work_step_output: '',
        },
        ruleValidate: {
          work_step_id: [
            { required: true, type: 'number', message: 'work_step_id 必须为数字且不能为空!', trigger: 'blur' },
          ],
          work_step_name: [
            { required: true, message: 'work_step_name 不能为空!', trigger: 'blur' }
          ],
          work_step_type: [
            { required: true, message: 'work_step_type 不能为空!', trigger: 'blur' }
          ],
          work_step_input: [
            { required: true, message: 'work_step_input 不能为空!', trigger: 'blur' }
          ],
          work_step_output: [
            { required: true, message: 'work_step_output 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditWorkStep(this.formValidate.work_id,
              this.formValidate.work_step_id,this.formValidate.work_step_name,this.formValidate.work_step_type,
              this.formValidate.work_step_input,this.formValidate.work_step_output);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 调用子组件隐藏 modal (this.refs.xxx.子组件定义的方法())
              this.$refs.triggerModal.hideModal();
              // 通知父组件添加成功
              this.$emit('handleSuccess');
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      loadWorkStepInfo:async function(){
        const result = await LoadWorkStepInfo(this.formValidate.work_id,this.formValidate.work_step_id);
        if(result.status == "SUCCESS"){
          this.formValidate.work_step_name = result.step.work_step_name;
          this.formValidate.work_step_type = result.step.work_step_type;
          this.formValidate.work_step_input = result.step.work_step_input;
          this.formValidate.work_step_output = result.step.work_step_output;
          this.paramDefinition = result.paramDefinition;
          this.paramDefinitionXml = result.paramDefinitionXml;
          this.formValidate.work_step_input = result.paramDefinitionXml;
        }else{
          // 加载失败
          this.$Message.error('错误的步骤ID,数据加载失败!数据已失效!');
          this.handleReset('formValidate');
        }
      },
      refreshAllWorkStepsInfo:async function(){
        const result = await GetAllWorkStepInfo(this.formValidate.work_id);
        if(result.status == "SUCCESS"){
          this.all_steps = result.steps;
        }
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    },
    mounted:function () {
      this.refreshAllWorkStepsInfo();
    }
  }
</script>

<style scoped>

</style>
