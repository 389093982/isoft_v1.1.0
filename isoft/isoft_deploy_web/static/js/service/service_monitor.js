var servername = $("input[name='_servername']").val();

$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var serviceMonitorsVueData = {
        serviceMonitors:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var serviceMonitorsVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#service_monitor',
        data: serviceMonitorsVueData
    });

    function pageToolFunction(obj) {
        // 渲染分页信息
        $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page, size, count){
                loadPageData(page, size, null);
            }});
    }

    function loadPageData(current_page, page_size, pageToolFunction) {
        $.ajax({
            url: servername + "/service/monitor",
            method:"post",
            data:{"current_page":current_page, "page_size":page_size},
            async: false,
            success:function (data) {
                var obj = JSON.parse(data);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                serviceMonitorsVue.$set(serviceMonitorsVueData, 'serviceMonitors', obj.serviceMonitors);
                if(pageToolFunction != null){
                    pageToolFunction(obj);              // 渲染分页
                }
            }
        });
    }

    // 加载第一页,10条记录,加载完成之后使用 pageToolFunction 函数进行分页渲染
    loadPageData(1,10,pageToolFunction);

    $(document).ModalEffects({"clearUIFunc":function () {
            // 置空异常信息
            $("#_editServerMonitor_error").html("");
            // 清空表单数据
            clearFormData();
        }});
});

function clearFormData() {

}


function editServiceMonitor(node) {
    url = $("#service_monitor #url").val();
    $.ajax({
        url: servername + "/service/editormonitor",
        type:"post",
        data:{"url":url},
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.reload();
            }else{
                alert("添加失败!");
            }
        }
    });
}



