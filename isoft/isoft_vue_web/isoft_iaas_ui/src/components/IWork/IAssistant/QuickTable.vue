<template>
  <div style="margin: 10px;">
    <Row>
      <div>
        <div>
          操作：
          <Button type="primary" size="small" icon="md-close" @click="deleteElement">删除</Button>
          <Button type="primary" size="small" icon="md-arrow-back" @click="moveElement(-1)">左移</Button>
          <Button type="primary" size="small" icon="md-arrow-forward" @click="moveElement(1)">右移</Button>
          <Button type="dashed" size="small" @click="renderSql">Render Sql</Button>
        </div>

        <Row style="margin-top: 10px;margin-bottom: 10px;padding:20px;background-color: #f8f8f9;">
          <Col span="8">
            <span v-for="(element,index) in hotSqlElements">
              <Tag>
                <span>{{element}}</span>
              </Tag>
            </span>
          </Col>
          <Col span="16">
            <span v-for="(element,index) in appendSqlElements">
              <Tag :color="choosedElementIndex == index ? 'primary' : 'default'">
                <span @click="choosedElementIndex=index">{{element}}</span>
              </Tag>
            </span>
          </Col>
        </Row>
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
          <Button type="dashed" size="small" @click="appendColumn">Apply</Button>
        </p>
      </Col>
      <Col span="20">
        <p style="color: red;">sql信息</p>
        <span>
          <p v-for="tableSql in tableSqls">
            {{tableSql}}
            <Button type="dashed" size="small" @click="insertAfter(tableSql)">Insert After</Button>
          </p>
          <p v-for="(customSql,index) in customSqls">
            自定义sql：{{customSql}}
            <Button type="dashed" size="small" @click="deleteCustom(index)">删除</Button>
            <Button type="dashed" size="small" @click="insertAfter(customSql)">Insert After</Button>
          </p>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {oneOf} from "../../../tools"
  import {swapArray} from "../../../tools"

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
        choosedElementIndex:-1,
        // 选中的列
        checkTableColumns:[],
        customSqls:[],
        // default line 默认线路
        appendSqlElements:["select","*","from dual"],
        hotSqlElements:["select", "(", ")", "count(*) as count","where","from"],
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
      },
      deleteElement:function () {
        this.appendSqlElements.splice(this.choosedElementIndex, 1);
      },
      moveElement:function (direction) {
        if(this.choosedElementIndex > 0 && this.choosedElementIndex < this.appendSqlElements.length - 1){
          swapArray(this.appendSqlElements, this.choosedElementIndex, this.choosedElementIndex + direction);
        }
      },
      insertAfter:function (customSql) {
        if(this.choosedElementIndex > 0){
          // 添加元素到指定位置
          this.appendSqlElements.splice(this.choosedElementIndex+1, 0, customSql);
        }else{
          this.appendSqlElements.push(customSql);
        }
      }
    }
  }
</script>

<style scoped>

</style>
