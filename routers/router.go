// @APIVersion 1.0.0
// @Title algorithm-dao-service API
// @Description algorithm-dao-service only serve the ALGORITHM_T
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dao-service/algorithm-dao-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/algorithm",
			beego.NSInclude(
				&controllers.AlgorithmController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
