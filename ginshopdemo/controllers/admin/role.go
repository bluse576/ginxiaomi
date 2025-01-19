package admin

import (
	"ginshopdemo/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})

	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) Doadd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == " " {
		con.Error(c, "角色标题不能为空", "/admin/role/add")
	}
	role := models.Role{}

	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "增加角色失败,请重试", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}

	c.String(http.StatusOK, "执行增加")
}
func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}

}
func (con RoleController) Doedit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == " " {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	err2 := models.DB.Save(&role).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
		return
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
	// c.String(http.StatusOK, "执行修改")
}
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
}
func (con RoleController) Auth(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取所有的权限
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	//3、获取当前角色拥有的权限，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", id).Find(&roleAccess)

	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}
	//循环遍历所有权限数据，判断当前权限id是否存在角色权限的Map对象中，

	// for _, value := range accessList {
	// 	if _, ok := roleAccessMap[value.Id]; ok {
	// 		value.Checked = true
	// 	}
	// 	for _, v := range value.AccessItem {
	// 		if _, ok := roleAccessMap[v.Id]; ok {
	// 			v.Checked = true
	// 		}
	// 	}
	// }
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
	}

	c.HTML(200, "/admin/role/auth.html", gin.H{
		"roleId":     id,
		"accessList": accessList,
	})
}
func (con RoleController) Doauth(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("roleId"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取权限Id 切片
	accessIds := c.PostFormArray("access_node")
	//删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	//增加数据
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := models.Int(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}
}
