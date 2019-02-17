<template>
  <!-- 设置 transfer 属性为 false 后,保证好两个 Modal 的前后顺序,可以解决顺序问题 -->
  <Modal
    v-model="showFormModal"
    width="800"
    title="快捷函数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false">
    <div>
      <Table :columns="columns1" :data="funcs" size="small"></Table>
    </div>
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
          {funcDemo:"IworkStringsJoin($str1,$str2)",funcDesc:"字符串拼接函数",},
          {funcDemo:"IworkStringsJoinWithSep($str1,$str2,-)",funcDesc:"字符串拼接函数",},
          {funcDemo:"IworkInt64Add($int1,$int2)",funcDesc:"数字相加函数",},
          {funcDemo:"IworkInt64Sub($int1,$int2)",funcDesc:"数字相减函数",},
          {funcDemo:"IworkInt64Multi($int1,$int2)",funcDesc:"数字相乘函数",},
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
