var MainJs = function () {
    var load = function () {
        $('.navbar li').find('a').each(function () {
            if (this.href == document.location.href || document.location.href.search(this.href) >= 0) {
                $(this).parent().siblings('li').removeClass('active');
                $(this).parent().addClass('active');
            }
        });
    };

    var nav = function () {
        $('.navbar li').click(function (e) {
            $(e.target).parent().siblings('li').removeClass('active');
            $(e.target).parent().addClass('active');
        });
    };
    return {
        init: function () {
            load();
            nav();
        }
    };
}();

jQuery(document).ready(function () {
    MainJs.init();
});