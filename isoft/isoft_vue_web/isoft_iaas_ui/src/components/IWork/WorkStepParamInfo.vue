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
                  <TabPane label="edit">
                    <WorkStepParamInputEdit :paramInputSchemaItems="paramInputSchema.ParamInputSchemaItems"/>
                  </TabPane>
                  <TabPane label="Xml">
                    <Input v-model.trim="formValidate.work_step_input" type="textarea" :rows="10" placeholder="请输入 work_step_input"></Input>
                  </TabPane>
                  <TabPane label="ParamMapping" v-if="showParamMapping">
                    <ParamMapping :paramMappings="paramMappings"/>
                  </TabPane>
                </Tabs>
              </FormItem>
            </Col>
            <Col span="12">
              <FormItem label="work_step_output" prop="work_step_output">
                <Tabs type="card" :animated="false">
                  <TabPane label="Tree">
                    <WorkStepParamOutputDisplay v-if="paramOutputSchemaTreeNode" :paramOutputSchemaTreeNode="paramOutputSchemaTreeNode"/>
                  </TabPane>
                  <TabPane label="Xml">
                    <Input v-model.trim="formValidate.work_step_output" type="textarea" :rows="10" placeholder="请输入 work_step_output"></Input>
                  </TabPane>
                </Tabs>
              </FormItem>
            </Col>
          </Row>
          <FormItem>
            <Row>
                <Button type="success" @click="handleSubmit('formValidate')">提交</Button>
            </Row>
          </FormItem>
        </Form>
      </Scroll>
    </Modal>
</template>

<script>
  import WorkStepParamInputEdit from "./WorkStepParamInputEdit"
  import WorkStepParamOutputDisplay from "./WorkStepParamOutputDisplay"
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal"
  import ParamMapping from "./ParamMapping"
  import {EditWorkStepParamInfo} from "../../api"
  import {LoadWorkStepInfo} from "../../api"
  import {oneOf} from "../../tools"

  export default {
    name: "WorkStepParamInfo",
    components:{WorkStepParamInputEdit,WorkStepParamOutputDisplay,ParamMapping,ISimpleConfirmModal},
    props: {
      workId: {
        type: Number,
        default: -1
      },
    },
    data(){
      return {
        showFormModal:false,
        // 输入参数
        paramInputSchema:"",
        paramInputSchemaXml:"",
        // 输出参数
        paramOutputSchema:"",
        paramOutputSchemaXml:"",
        paramOutputSchemaTreeNode:null,
        // 显示效果
        showParamMapping:false,
        // 参数映射
        paramMappings:[],
        default_work_step_types: this.GLOBAL.default_work_step_types,
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
            const paramInputSchemaStr = JSON.stringify(this.paramInputSchema);
            const paramMappingsStr = JSON.stringify(this.paramMappings);
            const result = await EditWorkStepParamInfo(this.formValidate.work_id, this.formValidate.work_step_id, paramInputSchemaStr, paramMappingsStr);
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
          this.formValidate.work_id = result.step.work_id;
          this.formValidate.work_step_id = result.step.work_step_id;
          this.formValidate.work_step_name = result.step.work_step_name;
          this.formValidate.work_step_type = result.step.work_step_type;

          if(oneOf(result.step.work_step_type, ["work_start","work_end","mapper"])){
            this.showParamMapping = true;
          }else{
            this.showParamMapping = false;
          }

          // 入参渲染
          this.paramInputSchema = result.paramInputSchema;
          this.paramInputSchemaXml = result.paramInputSchemaXml;
          this.formValidate.work_step_input = result.paramInputSchemaXml;
          // 出参渲染
          this.paramOutputSchema = result.paramOutputSchema;
          this.paramOutputSchemaXml = result.paramOutputSchemaXml;
          this.paramOutputSchemaTreeNode = result.paramOutputSchemaTreeNode;
          this.formValidate.work_step_output = result.paramOutputSchemaXml;
          // 参数映射渲染
          this.paramMappings = result.paramMappings != null ? result.paramMappings : [];
          // 提交 action
          this.$store.dispatch('commitSetCurrent',{"current_work_id":result.step.work_id, "current_work_step_id":result.step.work_step_id});
          // 异步请求加载完成之后才显示模态对话框
          this.showFormModal = true;
        }else{
          // 加载失败
          this.$Message.error('加载失败!');
          this.handleReset('formValidate');
        }
      },
      showWorkStepParamInfo:function (work_id, work_step_id) {
        this.formValidate.work_id = work_id;
        this.formValidate.work_step_id = work_step_id;
        this.loadWorkStepInfo();
      },
    },
  }
</script>

<style scoped>

</style>
