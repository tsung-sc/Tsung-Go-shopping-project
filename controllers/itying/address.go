package itying

import "xiaomi/models"

type AddressController struct {
	BaseController
}

func (c *AddressController) AddAddress() {
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	zipcode := c.GetString("zipcode")
	var addressCount int
	models.DB.Where("uid=?", user.Id).Table("address").Count(&addressCount)
	if addressCount > 10 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "增加收货地址失败，收货地址数量超过限制",
		}
		c.ServeJSON()
		return
	}
	models.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressResult := models.Address{
		Uid:            user.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		Zipcode:        zipcode,
		DefaultAddress: 1,
	}
	models.DB.Create(&addressResult)
	allAddressResult := []models.Address{}
	models.DB.Where("uid=?", user.Id).Find(&allAddressResult)
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.ServeJSON()
}

func (c *AddressController) GetOneAddressList() {
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	address := models.Address{}
	models.DB.Where("id=?", addressId).Find(&address)
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  address,
	}
	c.ServeJSON()
}

func (c *AddressController) DoEditAddressList() {
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	zipcode := c.GetString("zipcode")
	models.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressModel := models.Address{}
	models.DB.Where("id=?", addressId).Find(&addressModel)
	addressModel.Name = name
	addressModel.Phone = phone
	addressModel.Address = address
	addressModel.Zipcode = zipcode
	addressModel.DefaultAddress = 1
	models.DB.Save(&addressModel)
	// 查询当前用户的所有收货地址并返回
	allAddressResult := []models.Address{}
	models.DB.Where("uid=?", user.Id).Order("default_address desc").Find(&allAddressResult)

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.ServeJSON()

}

func (c *AddressController) ChangeDefaultAddress() {
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	models.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	models.DB.Table("address").Where("id=?", addressId).Updates(map[string]interface{}{"default_address": 1})
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  "更新默认收获地址成功",
	}
	c.ServeJSON()
}
