<template>
  <div style="margin: 10px;">
    <Row>
      <div>
        <span v-for="element in appendSqlElements">
           <Tag>{{element}}</Tag>
        </span>
        <Button icon="ios-add" type="dashed" size="small" @click="alert(111)">添加标签</Button>
        <Button type="dashed" size="small" @click="renderSql">Render Sql</Button>
      </div>
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
          <Button type="dashed" size="small" @click="chooseAll">全选</Button>
          <Button type="dashed" size="small" @click="toggleAll">反选</Button>
          <Button type="dashed" size="small" @click="appendColumn">拼接</Button>
        </p>
      </Col>
      <Col span="20">
        <p style="color: red;">sql信息</p>
        <span>
          <p v-for="tableSql in tableSqls">
            {{tableSql}}
            <Button type="dashed" size="small">应用</Button>
          </p>
          <p v-for="(customSql,index) in customSqls">
            自定义sql：{{customSql}}
            <Button type="dashed" size="small" @click="deleteCustom(index)">删除</Button>
            <Button type="dashed" size="small">应用</Button>
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
      },
      deleteCustom:function (index) {
        this.customSqls.splice(index,1);
      },
      renderSql:function () {
        alert(this.appendSqlElements.join(" "));
      }
    }
  }
</script>

<style scoped>

</style>
