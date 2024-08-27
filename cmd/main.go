package main

import (
	"golang-gorm/internal/delivery/server"
)

func main() {

	server.NewServer().Run()

	//fmt.Println("Hello World")
	//
	//dsn := "user=this_user password=password dbname=golang_gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//db.AutoMigrate(&domain.User{})
	//
	////result := db.Create(&domain.User{
	////	ID:        "001",
	////	Username:  "test",
	////	Email:     "test@test.com",
	////	Password:  "123456",
	////	Role:      "admin",
	////	CreatedAt: time.Time{},
	////	UpdatedAt: time.Time{},
	////})
	//
	////fmt.Println(result.RowsAffected)
	//
	//var user domain.User
	//result := db.First(&user, "0011")
	//if result.RowsAffected != 1 {
	//	fmt.Println("user with ID = BLA not found")
	//}
	//fmt.Println(user)
}
