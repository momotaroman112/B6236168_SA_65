package entity

import (
	"time"

	"gorm.io/gorm"
)

//การกำหนด schema
type User struct {
	gorm.Model
	Name string
	Email		string			`gorm:"uniqueIndex"`
	Password     string

	Bookings	[]Booking	 `gorm:"foreignKey:UserID"`
}

type Employee struct {
	gorm.Model
	Name string
	//StudentCode
	Email		string			`gorm:"uniqueIndex"`
	Password     string

	Bills	[]Bill	 `gorm:"foreignKey:EmployeeID"`
}

type FoodOrdered struct {
	gorm.Model
	Name string
	FoodTime time.Time
	TotalPrice     uint
}

type PaymentType struct {
	gorm.Model
	Name string
	Bill []Bill `gorm:"foreignKey:PaymentTypeID"`
}


type Booking struct {
	gorm.Model
	BookingTimeStart  time.Time
	BookingTimeStop time.Time
	Room		  string
	TotalPrice     uint

	UserID     *uint
	User      	User
}


type Bill struct {
	gorm.Model
	
	BillTime time.Time

	EmployeeID     *uint
	Employee        Employee

	PaymentTypeID *uint
	PaymentType   PaymentType

	BookingID *uint    `gorm:"uniqueIndex"`
	Booking    Booking `gorm:"constraint:OnDelete:CASCADE"` //belong to ลบใบลงทะเบียน บิลหาย

	FoodOrderedID *uint
	FoodOrdered   FoodOrdered

	TotalPrice uint
}
