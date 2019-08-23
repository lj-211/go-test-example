package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Model struct {
	gorm.Model

	One string
	Two int32
}

func (self *Model) TableName() string {
	return "model"
}

// 创建凭据
func SaveModel(db *gorm.DB, model *Model) error {
	if db == nil {
		return errors.Errorf("db实例为空")
	}
	if model == nil {
		return errors.Errorf("数据为空")
	}
	if cerr := db.Create(model).Error; cerr != nil {
		return errors.Wrap(cerr, "保存错误")
	}

	return nil
}
