<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="流程变量" btn-size="small" modal-title="流程变量" :modal-width="800">

    <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
      <FormItem prop="workVarName" style="width: 250px;">
        <Input type="text" v-model="formInline.workVarName" placeholder="workVarName">
        </Input>
      </FormItem>
      <FormItem prop="workVarType" style="width: 250px;">
        <Select v-model="formInline.workVarType">
          <Option value="string">string</Option>
          <Option value="map">map</Option>
          <Option value="int">int</Option>
        </Select>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formInline')">新增</Button>
      </FormItem>
    </Form>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "WorkWarList",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        formInline: {
          workVarName:'',
          workVarType:'',
        },
        ruleInline: {
          workVarName: [
            { required: true, message: 'Please fill in the user workVarName.', trigger: 'blur' }
          ],
          workVarType: [
            { required: true, message: 'Please fill in the workVarType.', trigger: 'blur' },
          ]
        }
      }
    },
    methods:{
      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('Success!');
          } else {
            this.$Message.error('Fail!');
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
