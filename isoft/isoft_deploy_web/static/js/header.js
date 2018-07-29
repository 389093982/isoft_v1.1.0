var servername = $("input[name='_servername']").val();

$(function () {
    $.ajax({
        url: servername + "/common/checkLoginUser",
        type:"post",
        success:function (data) {
            if(data.isLogin == true){
                $(".login").html(data.username);
                $(".login").mouseenter(function(){
                    $(".login").html("注销");
                });
                $(".login").mouseleave(function(){
                    $(".login").html(data.username);
                });
            }
        }
    });

    $(".login").click(function () {
        var html = $(this).html();
        if(html == "注销"){
            var redirectUrl = document.location.href;
            window.location.href = servername + "/common/logout?redirectUrl=" + redirectUrl + "&time=" + new Date().getTime();
        }else{
            window.location.href = servername + "/common/login?redirectUrl=" + redirectUrl + "&time=" + new Date().getTime();
        }
    });
});