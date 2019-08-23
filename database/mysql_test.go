package database

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModel_CreateReceiptFail(t *testing.T) {
	mdb, mock, err := sqlmock.New()
	require.Nil(t, err)

	db, err := gorm.Open("mysql", mdb)
	require.Nil(t, err)
	defer mdb.Close()

	model := Model{
		One: "one",
		Two: 2,
	}
	// 1. insert ok
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT  INTO `model` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = SaveModel(db, &model)
	if !assert.Nil(t, err) {
		log.Println("正常保存预期外的错误: ", err.Error())
	}
	// 2. 测试插入失败
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT  INTO `model` (.+)$").WillReturnError(errors.New("fail"))
	mock.ExpectCommit()

	err = SaveModel(db, &model)
	if assert.NotNil(t, err) {
		assert.Equal(t, "fail", errors.Cause(err).Error(), "返回错误消息必须是fail")
		log.Println(err.Error())
	}
}
