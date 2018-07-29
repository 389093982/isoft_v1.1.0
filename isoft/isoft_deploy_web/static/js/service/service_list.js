var servername = $("input[name='_servername']").val();

$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var serviceInfosVueData = {
        serviceInfos:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var serviceInfosVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#service_list',
        data: serviceInfosVueData
    });

    function pageToolFunction(obj) {
        // 渲染分页信息
        $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page, size, count){
                loadPageData(page, size, null);
            }
        });
    }

    function loadPageData(current_page, page_size, pageToolFunction) {
        $.ajax({
            url: servername + "/service/list",
            method:"post",
            data: {
                "current_page": current_page,
                "page_size": page_size,
                "service_type": $("input[name='_service_type']").val()
            },
            async: false,
            success:function (data) {
                var obj = JSON.parse(data);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                serviceInfosVue.$set(serviceInfosVueData, 'serviceInfos', obj.serviceInfos);
                if(pageToolFunction != null){
                    pageToolFunction(obj);              // 渲染分页
                }
                // $nextTick 是在下次 DOM 更新循环结束之后执行延迟回调，在修改数据之后使用 $nextTick，则可以在回调中获取更新后的 DOM
                serviceInfosVue.$nextTick(function () {
                    renderLastDeployStatus();
                })

            }
        });
    }

    // 加载第一页,10条记录,加载完成之后使用 pageToolFunction 函数进行分页渲染
    loadPageData(1,10,pageToolFunction);

    $(document).ModalEffects({"clearUIFunc":function () {
            // 置空异常信息
            $("#_editEnvInfo_error").html("");
            // 清空表单数据
            clearFormData();
        }});
});

// 清空表单数据
function clearFormData() {

}

function editServiceInfo(){
    var env_ids = $("select[id='env_id']").val();
    var service_name = $("input[id='service_name']").val();
    var service_type = $("select[id='service_type']").val();
    var service_port = $("input[id='service_port']").val();

    $.ajax({
        url: servername + "/service/edit",
        type:"post",
        data: {
            "env_ids": env_ids,
            "service_name": service_name,
            "service_type": service_type,
            "service_port": service_port
        },
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.reload();
            }else{
                $("#_editServerInfo_error").html("*" + data.errorMsg);
            }
        }
    });
}

function renderLastDeployStatus() {
    $("table td").each(function () {
        if ($(this).attr("service_type") == "beego" || $(this).attr("service_type") == "nginx" || $(this).attr("service_type") == "mysql") {
            _renderLastDeployStatus($(this).attr("service_id"), this)
        }
    });
}

function _renderLastDeployStatus(service_id, tdNode) {
    // 渲染成转圈效果
    renderTrackingStatus(null, tdNode);
    var interval = setInterval(function () {
        // 查询最后一次部署状态
        queryLastDeployStatus(service_id, tdNode);
    }, 2000);

    function queryLastDeployStatus(service_id, tdNode) {
        $.ajax({
            url: servername + "/service/queryLastDeployStatus",
            type: "post",
            data: {"service_id": service_id},
            success: function (data) {
                if (data.status == "SUCCESS") {
                    renderTrackingStatus(data.trackingStatus, tdNode);
                    if (data.finish == true) {
                        clearInterval(interval);
                    }
                }
            }
        });
    }

    function renderTrackingStatus(trackingStatus, tdNode) {
        if (trackingStatus !== null && trackingStatus !== undefined && trackingStatus !== '') {
            $(tdNode).siblings(".deploy_status").html("<div style='color: green;'>" + trackingStatus + "</div>");
        } else {
            $(tdNode).siblings(".deploy_status").html("<div class='loading'></div>");
        }
    }
}

function runDeployTask(node, operate_type) {
    var tdNode = $(node).parents("td");
    var env_id = $(tdNode).attr("env_id");
    var service_id = $(tdNode).attr("service_id");

    $.ajax({
        url: servername + "/service/runDeployTask",
        type:"post",
        data:{"env_id":env_id, "service_id":service_id, "operate_type":operate_type},
        success:function (data) {
            if(data.status=="SUCCESS"){
                _renderLastDeployStatus(service_id, tdNode);
            }
        }
    });
}

// 跳往服务编辑页面
function forwardEditPage(node) {
    window.location.href = servername + "/service/edit?service_id=" + $(node).parents("td").attr("service_id");
}

function showServiceTrackingLogDetail(node) {
    window.location.href = servername + "/service/showServiceTrackingLogDetail?service_id=" + $(node).parents("td").attr("service_id");
}

function DockerInfoCheck(node) {
    $.ajax({
        url: servername + "/docker/dockerInfoCheck",
        type: "post",
        data: {"env_id": $(node).parents("td").attr("env_id")},
        success: function (data) {
            if (data.status == "SUCCESS") {
                console.log(data)
            }
        }
    });
}