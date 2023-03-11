package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

type UserBasic struct {
	gorm.Model
	Name       string `json:"name" gorm:"type:varchar(11);not null;default:'';comment:姓名"`
	Password   string `json:"password" gorm:"type:varchar(32);not null;default:'';comment:密码"`
	Phone      string `json:"phone" gorm:"type:varchar(32);not null;default:'';comment:手机号"`
	Email      string
	Identity   string
	ClientIp   string
	ClientPort string
}

func TestGormTest(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "root:123456@tcp(127.0.0.1:3306)/jyoa?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "as_",
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	_ = db.AutoMigrate(&UserBasic{})
}
