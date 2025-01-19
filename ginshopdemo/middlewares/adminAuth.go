package middlewares

import (
	"encoding/json"
	"fmt"
	"ginshopdemo/models"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	excludeAuthPath()
	fmt.Println("InitAdminAuthMiddleware")
	//进行权限判断 没有登录的用户 不能进入后台管理中心

	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	//Split以"？"分割
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// // 2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		} else {
			if userinfoStruct[0].IsSuper == 0 {
				roleAccess := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}

				urlPath := strings.Replace(pathname, "/admin/", "", 1)
				access := models.Access{}
				models.DB.Where("url=?", urlPath).Find(&access)
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}

		}
	} else { //没有登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}

func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	// fmt.Println(excludeAuthPath)
	excludeAuthPathSlice = strings.Split(excludeAuthPath, ",")
	// return true
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
