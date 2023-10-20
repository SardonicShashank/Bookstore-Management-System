package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err:=gorm.Open("mysql","root@tcp(localhost)/bookstore?parseTime=true")
	if err!=nil{
		panic(err)
	}else{
		log.Println("DataBase Connected")
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}


