(function($){
    var app={
        init:function(){    
          
            this.addAddress();
            this.changeDefaultAddress();
            this.editAddress();
            this.doEditAddress();
            this.onSubmit()
       
        },		     
        onSubmit:function(){
     
            $("#checkoutForm").submit(function(){
                var addressCount=$("#addressList .address-item.selected").length;
                if(addressCount>0){
                    return true;
                }				
                alert('请选择收货地址');
                return false;				
            })
            
        },     
        addAddress:function(){

            $("#addAddress").click(function () {
				var name = $('#add_name').val();
				var phone = $('#add_phone').val();
				var address = $('#add_address').val();
				var zipcode = $('#add_zipcode').val();
				if (name == '' || phone == "" || address == "") {
					alert('姓名、电话、地址不能为空')
					return false;
				}
				var reg = /^[\d]{11}$/;
				if (!reg.test(phone)) {
					alert('手机号格式不正确');
					return false;
				}
				$.post('/address/addAddress', { name: name, phone: phone, address: address, zipcode: zipcode }, function (response) {
                    console.log(response)
                    if(response.success){

                        var addressList=response.result;
                        var str=""
                        for (var i = 0; i < addressList.length; i++) {
                            if (addressList[i].default_address) {
								str += '<div class="address-item J_addressItem selected" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
								str += '<dl>';
								str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
								str += '<dd class="utel">' + addressList[i].phone + '</dd>';
								str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
								str += '</dl>';
								str += '<div class="actions">';
								str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
								str += '</div>';
								str += '</div>';

							} else {
								str += '<div class="address-item J_addressItem" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
								str += '<dl>';
								str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
								str += '<dd class="utel">' + addressList[i].phone + '</dd>';
								str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
								str += '</dl>';
								str += '<div class="actions">';
								str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
								str += '</div>';
								str += '</div>';
							}
                        }

                        $("#addressList").html(str)
                    }

                    $('#addModal').modal('hide')
				});
			})
        },   
        changeDefaultAddress:function(){          
            //事件委托 
            $("#addressList").on("click",".J_addressItem",function(){
                $(this).addClass("selected").siblings().removeClass("selected");

                var id = $(this).attr("data-id")
                $.get('/address/changeDefaultAddress?address_id='+id,function(response){})
            })
        },
        editAddress:function(){

            $("#addressList").on("click",".modify",function(e){
				e.stopPropagation();                
                
                var id=$(this).attr('data-id');

                $.get('/address/getOneAddressList?address_id='+id,function(response){
					console.log(response)
                    if(response.success){
                        var addressInfo=response.result;

                        $("#edit_id").val(addressInfo.id);
                        $('#edit_name').val(addressInfo.name);
                        $('#edit_phone').val(addressInfo.phone);
                        $('#edit_address').val(addressInfo.address);
                        $('#edit_zipcode').val(addressInfo.zipcode);

                        $('#editModal').modal('show');
                    } else{
                        alert(response.message)
                    }                
                })
            })

        },
        doEditAddress:function(){            
            $("#doEditAddress").click(function(){

                var id=$("#edit_id").val();
                var name=$('#edit_name').val();
                var phone=$('#edit_phone').val();
                var address=$('#edit_address').val();
                var zipcode=$('#edit_zipcode').val();

                $.post('/address/doEditAddressList',{address_id:id,name:name,phone:phone,address:address,zipcode:zipcode},function(response){
                    
                    if(response.success){

                        var addressList=response.result;
                        var str=""
                        for (var i = 0; i < addressList.length; i++) {
                            if (addressList[i].default_address) {
								str += '<div class="address-item J_addressItem selected" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
								str += '<dl>';
								str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
								str += '<dd class="utel">' + addressList[i].phone + '</dd>';
								str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
								str += '</dl>';
								str += '<div class="actions">';
								str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
								str += '</div>';
								str += '</div>';

							} else {
								str += '<div class="address-item J_addressItem" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
								str += '<dl>';
								str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
								str += '<dd class="utel">' + addressList[i].phone + '</dd>';
								str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
								str += '</dl>';
								str += '<div class="actions">';
								str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
								str += '</div>';
								str += '</div>';
							}
                        }

                        $("#addressList").html(str)
                    }

                    $('#editModal').modal('hide')
                })
                
            })
        }
    }
    $(function(){
        app.init();
    })    
})($)
