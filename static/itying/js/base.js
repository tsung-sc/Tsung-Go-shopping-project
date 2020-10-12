(function($){

    var app={
        init:function(){
    
            this.initSwiper();

            this.initNavSlide();

            this.initContentTabs();

            this.initContentColor();

        },
        initSwiper:function(){    
            new Swiper('.swiper-container', {
                loop : true,
                navigation: {
                  nextEl: '.swiper-button-next',
                  prevEl: '.swiper-button-prev'                 
                },
                pagination: {
                    el: '.swiper-pagination',
                    clickable :true
                },
                effect:'fade'   //动画效果：渐变
                
            });
        },
        initNavSlide:function(){
             $("#nav_list>li").hover(function(){

                $(this).find('.children-list').show();
             },function(){
                $(this).find('.children-list').hide(); 
             })          

        },
        initContentTabs:function(){

            $(function () {
                $('.detail_info .detail_info_item:first').addClass('active');
                $('.detail_list li:first').addClass('active');
                $('.detail_list li').click(function () {
                    var index = $(this).index();
                    $(this).addClass('active').siblings().removeClass('active');
                    $('.detail_info .detail_info_item').removeClass('active').eq(index).addClass('active');
        
                })
            })	
        },

        /*
         {
                        success:true,
                        result:[
                            {
                                img_url:"图片的地址"
                            }，
                             {
                                img_url:"图片的地址"
                            }
                        ]
          }
        */
        initContentColor:function(){
            var _that=this;
          
            $("#color_name").html($("#color_list .active .yanse").html())

            $("#color_list .banben").click(function(){
                $(this).addClass("active").siblings().removeClass("active")

                $("#color_name").html($("#color_list .active .yanse").html())
                var color_id = $(this).attr('color_id');
                var goods_id = $(this).attr('goods_id');
                $.get("/product/getImgList",{"goods_id":goods_id,"color_id":color_id},function(response){
                    if(response.success){
                        var data=response.result;
                        var str=""
                        for(var i=0;i<data.length;i++){
                            str += '<div class="swiper-slide"><img src="' + data[i].img_url + '"> </div>';
                        }

                        $("#item_focus").html(str)
                        _that.initSwiper()
                    }
                   
                })
                
            })
        }
    }   

    
    $(function(){
    
    
        app.init();
    })

    

})($)
