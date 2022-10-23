package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Booking{},
		&Employee{},
		&User{},
		&FoodOrdered{},
		&PaymentType{},
		&Bill{},
	)
	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("45191"), 14)
	if err != nil {
		panic("failed")
	}

	db.Model(&User{}).Create(&User{
		Email: "moon@gmil.com",
		Name:     "Kiadtisak J",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Fei",
		Email: "man@gmil.com",
		Password: string(password),
	})

	var jo User
	var fei User

	db.Raw("SELECT * FROM users WHERE email = ?", "moon@gmil.com").Scan(&jo)
	db.Raw("SELECT * FROM users WHERE email = ?", "man@gmil.com").Scan(&fei)

	//Booking
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Komson",
		Password: string(password),
		Email:    "komson@gmai.com",
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "wichi",
		Password: string(password),
		Email:    "wichai@gmai.com",
	})

	 var a User
	 var b User

	 db.Raw("SELECT * FROM employees WHERE email = ?", "komson@gmai.com").Scan(&a)
	 db.Raw("SELECT * FROM employees WHERE email = ?", "wichai@gmai.com").Scan(&b)


	Booking1 := Booking{
		BookingTimeStart: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		BookingTimeStop: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		Room:		  "1012",
		TotalPrice:     3000,
		User :     	jo,
	}
	db.Model(&Booking{}).Create(&Booking1)

	Booking2 := Booking{
		BookingTimeStart: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		BookingTimeStop: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		Room:		  "1112",
		TotalPrice:     2000,
		User :     	fei,
	}
	db.Model(&Booking{}).Create(&Booking2)

	//FoodOrderes
	Foodordered1 := FoodOrdered{
		Name: "Set A",
		FoodTime: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
	TotalPrice: 500,
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered1)

	Foodordered2 := FoodOrdered{
		Name: "Set B",
		FoodTime: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
	TotalPrice: 1000,
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered2)

	Foodordered3 := FoodOrdered{
		Name: "Set C",
		FoodTime: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		TotalPrice: 1500,
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered3)
	//PaymentType
	PaymentType1 := PaymentType{
		Name: "เงินสด",
	}
	db.Model(&PaymentType{}).Create(&PaymentType1)
	PaymentType2 := PaymentType{
		Name: "ธนาคาร A",
	}
	db.Model(&PaymentType{}).Create(&PaymentType2)
	PaymentType3 := PaymentType{
		Name: "ธนาคาร B",
	}
	db.Model(&PaymentType{}).Create(&PaymentType3)

	//Bill1
	/*db.Model(&Bill{}).Create(&Bill{

		Booking: Booking1,
		FoodOrdered:        Place1,
		PaymentType:  PaymentType1,
		BillTime:     time.Now(),
		TotalPrice:   (Booking1.TotalCredit * 800),
	})*/

	/*//Bill2
	db.Model(&Bill{}).Create(&Bill{

		Booking: Booking2,
		FoodOrdered:        Place2,
		PaymentType:  PaymentType2,
		BillTime:     time.Now(),
		TotalPrice:   (Booking2.TotalCredit * 800),
	})

	//Bill3
	db.Model(&Bill{}).Create(&Bill{

		Booking: Booking3,
		FoodOrdered:        Place3,
		PaymentType:  PaymentType3,
		BillTime:     time.Now(),
		TotalPrice:   (Booking2.TotalCredit * 800),
	})*/

}
