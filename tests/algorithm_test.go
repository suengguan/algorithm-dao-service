package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	_ "dao-service/algorithm-dao-service/routers"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/algorithm"
)

func init() {
}

func Test_Create(t *testing.T) {
	var algorithm model.Algorithm
	algorithm.Id = 0
	algorithm.Name = "algorithm-name"
	algorithm.Version = "v1"
	algorithm.Description = "this is algorithm"
	algorithm.Author = "author"
	algorithm.Parameters = "parameter=1"
	algorithm.Image = ""
	algorithm.Downloads = 0
	algorithm.Stars = 0

	// post create algorithm
	requestData, err := json.Marshal(&algorithm)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err := http.Post(base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}

func Test_GetAll(t *testing.T) {
	res, err := http.Get(base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}

func Test_GetByNameAndVersion(t *testing.T) {
	res, err := http.Get(base_url + "/algorithm-name" + "/v1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}

func Test_Update(t *testing.T) {
	// get all
	res1, err := http.Get(base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res1.Body.Close()
	resBody1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody1))

	var response model.Response
	json.Unmarshal(resBody1, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var algorithms []*model.Algorithm
	json.Unmarshal(([]byte)(response.Result), &algorithms)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(algorithms) <= 0 {
		t.Log("error : ", "there is no algorithm to update!")
		return
	}

	algorithms[0].Name = "algorithm-name-update"
	algorithms[0].Version = "v1-update"
	algorithms[0].Description = "this is algorithm-update"
	algorithms[0].Author = "author-update"

	// post create algorithm
	requestData, err := json.Marshal(&algorithms[0])
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", base_url, strings.NewReader(string(requestData)))

	res, err := client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
	}

	t.Log(string(resBody))
}

func Test_DeleteById(t *testing.T) {
	// get all
	res1, err := http.Get(base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res1.Body.Close()
	resBody1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody1))

	var response model.Response
	json.Unmarshal(resBody1, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var algorithms []*model.Algorithm
	json.Unmarshal(([]byte)(response.Result), &algorithms)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(algorithms) <= 0 {
		t.Log("error : ", "there is no algorithm to delete!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", base_url+"/id/"+strconv.FormatInt(algorithms[0].Id, 10), nil)

	res, err := client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}
