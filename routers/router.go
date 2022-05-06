// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-property-server/controllers"
	"go-property-server/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TestController{}, "*:Index")
	beego.Router("/admin/login/dologin", &admin.LoginController{}, "Post:Dologin")
	beego.Router("/admin/login/loginout", &admin.LoginController{}, "Post:Logout")
}
