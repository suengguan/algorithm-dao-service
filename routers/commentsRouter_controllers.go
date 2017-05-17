package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "GetByNameAndVersion",
			Router: `/:name/:version`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"] = append(beego.GlobalControllerRouter["dao-service/algorithm-dao-service/controllers:AlgorithmController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
