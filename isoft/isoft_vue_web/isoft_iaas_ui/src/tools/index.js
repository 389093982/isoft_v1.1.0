// 获取 cookie 值
export const getCookie = function getCookie(c_name) {
  if (document.cookie.length > 0) {
    var c_start = document.cookie.indexOf(c_name + "=");
    if (c_start != -1) {
      c_start = c_start + c_name.length+1;
      var c_end=document.cookie.indexOf(";", c_start);
      if (c_end == -1) {
        c_end = document.cookie.length;
      }
      return unescape(document.cookie.substring(c_start,c_end));
    }
  }
  return "";
};

//删除cookie
export const delCookie = function delCookie(name) {
  document.cookie = name+"=;expires="+(new Date(0)).toGMTString();
};

export const checkEmpty = function checkEmpty(checkStr){
  if(checkStr == null || checkStr == undefined || checkStr == ""){
    return true;
  }
  return false;
};

// 跨域设置 cookie
function setCookie (c_name,value,expiredays,domain){
  var exdate = new Date();
  exdate.setDate(exdate.getDate() + expiredays);
  //判断是否需要跨域存储
  if (domain) {
    // egg：path=/;domain=xueersi.com";
    document.cookie = c_name+ "=" +escape(value)+((expiredays==null) ? "" : ";expires="+exdate.toGMTString())+";path=/;domain=" + domain;
  } else {
    document.cookie = c_name+ "=" +escape(value)+((expiredays==null) ? "" : ";expires="+exdate.toGMTString())+";path=/";
  }
}

// 判断值 value 是否是列表 validList 中
export function oneOf (value, validList) {
  for (let i = 0; i < validList.length; i++) {
    if (value === validList[i]) {
      return true;
    }
  }
  return false;
}

export function checkContainsInString(str, subStr) {
  return str.indexOf(subStr) > 0
}
