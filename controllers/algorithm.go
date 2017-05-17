package controllers

import (
	"dao-service/algorithm-dao-service/models"
	"dao-service/algorithm-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Algorithm
type AlgorithmController struct {
	beego.Controller
}

// @Title Create
// @Description create algorithm
// @Param	body		body 	models.Algorithm	true		"body for algorithm content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *AlgorithmController) Create() {
	var err error
	var algorithm model.Algorithm
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &algorithm)
	if err == nil {
		var svc service.AlgorithmService
		var result []byte
		err = svc.Create(&algorithm)
		if err == nil {
			result, err = json.Marshal(&algorithm)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetAll
// @Description get all algorithms
// @Success 200 {object} models.Response
// @router / [get]
func (this *AlgorithmController) GetAll() {
	var err error
	var response models.Response

	var svc service.AlgorithmService
	var algorithms []*model.Algorithm
	var result []byte
	algorithms, err = svc.GetAll()
	if err == nil {
		result, err = json.Marshal(algorithms)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = string(result)
		}
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetByNameAndVersion
// @Description get algorithm by name and version
// @Param	name		path 	string	true		"The key for staticblock"
// @Param	version		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:name/:version [get]
func (this *AlgorithmController) GetByNameAndVersion() {
	var err error
	var response models.Response

	//var name string
	//var version string
	name := this.GetString(":name")
	version := this.GetString(":version")

	var svc service.AlgorithmService
	var algorithm *model.Algorithm
	var result []byte
	algorithm, err = svc.GetByNameAndVersion(name, version)
	if err == nil {
		result, err = json.Marshal(algorithm)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = string(result)
		}
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title Update
// @Description update the algorithm
// @Param	body		body 	models.Algorithm	true		"body for algorithm content"
// @Success 200 {object} models.Response
// @router / [put]
func (this *AlgorithmController) Update() {
	var err error
	var algorithm model.Algorithm
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &algorithm)
	if err == nil {
		var svc service.AlgorithmService
		var result []byte
		err = svc.Update(&algorithm)
		if err == nil {
			result, err = json.Marshal(&algorithm)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title DeleteById
// @Description delete the algorithm by id
// @Param	id		path 	int64	true		"The int you want to delete"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router /id/:id [delete]
func (this *AlgorithmController) DeleteById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("DeleteById", id)
	if id > 0 && err == nil {
		var svc service.AlgorithmService
		err = svc.DeleteById(id)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "algorithm id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
