<template>
    <Modal
      v-model="showFormModal"
      width="1000"
      title="查看/编辑 workstep"
      :footer-hide="true"
      :mask-closable="false"
      :styles="{top: '10px'}">
      <Scroll height="450">
        <!-- 表单信息 -->
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
          <Row style="margin-bottom: 20px;text-align: center;color: green;">
            <Col span="4">
              <Button type="warning" size="small" @click="showNext(-1)">加载前置步骤</Button>
            </Col>
            <Col span="16">
                <h2 style='font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;'>步骤名称:{{formValidate.work_step_name}},步骤类型:{{formValidate.work_step_type}}</h2>
            </Col>
            <Col span="4">
              <Button type="warning" size="small" @click="showNext(1)">加载后置步骤</Button>
            </Col>
          </Row>
          <Row style="margin-right: 10px;">
            <Col span="14">
              <FormItem label="work_step_input" prop="work_step_input">
                <Tabs type="card" :animated="false">
                  <TabPane label="edit">
                    <WorkStepParamInputEdit :paramInputSchemaItems="paramInputSchema.ParamInputSchemaItems"/>
                  </TabPane>
                  <TabPane label="ParamMapping" v-if="showParamMapping">
                    <ParamMapping :paramMappings="paramMappings"/>
                  </TabPane>
                </Tabs>
              </FormItem>
            </Col>
            <Col span="10">
              <FormItem label="work_step_output" prop="work_step_output">
                <Tabs type="card" :animated="false">
                  <TabPane label="Tree">
                    <WorkStepPreParamOutputTree v-if="paramOutputSchemaTreeNode" :paramOutputSchemaTreeNode="paramOutputSchemaTreeNode"/>
                  </TabPane>
                </Tabs>
              </FormItem>
            </Col>
          </Row>
          <FormItem>
            <Row>
                <Button type="success" size="small" @click="handleSubmit('formValidate', false)">提交</Button>
              <Button type="warning" size="small" @click="handleSubmit('formValidate', true)">提交并关闭</Button>
            </Row>
          </FormItem>
        </Form>
      </Scroll>
    </Modal>
</template>

<script>
  import WorkStepParamInputEdit from "./WorkStepParamInputEdit"
  import WorkStepPreParamOutputTree from "./WorkStepPreParamOutputTree"
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal"
  import ParamMapping from "./ParamMapping"
  import {EditWorkStepParamInfo} from "../../api"
  import {LoadWorkStepInfo} from "../../api"
  import {oneOf} from "../../tools"

  export default {
    name: "WorkStepParamInfo",
    components:{WorkStepParamInputEdit,WorkStepPreParamOutputTree,ParamMapping,ISimpleConfirmModal},
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
        // 输出参数
        paramOutputSchema:"",
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
      handleSubmit (name, closable) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const paramInputSchemaStr = JSON.stringify(this.paramInputSchema);
            const paramMappingsStr = JSON.stringify(this.paramMappings);
            const result = await EditWorkStepParamInfo(this.formValidate.work_id, this.formValidate.work_step_id, paramInputSchemaStr, paramMappingsStr);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 通知父组件添加成功
              this.$emit('handleSuccess');
              if(closable == true){
                // 直接关闭
                this.showFormModal =false;
              }else{
                // 直接刷新不关闭
                this.showWorkStepParamInfo(this.formValidate.work_id, this.formValidate.work_step_id);
              }
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

          if(oneOf(result.step.work_step_type, ["work_start","work_end","mapper","entity_parser","goto_condition"])){
            this.showParamMapping = true;
          }else{
            this.showParamMapping = false;
          }

          // 入参渲染
          this.paramInputSchema = result.paramInputSchema;
          // 出参渲染
          this.paramOutputSchema = result.paramOutputSchema;
          this.paramOutputSchemaTreeNode = result.paramOutputSchemaTreeNode;
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
      showNext: function(num){
        if((this.formValidate.work_step_type == "work_end" && num > 0) || (this.formValidate.work_step_type == "work_start" && num < 0)){
          return;
        }
        this.showWorkStepParamInfo(this.formValidate.work_id,this.formValidate.work_step_id + num);
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
