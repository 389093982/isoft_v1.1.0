<template>
  <!-- 设置 transfer 属性为 false 后,保证好两个 Modal 的前后顺序,可以解决顺序问题 -->
  <Modal
    v-model="showFormModal"
    width="800"
    title="快捷函数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
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
          {funcDemo:"stringsToUpper($str)",funcDesc:"字符串转大写函数",},
          {funcDemo:"stringsToLower($str)",funcDesc:"字符串转小写函数",},
          {funcDemo:"stringsJoin($str1,$str2)",funcDesc:"字符串拼接函数",},
          {funcDemo:"stringsJoinWithSep($str1,$str2,-)",funcDesc:"字符串拼接函数",},
          {funcDemo:"int64Add($int1,$int2)",funcDesc:"数字相加函数",},
          {funcDemo:"int64Sub($int1,$int2)",funcDesc:"数字相减函数",},
          {funcDemo:"int64Multi($int1,$int2)",funcDesc:"数字相乘函数",},
          {funcDemo:"stringsContains($str1,$str2)",funcDesc:"字符串包含函数",},
          {funcDemo:"stringsHasPrefix($str1,$str2)",funcDesc:"字符串前缀判断函数",},
          {funcDemo:"stringsHasSuffix($str1,$str2)",funcDesc:"字符串后缀判断函数",},
          {funcDemo:"int64Gt($int1,$int2)",funcDesc:"判断数字1是否大于数字2",},
          {funcDemo:"int64Lt($int1,$int2)",funcDesc:"判断数字1是否小于数字2",},
          {funcDemo:"int64Eq($int1,$int2)",funcDesc:"判断数字1是否等于数字2",},
          {funcDemo:"and($bool1,$bool2)",funcDesc:"判断bool1和bool2同时满足",},
          {funcDemo:"or($bool,$bool2)",funcDesc:"判断bool1和bool2只要一个满足即可",},
          {funcDemo:"not($bool)",funcDesc:"bool值取反",},
          {funcDemo:"uuid()",funcDesc:"生成随机UUID信息",},
          {funcDemo:"isempty($var)",funcDesc:"判断变量或者字符串是否为空",},
          {funcDemo:"getDirPath($filepath)",funcDesc:"获取当前文件父级目录的绝对路径",},
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
