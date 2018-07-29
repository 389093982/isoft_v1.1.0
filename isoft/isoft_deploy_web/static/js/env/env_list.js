var servername = $("input[name='_servername']").val();

$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var envInfosVueData = {
        envInfos:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var envInfosVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#envInfo_list',
        data: envInfosVueData
    });

    function pageToolFunction(obj) {
        // 渲染分页信息
        $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page, size, count){
                loadPageData(page, size, null);
            }});
    }

    function loadPageData(current_page, page_size, pageToolFunction) {
        $.ajax({
            url: servername + "/env/list",
            method:"post",
            data:{"current_page":current_page, "page_size":page_size},
            async: false,
            success:function (data) {
                var obj = JSON.parse(data);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                envInfosVue.$set(envInfosVueData, 'envInfos', obj.envInfos);
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
            $("#_editEnvInfo_error").html("");
            // 清空表单数据
            clearFormData();
        }});
});

// 清空表单数据
function clearFormData() {
    $("input[id='env_name']").val("");
    $("input[id='env_ip']").val("");
    $("input[id='env_account']").val("");
    $("input[id='env_passwd']").val("");
    $("input[id='deploy_home']").val("");
}

function editEnvInfo() {
    var env_name = $("input[id='env_name']").val();
    var env_ip = $("input[id='env_ip']").val();
    var env_account = $("input[id='env_account']").val();
    var env_passwd = $("input[id='env_passwd']").val();
    var deploy_home = $("input[id='deploy_home']").val();

    $.ajax({
        url: servername + "/env/edit",
        type:"post",
        data:{"env_name":env_name, "env_ip":env_ip, "env_account": env_account, "env_passwd": env_passwd, "deploy_home": deploy_home},
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.reload();
            }else{
                $("#_editEnvInfo_error").html("*" + data.errorMsg);
            }
        }
    });
}

// ssh 连接测试
function connect_test(currentNode, env_id) {
    // 显示加载转圈效果
    $(currentNode).parents("tr").find(".connect").html("");
    $(currentNode).parents("tr").find(".connect").addClass("loading");
    $.ajax({
        url: servername + "/env/connect_test",
        type:"post",
        data:{"env_id":env_id},
        success:function (data) {
            if(data.status=="SUCCESS"){
                $(currentNode).parents("tr").find(".connect").html("<span style='color: green;'>连接正常</span>");
            }else{
                $(currentNode).parents("tr").find(".connect").html("<span style='color: red;'>连接异常</span>");
            }
            // 取消转圈效果
            $(currentNode).parents("tr").find(".connect").removeClass("loading");
        }
    });
}

// 同步 deploy_home 到目标机器指定路径
function sync_deploy_home(currentNode, env_id) {
    // 显示加载转圈效果
    $(currentNode).parents("tr").find(".connect").html("");
    $(currentNode).parents("tr").find(".connect").addClass("loading");
    $.ajax({
        url: servername + "/env/sync_deploy_home",
        type: "post",
        data: {"env_id": env_id},
        success: function (data) {
            if (data.status == "SUCCESS") {
                $(currentNode).parents("tr").find(".connect").html("<span style='color: green;'>同步成功!</span>");
            } else {
                $(currentNode).parents("tr").find(".connect").html("<span style='color: red;'>同步失败!</span>");
            }
            // 取消转圈效果
            $(currentNode).parents("tr").find(".connect").removeClass("loading");
        }
    });
}