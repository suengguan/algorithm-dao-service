package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"config"
	"model"
)

type AlgorithmDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Algorithm
}

func NewAlgorithmDao() *AlgorithmDao {
	d := new(AlgorithmDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(config.DB_NAME)

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *AlgorithmDao) Create(algorithm *model.Algorithm) error {
	num, err := this.m_Orm.Insert(algorithm)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *AlgorithmDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no algorithm to delete")
		return err
	}

	return err
}

// update
func (this *AlgorithmDao) Update(algorithm *model.Algorithm) error {
	num, err := this.m_Orm.Update(algorithm)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no algorithm to update")
	}

	return err
}

// find
func (this *AlgorithmDao) GetAll() ([]*model.Algorithm, error) {
	var algorithms []*model.Algorithm

	num, err := this.m_QuerySeter.All(&algorithms)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return algorithms, err
}

func (this *AlgorithmDao) GetById(Id int64) (*model.Algorithm, error) {
	var algorithm model.Algorithm

	err := this.m_QuerySeter.Filter("ID", Id).One(&algorithm)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &algorithm, err
}

func (this *AlgorithmDao) GetByNameAndVersion(name string, version string) (*model.Algorithm, error) {
	var algorithm model.Algorithm

	err := this.m_QuerySeter.Filter("NAME", name).Filter("VERSION", version).One(&algorithm)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &algorithm, err
}
