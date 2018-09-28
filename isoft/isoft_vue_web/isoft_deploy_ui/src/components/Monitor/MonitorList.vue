<template>
  <div>
    <Table :columns="columns1" :data="monitorHeartBeats" size="small" height="450"></Table>
    <Page :total="total" :page-size="pageSize" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageMonitorHeartBeat} from '../../api'

  export default {
    name: "MonitorList",
    data(){
      var _this = this;
      return {
        monitorHeartBeats:[],
        alives:[],
        // 需要隐藏的列
        hiddenColumn:[],
        columns1 : [
          {
            title: 'addr',
            key: 'addr',
            width:150
          },
          {
            title: 'created_by',
            key: 'created_by',
            width:150
          },
          {
            title: 'created_time',
            key: 'created_time',
            width:200,
            render: function (h, params) {
              return h('div',
                new Date(this.row.created_time).Format('yyyy-MM-dd hh:mm:ss'));  // 这里的this.row能够获取当前行的数据
            }
          },
          {
            title: 'last_updated_by',
            key: 'last_updated_by',
            width:150
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time',
            width:200,
            render: function (h, params) {
              return h('div',
                new Date(this.row.last_updated_time).Format('yyyy-MM-dd hh:mm:ss'));  // 这里的this.row能够获取当前行的数据
            }
          },
          {
            title: '运行状态',
            key: 'running_status',
            render: function (h, params) {
              if(_this.alives != null && _this.alives != undefined && $.inArray(this.row.addr, _this.alives)>=0){
                return h('p',{style:{color:"green",}},"正常");
              }else{
                return h('div',{style:{color:"red",}},"异常");
              }
            }
          },
        ],
        // 总页数
        total:1,
        pageNum:1,    // 当前页
        pageSize:10   // 每页数量
      }
    },
    methods:{
      async refreshMonitorHeartBeatList(){
        const _this = this;
        const data = await FilterPageMonitorHeartBeat(this.pageNum, this.pageSize);
        if(data.status=="SUCCESS"){
          this.alives = data.alives;
          this.monitorHeartBeats = data.monitorHeartBeats;
          this.total = data.paginator.totalcount;
        }
      },
      handlePageSizeChange (value){
        // 设置每页数量
        if(value != 10){
          this.pageSize= value;
          this.refreshMonitorHeartBeatList();
        }
      },
      handleChange (value){
        if(value != 1){
          // 设置当前页数
          this.pageNum = value;
          // 重新加载表格数据
          this.refreshMonitorHeartBeatList();
        }
      },
      initFormatter(){
        Date.prototype.Format = function (formatStr) {
          var str = formatStr;
          var Week = ['日', '一', '二', '三', '四', '五', '六'];
          str = str.replace(/yyyy|YYYY/, this.getFullYear());
          str = str.replace(/yy|YY/, (this.getYear() % 100) > 9 ? (this.getYear() % 100).toString() : '0' + (this.getYear() % 100));
          var month = this.getMonth() + 1;
          str = str.replace(/MM/, month > 9 ? month.toString() : '0' + month);
          str = str.replace(/M/g, month);
          str = str.replace(/w|W/g, Week[this.getDay()]);
          str = str.replace(/dd|DD/, this.getDate() > 9 ? this.getDate().toString() : '0' + this.getDate());
          str = str.replace(/d|D/g, this.getDate());
          str = str.replace(/hh|HH/, this.getHours() > 9 ? this.getHours().toString() : '0' + this.getHours());
          str = str.replace(/h|H/g, this.getHours());
          str = str.replace(/mm/, this.getMinutes() > 9 ? this.getMinutes().toString() : '0' + this.getMinutes());
          str = str.replace(/m/g, this.getMinutes());
          str = str.replace(/ss|SS/, this.getSeconds() > 9 ? this.getSeconds().toString() : '0' + this.getSeconds());
          str = str.replace(/s|S/g, this.getSeconds());
          return str;
        }
      },
    },
    mounted:function(){
      this.refreshMonitorHeartBeatList();
    },
    created(){
      // 为Date 对象添加Format方法
      this.initFormatter();
    },
  }
</script>

<style scoped>

</style>
