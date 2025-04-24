package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(100);not null;index"`
	Age       int
	CreatedAt time.Time
	UpdateAt  time.Time
	UserInfo  UserInfo `gorm:"foreignKey:ID;references:ID"`
}

type UserInfo struct {
	ID        uint   `gorm:"primarykey"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	phone     int    `gorm:"type:varchar(20);unique;not null"`
	CreatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedAt = time.Now()
	return nil
}

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(200)
	sqlDB.SetMaxOpenConns(2000)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return db, nil

}

func main() {
	dsn := "root:mismis@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := InitDB(dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()
	// 建表
	db.AutoMigrate(&User{})
	// db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100), age INT)")

	// 插入数据
	db.Create(&User{Name: "zhi", Age: 21})

	// 查询数据
	var user User
	db.Where("id =?", 15).First(&user)
	db.Where("name = ?", "zhi").Find(&user)
	fmt.Printf("User: %v\n", user)

	// 更新数据
	db.Model(&user).Update("Age", 40)
	fmt.Printf("User: %v\n", user)

	// 删除数据
	// db.Delete(&user)
}
