package admin

import (
	"encoding/json"
	"ginshopdemo/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		//获取左侧权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

		//获取当前角色权限
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		//4、循环遍历所有的权限数据,判断id是否在Map对象中
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username": userinfoStruct[0].Username,
		})

	} else {
		c.Redirect(302, "/admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
