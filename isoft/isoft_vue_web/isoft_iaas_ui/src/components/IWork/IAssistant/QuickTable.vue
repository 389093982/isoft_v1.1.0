<template>
  <div style="margin: 20px;">
    <Row>
      <p>
        <span v-for="element in appendSqlElements">
           <Tag>{{element}}</Tag>
        </span>
        <Button icon="ios-add" type="dashed" size="small" @click="alert(111)">添加标签</Button>
      </p>
      <Col span="4">
        <p style="color: red;">表名：{{tableName}}</p>
        <CheckboxGroup v-model="checkTableColumns">
          <ul>
            <li v-for="tableColumn in tableColumns" style="list-style: none;">
              <Checkbox :label="tableColumn"></Checkbox>
            </li>
          </ul>
        </CheckboxGroup>
        <p style="margin-top: 10px;">
          <Button size="small" type="success" @click="chooseAll">全选</Button>
          <Button size="small" type="info" @click="toggleAll">反选</Button>
          <Button size="small" type="warning" @click="appendColumn">拼接</Button>
        </p>
      </Col>
      <Col span="20">
        <p style="color: red;">sql信息</p>
        <span>
          <p v-for="tableSql in tableSqls">
            {{tableSql}} &nbsp;<a href="javascript:;">拷贝</a>
          </p>
          <p v-for="customSql in customSqls">
            自定义sql：{{customSql}} &nbsp;<a href="javascript:;">拷贝</a>
          </p>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {oneOf} from "../../../tools"

  export default {
    name: "QuickTable",
    props:{
      tableName:{
        type:String,
        default:'',
      },
      tableColumns:{
        type:Array,
        default:[],
      },
      tableSqls:{
        type:Array,
        default:[],
      }
    },
    data(){
      return {
        // 选中的列
        checkTableColumns:[],
        customSqls:[],
        appendSqlElements:["select","*","from dual"],
      }
    },
    methods:{
      chooseAll:function () {
        if(this.checkTableColumns.length > 0){
          this.checkTableColumns = [];
        }else{
          this.checkTableColumns = this.tableColumns;
        }
      },
      toggleAll:function () {
        this.checkTableColumns = this.tableColumns.filter(column => !oneOf(column, this.checkTableColumns));
      },
      appendColumn:function () {
        if(this.checkTableColumns.length > 0){
          this.customSqls.push(this.checkTableColumns.join(","));
        }
      }
    }
  }
</script>

<style scoped>

</style>
