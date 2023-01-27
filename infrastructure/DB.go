package infrastructure

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func NewTestDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName:   c.DB.Test.DBName,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open(mysql.Open(d.Username+":"+d.Password+"@tcp(db-api:3306)/"+d.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(errors.DB0001)
	}
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}

func (db *DB) DBInit() {
	DBEngine := db.Connect()
	var err error
	err = errors.Combine(err, DBEngine.AutoMigrate(&models.UserEntity{}))
	err = errors.Combine(err, DBEngine.AutoMigrate(&models.BlogEntity{}))
	err = errors.Combine(err, DBEngine.AutoMigrate(&models.CategoryEntity{}))
	if err != nil {
		panic(errors.DB0002)
	}

}
