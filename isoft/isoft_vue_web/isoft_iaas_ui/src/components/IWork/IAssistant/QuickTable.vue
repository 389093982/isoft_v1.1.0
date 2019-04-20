<template>
  <div style="margin: 20px;">
    <Row>
      <Col span="4">
        <p style="color: red;">表名：{{tableName}}</p>
        <CheckboxGroup>
          <ul>
            <li v-for="tableColumn in tableColumns" style="list-style: none;">
              <Checkbox :label="tableColumn"></Checkbox>
            </li>
          </ul>
        </CheckboxGroup>
      </Col>
      <Col span="20">
        <p style="color: red;">sql信息</p>
        <span v-for="tableSql in tableSqls">
          <p>{{tableSql}}</p>
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
  }
</script>

<style scoped>

</style>
