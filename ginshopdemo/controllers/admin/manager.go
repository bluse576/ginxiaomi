package admin

import (
	"ginshopdemo/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

}
func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}
func (con ManagerController) Doadd(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	//用户名密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或者密码长度不合法", "/admin/manager/add")
	}

	//判断管理员是否存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "管理员已经存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(models.GetUnix()),
	}
	err2 := models.DB.Create(&manager).Error
	if err2 != nil {
		con.Error(c, "增加管理员失败", "/admin/add")
		return
	}
	con.Success(c, "增加管理员成功", "/admin/manager")

	c.String(http.StatusOK, "执行成功")
}
func (con ManagerController) Edit(c *gin.Context) {
	//获取管理员
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	manager := models.Manager{}
	models.DB.Where("id=?", id).Find(&manager)
	//获取角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}
func (con ManagerController) Doedit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := models.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(mobile) > 11 {
		con.Error(c, "电话长度不合法,电话号码长度需要大于11位", "/admin/manager/edit?id="+models.String(id))
		return
	}
	//执行修改
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	//判断是否为空 为空表示不修改密码 不为空表示修改密码
	if password != "" {
		if len(password) < 6 {
			con.Error(c, "密码长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
			return
		}
		manager.Password = models.Md5(password)
	}

	err3 := models.DB.Save(&manager).Error
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/manager")
}
func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Manager{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "删除数据成功", "/admin/role")
	}
}
