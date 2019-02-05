<template>
    <Modal
      v-model="showFormModal"
      width="1000"
      title="查看/编辑 workstep"
      :footer-hide="true"
      :mask-closable="false"
      :styles="{top: '10px'}">
      <Scroll height="500">
        <!-- 表单信息 -->
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
          <FormItem label="work_step_name" prop="work_step_name">
            <Input v-model.trim="formValidate.work_step_name" readonly placeholder="请输入 work_step_name"></Input>
          </FormItem>
          <FormItem label="work_step_type" prop="work_step_type">
            <Input v-model.trim="formValidate.work_step_type" readonly placeholder="请输入 work_step_type"></Input>
          </FormItem>
          <Row>
            <Col span="12">
              <FormItem label="work_step_input" prop="work_step_input">
                <Tabs type="card" :animated="false">
                  <TabPane label="Xml">
                    <Input v-model.trim="formValidate.work_step_input" type="textarea" :autosize="{minRows: 10,maxRows: 20}" placeholder="请输入 work_step_input"></Input>
                  </TabPane>
                  <TabPane label="edit">
                    <WorkStepParamInputEdit :paramSchemaItems="paramSchema.ParamSchemaItems"/>
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
    </Modal>
</template>

<script>
  import  WorkStepParamInputEdit from "./WorkStepParamInputEdit"
  import {EditWorkStepParamInfo} from "../../api"
  import {LoadWorkStepInfo} from "../../api"

  export default {
    name: "WorkStepParamInfoDialog",
    components:{WorkStepParamInputEdit},
    props: {
      workId: {
        type: Number,
        default: -1
      },
    },
    data(){
      return {
        showFormModal:false,
        paramSchema:"",
        paramSchemaXml:"",
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
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const paramSchemaStr = JSON.stringify(this.paramSchema);
            const result = await EditWorkStepParamInfo(this.formValidate.work_id, this.formValidate.work_step_id, paramSchemaStr);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              this.showFormModal =false;
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
          this.paramSchema = result.paramSchema;
          this.paramSchemaXml = result.paramSchemaXml;
          this.formValidate.work_step_input = result.paramSchemaXml;
        }else{
          // 加载失败
          this.$Message.error('错误的步骤ID,数据加载失败!数据已失效!');
          this.handleReset('formValidate');
        }
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      showWorkStepParamInfoDialog:function (work_id, work_step_id) {
        this.formValidate.work_id = work_id;
        this.formValidate.work_step_id = work_step_id;
        this.loadWorkStepInfo();
        this.showFormModal = true;
      }
    },
  }
</script>

<style scoped>

</style>
