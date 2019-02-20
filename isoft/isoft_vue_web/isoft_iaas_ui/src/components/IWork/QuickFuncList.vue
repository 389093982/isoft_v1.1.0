<template>
  <!-- 设置 transfer 属性为 false 后,保证好两个 Modal 的前后顺序,可以解决顺序问题 -->
  <Modal
    v-model="showFormModal"
    width="800"
    title="快捷函数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false">
    <Scroll height="380">
      <Table :columns="columns1" :data="funcs" size="small"></Table>
    </Scroll>
  </Modal>
</template>

<script>
  export default {
    name: "QuickFuncList",
    data(){
      return {
        showFormModal:false,
        funcs:[
          {funcDemo:"IworkStringsToUpper($str)",funcDesc:"字符串转大写函数",},
          {funcDemo:"IworkStringsToLower($str)",funcDesc:"字符串转小写函数",},
          {funcDemo:"IworkStringsJoin($str1,$str2)",funcDesc:"字符串拼接函数",},
          {funcDemo:"IworkStringsJoinWithSep($str1,$str2,-)",funcDesc:"字符串拼接函数",},
          {funcDemo:"IworkInt64Add($int1,$int2)",funcDesc:"数字相加函数",},
          {funcDemo:"IworkInt64Sub($int1,$int2)",funcDesc:"数字相减函数",},
          {funcDemo:"IworkInt64Multi($int1,$int2)",funcDesc:"数字相乘函数",},
          {funcDemo:"IworkStringsContains($str1,$str2)",funcDesc:"字符串包含函数",},
          {funcDemo:"IworkStringsHasPrefix($str1,$str2)",funcDesc:"字符串前缀判断函数",},
          {funcDemo:"IworkStringsHasSuffix($str1,$str2)",funcDesc:"字符串后缀判断函数",},
          {funcDemo:"IworkInt64Gt($int1,$int2)",funcDesc:"判断数字1是否大于数字2",},
          {funcDemo:"IworkInt64Lt($int1,$int2)",funcDesc:"判断数字1是否小于数字2",},
          {funcDemo:"IworkInt64Eq($int1,$int2)",funcDesc:"判断数字1是否等于数字2",},
          {funcDemo:"IworkBoolAnd($bool1,$bool2)",funcDesc:"判断bool1和bool2同时满足",},
          {funcDemo:"IworkBoolOr($bool,$bool2)",funcDesc:"判断bool1和bool2只要一个满足即可",},
          {funcDemo:"IworkBoolNot($bool)",funcDesc:"bool值取反",},
          {funcDemo:"IworkStringsUUID()",funcDesc:"生成随机UUID信息",},
        ],
        columns1: [
          {
            title: 'funcDesc',
            key: 'funcDesc',
          },
          {
            title: 'funcDemo',
            key: 'funcDemo',
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.showFormModal = false;
                      this.$emit("chooseFunc", this.funcs[params.index]['funcDemo']);
                    }
                  }
                }, '复制'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      showModal: function(){
        this.showFormModal = true;
      }
    }
  }
</script>

<style scoped>

</style>
