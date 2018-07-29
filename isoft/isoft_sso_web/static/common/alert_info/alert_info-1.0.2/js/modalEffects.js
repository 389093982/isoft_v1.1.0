(function ($) {

    $.fn.ModalEffects = function(options) {             // options 用于扩展属性
        // 默认属性
        var defaults = {
            "clearUIFunc":function () {}                // 清理 UI界面数据函数,默认为空函数
        };
        // 函数列表
        var methods = {
            initModalEffects:function (options) {
                // 获取文档中 class="md-overlay" 的元素,全局只有一个
                var overlay = document.querySelector('.md-overlay');

                // slice() 方法可从已有的数组中返回选定的元素
                // [].slice.call()或者Array.prototype.slice.call()
                // 这句话相当于Array.slice.call(arguments),目的是将arguments对象的数组提出来转化为数组,arguments本身并不是数组而是对象
                [].slice.call(document.querySelectorAll('.md-trigger')).forEach(function(el, i) {   // 获取所有的 trigger 并循环
                    var modal = document.querySelector('#' + el.getAttribute('data-modal')),    // 获取 data-modal 弹出层元素
                        close = modal.querySelector('.md-close');                               // 进一步获取关闭操作元素

                    function removeModal(hasPerspective) {
                        classie.remove(modal, 'md-show');

                        if (hasPerspective) {
                            classie.remove(document.documentElement, 'md-perspective');
                        }
                    }

                    function removeModalHandler() {
                        removeModal(classie.has(el, 'md-setperspective'));
                    }

                    // addEventListener() 方法用于向指定元素添加事件句柄
                    el.addEventListener('click', function (ev) {                // 点击按钮弹出对话框
                        classie.add(modal, 'md-show');

                        if (classie.has(el, 'md-setperspective')) {     // 判断 el 元素是否有 md-setperspective 样式
                            setTimeout(function () {
                                classie.add(document.documentElement, 'md-perspective');
                            }, 25);
                        }
                    });

                    close.addEventListener('click', function (ev) {   // 关闭操作
                        ev.stopPropagation();
                        removeModalHandler();

                        // 关闭时调用清理 UI 界面数据函数
                        options.clearUIFunc();
                    });
                })
            }
        };


        options = $.extend(defaults, options);
        methods.initModalEffects(options);
        return this;
    };
})(jQuery);
