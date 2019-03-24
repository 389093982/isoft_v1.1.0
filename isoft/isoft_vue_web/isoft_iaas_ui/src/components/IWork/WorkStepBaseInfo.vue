<template>
  <Modal
    v-model="showFormModal"
    width="800"
    title="查看/编辑 workstep"
    :footer-hide="true"
    :mask-closable="false"
    :styles="{top: '20px'}">
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
          <Row>
            <Col span="18">
              <Input v-model.trim="formValidate.work_step_type" placeholder="请输入 work_step_type"></Input>
            </Col>
            <Col span="6">
              <Poptip v-model="visible" placement="left-start" width="420">
                <Button style="margin-left: 5px;">选择步骤类型</Button>
                <div slot="content">
                  <span v-for="default_work_step_type in default_work_step_types" style="margin: 5px;float: left;">
                    <Tag><span @click="closePoptip(default_work_step_type.name)">{{default_work_step_type.name}}</span></Tag>
                  </span>
                </div>
              </Poptip>
            </Col>
          </Row>
        </FormItem>
        <FormItem label="is_defer" prop="is_defer">
          <Select v-model="formValidate.is_defer">
            <Option value="true">true</Option>
            <Option value="false">false</Option>
          </Select>
        </FormItem>
        <FormItem label="work_step_desc" prop="work_step_desc">
          <Input v-model.trim="formValidate.work_step_desc" type="textarea" :rows="4" placeholder="请输入 work_step_desc"></Input>
        </FormItem>
        <FormItem>
          <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        </FormItem>
      </Form>
    </div>
  </Modal>
</template>

<script>
  import {EditWorkStepBaseInfo} from "../../api"
  import {LoadWorkStepInfo} from "../../api"
  import {validateCommonPatternForString} from "../../tools"

  export default {
    name: "WorkStepBaseInfo",
    data(){
      const _validateWorkStepName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('字段值不能为空!'));
        } else if (!validateCommonPatternForString(value)) {
          callback(new Error('存在非法字符，只能包含字母，数字，下划线!'));
        } else {
          callback();
        }
      };
      return {
        showFormModal:false,
        visible:false,
        default_work_step_types: this.GLOBAL.default_work_step_types,
        formValidate: {
          work_id: -1,
          work_step_id: -1,
          work_step_name: '',
          work_step_type: '',
          work_step_desc: '',
          is_defer: 'false',
        },
        ruleValidate: {
          work_step_name: [
            { validator: _validateWorkStepName, trigger: 'blur' },
          ],
          work_step_type: [
            { required: true, message: 'work_step_type 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      closePoptip (type) {
        this.formValidate.work_step_type=type;
        this.visible = false;
      },
      loadWorkStepInfo:async function(){
        const result = await LoadWorkStepInfo(this.formValidate.work_id,this.formValidate.work_step_id);
        if(result.status == "SUCCESS"){
          this.formValidate.work_step_name = result.step.work_step_name;
          this.formValidate.work_step_type = result.step.work_step_type;
          this.formValidate.work_step_desc = result.step.work_step_desc;
          this.formValidate.is_defer = result.step.is_defer;
        }
      },
      showWorkStepBaseInfo:function (work_id, work_step_id) {
        // 重置表单,清除缓存
        this.$refs["formValidate"].resetFields();
        this.formValidate.work_id = work_id;
        this.formValidate.work_step_id = work_step_id;
        this.loadWorkStepInfo();
        this.showFormModal = true;
      },
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditWorkStepBaseInfo(this.formValidate.work_id, this.formValidate.work_step_id,
                this.formValidate.work_step_name,this.formValidate.work_step_desc, this.formValidate.work_step_type, this.formValidate.is_defer);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              this.showFormModal = false;
              // 通知父组件添加成功
              this.$emit('handleSuccess');
              // 重置表单,清除缓存
              this.$refs[name].resetFields();
            }else{
              this.$Message.error('提交失败!参数不合法或者步骤名称已存在!');
            }
          }
        })
      },
    },
  }
</script>

<style scoped>

</style>
