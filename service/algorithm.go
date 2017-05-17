package service

import (
	"fmt"
	"model"

	"dao-service/algorithm-dao-service/dao"

	"github.com/astaxie/beego"
)

type AlgorithmService struct {
}

func (this *AlgorithmService) Create(algorithm *model.Algorithm) error {
	var err error
	var algorithmDao = dao.NewAlgorithmDao()

	err = algorithmDao.Create(algorithm)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create algorithm failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *AlgorithmService) GetAll() ([]*model.Algorithm, error) {
	var err error
	var algorithmDao = dao.NewAlgorithmDao()
	var algorithms []*model.Algorithm

	// get algorithms
	algorithms, err = algorithmDao.GetAll()
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	return algorithms, err
}

func (this *AlgorithmService) GetByNameAndVersion(name, version string) (*model.Algorithm, error) {
	var err error
	var algorithmDao = dao.NewAlgorithmDao()
	var algorithm *model.Algorithm

	// get algorithm
	algorithm, err = algorithmDao.GetByNameAndVersion(name, version)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	return algorithm, err
}

func (this *AlgorithmService) Update(algorithm *model.Algorithm) error {
	var err error
	var algorithmDao = dao.NewAlgorithmDao()

	err = algorithmDao.Update(algorithm)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update algorithm failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *AlgorithmService) DeleteById(id int64) error {
	var err error
	var algorithmDao = dao.NewAlgorithmDao()

	// delete algorithm
	err = algorithmDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete algorithm failed!", "reason:"+err.Error())
		return err
	}

	return err
}
