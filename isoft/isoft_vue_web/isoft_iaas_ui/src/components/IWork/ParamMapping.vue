<template>
  <span>
    <Row>
      <ParamMappingAdd @handleSubmit="paramMappingAdd"/>
      <div v-for="paramMapping in paramMappings">
        <Row>
          <Col span="12">{{paramMapping}}</Col>
          <Col span="12" style="text-align: right;">
            <Button type="success" size="small" @click="handleDelete(paramMapping)" style="margin-left: 6px">删除</Button>
          </Col>
        </Row>
        <hr/>
      </div>
    </Row>
  </span>
</template>

<script>
  import ParamMappingAdd from "./ParamMappingAdd"
  import {oneOf} from "../../tools"
  import {strSplit} from "../../tools"

  export default {
    name: "ParamMapping",
    components:{ParamMappingAdd},
    props:{
      paramMappings:{
        type:Array,
        default:[],
      }
    },
    methods:{
      paramMappingAdd:function (data) {
        var strs = strSplit(data, ",");
        for(var i=0; i<strs.length; i++){
          if(!oneOf(strs[i].trim(), this.paramMappings)){
            this.paramMappings.push(strs[i].trim());
          }
        }
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
