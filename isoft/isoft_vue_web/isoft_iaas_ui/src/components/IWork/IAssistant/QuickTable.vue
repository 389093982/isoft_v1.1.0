<template>
  <div style="margin: 10px;">
    <Row>
      <div>
        <div>
          操作：
          <Button type="primary" size="small" icon="md-close" @click="deleteElement">删除</Button>
          <Button type="dashed" size="small" @click="renderSql">Render Sql</Button>
        </div>

        <Row style="margin-top: 10px;margin-bottom: 10px;padding:20px;background-color: #f8f8f9;">
          <Col span="8">
            <span v-for="(element,index) in hotSqlElements" draggable="true" @dragstart="dragstart($event, element, -1)">
              <Button style="margin: 2px;" size="small">{{element}}</Button>
            </span>
          </Col>
          <Col span="16">
            <span v-for="(element,index) in appendSqlElements"
                  draggable="true" @dragstart="dragstart($event, element, index)"
                  @drop="drop($event, index)" @dragover="allowDrop($event)">
              <Button :type="choosedElementIndex == index ? 'primary' : 'default'"
                      style="margin: 2px;" size="small" @click="choosedElementIndex=index">
                {{element}}
              </Button>
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
        <ul>
          <li style="list-style: none;" v-for="tableSql in tableSqls" draggable="true" @dragstart="dragstart($event, tableSql, -1)">
            <Button style="margin: 2px;" size="small">{{tableSql}}</Button>
          </li>
          <li style="list-style: none;" v-for="(customSql,index) in customSqls" draggable="true" @dragstart="dragstart($event, customSql, -1)">
            <Button style="margin: 2px;" size="small">自定义sql：{{customSql}}</Button>
            <Button type="dashed" size="small" @click="deleteCustom(index)">删除</Button>
          </li>
        </ul>
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
      dragstart:function(event, transferData, index){
        event.dataTransfer.setData("Text", JSON.stringify({'transferData':transferData, 'index':index}));
      },
      allowDrop:function(event){
        event.preventDefault();
      },
      drop:function(event, index){
        event.preventDefault();
        var dataStr = event.dataTransfer.getData("Text");
        var data = JSON.parse(dataStr);
        var sourceIndex = data.index;
        var transferData = data.transferData;
        if(sourceIndex >= 0){
          // 交换位置
          swapArray(this.appendSqlElements, sourceIndex, index);
        }else{
          // index 后面添加
          this.appendSqlElements.splice(index + 1, 0, data.transferData);
        }
      }
    }
  }
</script>

<style scoped>

</style>
