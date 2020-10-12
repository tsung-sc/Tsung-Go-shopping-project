package admin

import (
	"xiaomi/models"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	userinfo, ok := c.GetSession("userinfo").(models.Manager)
	if ok {
		c.Data["username"] = userinfo.Username
		roleId := userinfo.RoleId
		access := []models.Access{}
		models.DB.Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort desc").Where("module_id=?", 0).Find(&access)
		//获取当前部门拥有的权限，并把权限ID放在一个MAP对象里面
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", roleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		for i := 0; i < len(access); i++ {
			if _, ok := roleAccessMap[access[i].Id]; ok {
				access[i].Checked = true
			}
			for j := 0; j < len(access[i].AccessItem); j++ {
				if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
					access[i].AccessItem[j].Checked = true
				}
			}
		}
		c.Data["accessList"] = access
		c.Data["isSuper"] = userinfo.IsSuper
	}
	c.TplName = "admin/main/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}

//修改公共状态
func (c *MainController) ChangeStatus() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法请求",
		}
		c.ServeJSON()
		return
	}
	table := c.GetString("table")
	field := c.GetString("field")
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "更新数据失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "更新数据成功",
	}
	c.ServeJSON()
}

func (c *MainController) EditNum() {
	id := c.GetString("id")
	table := c.GetString("table")
	field := c.GetString("field")
	num := c.GetString("num")
	err1 := models.DB.Exec("update " + table + " set " + field + "=" + num + " where id=" + id).Error
	if err1 != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "修改数量失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "修改数量成功",
	}
	c.ServeJSON()
}
