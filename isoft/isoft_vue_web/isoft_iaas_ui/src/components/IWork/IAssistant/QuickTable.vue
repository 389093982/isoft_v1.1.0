<template>
  <div style="margin: 10px;">
    <Row>
      <div>
        <div>
          操作：
          <Button type="primary" size="small" icon="md-close" @click="deleteElement">删除</Button>
          <Button type="dashed" size="small" @click="renderSql">Render Sql</Button>
        </div>

        <div class="demo-split">
          <Split v-model="split1">
            <div slot="left" class="demo-split-pane">
              <span v-for="(element,index) in hotSqlElements" draggable="true" @dragstart="dragstart($event, element, -1)">
                <Button style="margin: 2px;" size="small">{{element}}</Button>
              </span>
            </div>
            <div slot="right" class="demo-split-pane">
              <div style="min-height: 100px;" @drop="drop($event, -1)" @dragover="allowDrop($event)">
                <span v-for="(element,index) in appendSqlElements"
                      draggable="true" @dragstart="dragstart($event, element, index)"
                      @drop="drop($event, index)" @dragover="allowDrop($event)">
                  <Button :type="choosedElementIndex == index ? 'primary' : 'default'"
                          style="margin: 2px;" size="small" @click="choosedElementIndex=index">
                    {{element}}
                  </Button>
                </span>
              </div>
            </div>
          </Split>
        </div>
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
        split1: 0.4,
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
      deleteElement:function (event) {
        this.appendSqlElements.splice(this.choosedElementIndex, 1);
      },
      dragstart:function(event, transferData, index){
        event.dataTransfer.setData("Text", JSON.stringify({'transferData':transferData, 'index':index}));
      },
      allowDrop:function(event){
        event.preventDefault();
      },
      drop:function(event, index){
        event.stopPropagation();
        event.preventDefault();
        var dataStr = event.dataTransfer.getData("Text");
        var data = JSON.parse(dataStr);
        var sourceIndex = data.index;
        var transferData = data.transferData;
        if(index > 0){
          if(sourceIndex >= 0){
            // 交换位置
            swapArray(this.appendSqlElements, sourceIndex, index);
          }else{
            // index 后面添加
            this.appendSqlElements.splice(index + 1, 0, data.transferData);
          }
        }else{
          this.appendSqlElements.push(transferData);
        }

      }
    }
  }
</script>

<style scoped>
  .demo-split{
    height: 200px;
    border: 1px solid #dcdee2;
  }
  .demo-split-pane{
    padding: 10px;
  }
</style>
