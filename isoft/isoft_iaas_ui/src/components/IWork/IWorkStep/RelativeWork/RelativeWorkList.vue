<template>
  <span>
    <h4 style="text-align: center;">{{title}}</h4>
    <Table :columns="columns1" :data="workList" size="small"></Table>
  </span>
</template>

<script>
  export default {
    name: "RelativeWorkList",
    props:{
      title:{
        type:String,
        default:"title",
      }
    },
    data(){
      return {
        workList:[],
        columns1: [
          {
            title: 'work_id',
            key: 'id',
            width: 100,
          },
          {
            title: 'work_name',
            key: 'work_name',
          },
          {
            title: 'work_desc',
            key: 'work_desc',
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.$router.push({ path: '/iwork/workstepList',
                        query: { work_id: this.workList[params.index]['id'], work_name: this.workList[params.index]['work_name'] }});
                    }
                  }
                }, '查看'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshWorkList:function (workList) {
        this.workList = workList;
      }
    }
  }
</script>

<style scoped>

</style>
