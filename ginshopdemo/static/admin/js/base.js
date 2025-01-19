
$(function(){
    baseApp.init();
})
var baseApp={
    init:function(){
        this.initAside()
    },
    initAside:function(){
        $(function(){
			$('.aside h4').click(function(){		
				
				$(this).siblings('ul').slideToggle();
			})
		})
    },//删除提示
    confirmDelete:function(){
		$(".delete").click(function(){
			var flag = confirm("您确定要删除吗?")
			return flag
		})
	}
}