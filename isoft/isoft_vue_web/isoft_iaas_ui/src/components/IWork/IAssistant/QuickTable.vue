<template>
  <div style="margin: 10px;">
    <Row>
      <div>
        <div style="text-align: right;margin-bottom: 5px;">
          <span @drop="deleteElement($event)" @dragover="allowDrop($event)">
           <Button type="primary" size="small" icon="md-close">拖拽至此处进行删除</Button>
          </span>
          <Button type="dashed" size="small" @click="renderSql">Render Sql</Button>
        </div>

        <Row :gutter="10">
          <Col span="8" style="border: 1px solid #dcdee2;padding: 20px;">
            <span v-for="(element,index) in hotSqlElements" draggable="true" @dragstart="dragstart($event, element, -1, 'left')">
              <Button style="margin: 2px;" size="small">{{element}}</Button>
            </span>
          </Col>
          <Col span="16">
            <div style="min-height: 100px;border: 1px solid #dcdee2;padding:20px;" @drop="drop($event, -1)" @dragover="allowDrop($event)">
              <span v-for="(element,index) in appendSqlElements"
                    draggable="true" @dragstart="dragstart($event, element, index, 'right')"
                    @drop="drop($event, index)" @dragover="allowDrop($event)">
                <Tooltip :content="element" placement="bottom" max-width="800">
                  <Button style="margin: 2px;" size="small">{{element | filterLimitFunc}}</Button>
                </Tooltip>
              </span>
            </div>
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
          <li style="list-style: none;" v-for="tableSql in tableSqls" draggable="true" @dragstart="dragstart($event, tableSql, -1, 'bottom')">
            <Tooltip :content="tableSql" placement="bottom" max-width="800">
              <Button style="margin: 2px;" size="small">{{tableSql | filterLimitFunc}}</Button>
            </Tooltip>
          </li>
          <li style="list-style: none;" v-for="(customSql,index) in customSqls" draggable="true" @dragstart="dragstart($event, customSql, -1, 'bottom')">
            <Tooltip :content="customSql" placement="bottom" max-width="800">
              <Button style="margin: 2px;" size="small">自定义sql：{{customSql | filterLimitFunc}}</Button>
            </Tooltip>
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
        split1: 0.4,
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
      deleteElement:function (event) {
        event.preventDefault();
        var dataStr = event.dataTransfer.getData("Text");
        var data = JSON.parse(dataStr);
        var sourceIndex = data.index;
        var transferData = data.transferData;
        var location = data.location;
        if(location == 'right'){
          this.appendSqlElements.splice(sourceIndex, 1);
        }
      },
      dragstart:function(event, transferData, index, location){
        event.dataTransfer.setData("Text", JSON.stringify({'transferData':transferData, 'index':index, 'location':location}));
      },
      allowDrop:function(event){
        event.preventDefault();
      },
      drop:function(event, index){
        // 取消冒泡
        event.stopPropagation();
        event.preventDefault();
        var dataStr = event.dataTransfer.getData("Text");
        var data = JSON.parse(dataStr);
        var sourceIndex = data.index;
        var transferData = data.transferData;
        var location = data.location;
        if(index > 0){  // 目标位置有元素
          if(sourceIndex >= 0){
            // 交换位置
            swapArray(this.appendSqlElements, sourceIndex, index);
          }else{
            // index 后面添加
            this.appendSqlElements.splice(index + 1, 0, data.transferData);
          }
        }else{         // 目标元素为空div,直接追加
          if(location != 'right'){
            this.appendSqlElements.push(transferData);
          }
        }
      }
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 80) {
          value= value.substring(0,80) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>

</style>
