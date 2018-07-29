var servername = $("input[name='_servername']").val();

function modifyService() {
    var service_id = $("input[name='service_id']").val();
    var service_name = $("input[name='service_name']").val();
    var service_type = $("input[name='service_type']").val();
    var service_port = $("input[name='service_port']").val();
    var package_name = $("input[name='package_name']").val();
    var run_mode = $("input[name='run_mode']").val();

    $.ajax({
        url: servername + "/service/modify",
        type:"post",
        data: {
            "service_id": service_id,
            "service_port": service_port,
            "package_name": package_name,
            "run_mode": run_mode
        },
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.href = servername + "/service/list?service_type=" + service_type;
            }
        }
    });
}