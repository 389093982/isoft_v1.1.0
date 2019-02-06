<template>
  <span>
    <Row>
      <ParamMappingAdd @handleSubmit="paramMappingAdd"/>
      <div v-for="paramMapping in paramMappings">
        <Row>
          <Col span="12">{{paramMapping}}</Col>
          <Col span="12" style="text-align: right;"><Button type="success" size="small" @click="handleDelete(paramMapping)" style="margin-left: 6px">删除</Button></Col>
        </Row>
        <hr/>
      </div>
    </Row>
  </span>
</template>

<script>
  import ParamMappingAdd from "./ParamMappingAdd"

  export default {
    name: "ParamMapping",
    components:{ParamMappingAdd},
    data(){
      return {
        paramMappings:[],
      }
    },
    methods:{
      paramMappingAdd:function (data) {
        for(var i=0; i< this.paramMappings.length; i++){
          if(this.paramMappings[i] == data){
            return;
          }
        }
        this.paramMappings.push(data);
      },
      handleDelete:function (data) {
        for(var i=0; i< this.paramMappings.length; i++){
          if(this.paramMappings[i] == data){
            this.paramMappings.splice(i,1);
          }
        }
      }
    }
  }
</script>

<style scoped>

</style>
