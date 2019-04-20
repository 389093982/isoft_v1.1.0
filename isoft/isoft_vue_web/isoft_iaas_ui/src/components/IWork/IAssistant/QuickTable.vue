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
        <CheckboxGroup>
          <ul>
            <li v-for="tableColumn in tableColumns" style="list-style: none;">
              <Checkbox :label="tableColumn"></Checkbox>
            </li>
          </ul>
        </CheckboxGroup>
        <Button size="small" type="primary">全选</Button>
        <Button size="small" type="primary">反选</Button>
      </Col>
      <Col span="20">
        <p style="color: red;">sql信息</p>
        <span v-for="tableSql in tableSqls">
          <p>
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
        customSqls:[],
        appendSqlElements:["select","*","from dual"],
      }
    }
  }
</script>

<style scoped>

</style>
