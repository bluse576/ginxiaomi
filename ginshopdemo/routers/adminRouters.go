package routers

import (
	"ginshopdemo/controllers/admin"
	"ginshopdemo/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doadd", admin.ManagerController{}.Doadd)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doedit", admin.ManagerController{}.Doedit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		adminRouters.POST("/role/doadd", admin.RoleController{}.Doadd)
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/role/doedit", admin.RoleController{}.Doedit)
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admin.RoleController{}.Add)
		adminRouters.POST("/role/doauth", admin.RoleController{}.Doauth)

		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doadd", admin.AccessController{}.Doadd)
		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doedit", admin.AccessController{}.Doedit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)

	}
}
