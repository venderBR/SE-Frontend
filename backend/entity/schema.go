package entity

import (
	"gorm.io/gorm"

	"time"
)

// ------------------------------------------------------------------------
type Admin struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string

	Book         []Book         `gorm:"foreignKey:Admin_id"`
	Rewards      []Rewards      `gorm:"foreignKey:Admin_id"`
	Room         []Room         `gorm:"foreignKey:Admin_id"`
	Borrow_book  []Borrow_book  `gorm:"foreignKey:Admin_id"`
	Book_request []Book_request `gorm:"foreignKey:Admin_id"`
	Donate_book  []Donate_book  `gorm:"foreignKey:Admin_id"`
}

// ------------------------------------------------------------------------
type Member struct {
	gorm.Model
	FirstName string
	LastName  string
	Tel       string
	Age       int
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Point     float32

	Review          []Review          `gorm:"foreignKey:Member_id"`
	Borrow_book     []Borrow_book     `gorm:"foreignKey:Member_id"`
	Rewards_history []Rewards_history `gorm:"foreignKey:Member_id"`
	Donate_book     []Donate_book     `gorm:"foreignKey:Member_id"`
	Point_history   []Point_history   `gorm:"foreignKey:Member_id"`
	Room_booking    []Room_booking    `gorm:"foreignKey:Member_id"`
	Book_request    []Book_request    `gorm:"foreignKey:Member_id"`

	Gender_id *uint  //fk
	Gender    Gender `gorm:"references:id"`

	Religion_id *uint    //fk
	Religion    Religion `gorm:"references:id"`

	Prefix_id *uint  //fk
	Prefix    Prefix `gorm:"references:id"`
}
type Gender struct {
	gorm.Model
	Gender string

	Member []Member `gorm:"foreignKey:Gender_id"`
}
type Religion struct {
	gorm.Model
	Religion string

	Member []Member `gorm:"foreignKey:Religion_id"`
}
type Prefix struct {
	gorm.Model
	Prefix string

	Member []Member `gorm:"foreignKey:Prefix_id"`
}

// -------------------------------------------------------------------------
type Room struct {
	gorm.Model

	Room_name      string
	Start_time     time.Time
	End_time       time.Time
	Room_image_src string

	RoomType_id *uint     //fk
	Room_type   Room_type `gorm:"references:id"`

	Room_status_id *uint       //fk
	Room_status    Room_status `gorm:"references:id"`

	Facilities_id *uint      //fk
	Facilities    Facilities `gorm:"references:id"`

	Admin_id *uint
	Admin    Admin `gorm:"references:id"`

	Room_booking []Room_booking `gorm:"foreignKey:Room_id"`
}

type Room_type struct {
	gorm.Model
	Name     string
	Capacity int

	Room []Room `gorm:"foreignKey:RoomType_id"`
}

type Room_status struct {
	gorm.Model
	Status_name string

	Room []Room `gorm:"foreignKey:Room_status_id"`
}
type Room_booking struct {
	gorm.Model

	Start_time time.Time
	End_time   time.Time

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Room_id *uint //fk
	Room    Room  `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Room_booking_id"`
}
type Facilities struct {
	gorm.Model
	Facilities string

	Room []Room `gorm:"foreignKey:Facilities_id"`
}

// -------------------------------------------------------------------------
type Rewards_history struct {
	gorm.Model
	// Rewardshis_date time.Time

	Rewards_id *uint   //fk
	Rewards    Rewards `gorm:"references:id"`

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`
}
type Point_history struct {
	gorm.Model
	Point  int
	Remark string

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Borrow_book_id *uint       //fk
	Borrow_book    Borrow_book `gorm:"references:id"`

	Room_booking_id *uint        //fk
	Room_booking    Room_booking `gorm:"references:id"`

	Donate_book_id *uint       //fk
	Donate_book    Donate_book `gorm:"references:id"`
}

// -------------------------------------------------------------------------
type Rewards struct {
	gorm.Model

	Rewards_picture   string
	Rewards_title     string
	Rewards_details   string
	Point_to_paid     int
	Rewards_stock     int
	Rewards_image_src string

	Rewards_history []Rewards_history `gorm:"foreignKey:Rewards_id"`

	Re_cat_id *uint // fk
	ReCat     ReCat `gorm:"references:id"`

	Admin_id *uint
	Admin    Admin `gorm:"references:id"`
}
type ReCat struct {
	gorm.Model
	Re_cat_name string

	Rewards []Rewards `gorm:"foreignKey:Re_cat_id"`
}

// -------------------------------------------------------------------------
type Donate_book struct {
	gorm.Model
	DonateBook_title     string
	Donate_point         int
	BookDonate_writer    string
	BookDonate_publisher string
	// DonateDate           time.Time

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Catagory_id *uint    //fk
	Catagory    Catagory `gorm:"references:id"`

	Admin_id *uint //fk
	Admin    Admin `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Donate_book_id"`
}

// -------------------------------------------------------------------------
type Book_request_status struct {
	gorm.Model
	Status_name string

	Book_request []Book_request `gorm:"foreignKey:Book_request_status_id"`
}
type Book_request struct {
	gorm.Model
	BookRequest_title  string
	BookRequest_reason string
	// BookRequestDate time.Time
	BookRequest_writer    string
	BookRequest_publisher string

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Catagory_id *uint    //fk
	Catagory    Catagory `gorm:"references:id"`

	Admin_id *uint //fk
	Admin    Admin `gorm:"references:id"`

	Book_request_status_id *uint               //fk
	Book_request_status    Book_request_status `gorm:"references:id"`
}

// -------------------------------------------------------------------------
type Catagory struct {
	gorm.Model
	Cat_name string

	Book         []Book         `gorm:"foreignKey:Catagory_id"`
	Book_request []Book_request `gorm:"foreignKey:Catagory_id"`
	Donate_book  []Donate_book  `gorm:"foreignKey:Catagory_id"`
}

// -------------------------------------------------------------------------
type Book struct {
	gorm.Model
	Book_title       string
	QuantityInStock  int
	QuantityInRental int
	Image_src        string
	TotalBook        int

	Review      []Review      `gorm:"foreignKey:Book_id"`
	Borrow_book []Borrow_book `gorm:"foreignKey:Book_id"`

	Catagory_id *uint
	Catagory    Catagory `gorm:"references:id"`

	Publisher_id *uint
	Publisher    Publisher `gorm:"references:id"`

	Status_id *uint
	Status    Status `gorm:"references:id"`

	Admin_id *uint
	Admin    Admin `gorm:"references:id"`
}

type Status struct {
	gorm.Model
	Status string

	Book []Book `gorm:"foreignKey:Status_id"`
}
type Publisher struct {
	gorm.Model
	Publisher string

	Book []Book `gorm:"foreignKey:Publisher_id"`
}

// -------------------------------------------------------------------------
type Borrow_book struct {
	gorm.Model

	Start_time time.Time
	End_time   time.Time
	Totalbook  int

	Borrow_status_id int           //fk
	Borrow_status    Borrow_status `gorm:"references:id"`

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Book_id *uint //fk
	Book    Book  `gorm:"references:id"`

	Admin_id *uint
	Admin    Admin `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Borrow_book_id"`
}
type Borrow_status struct {
	gorm.Model
	Name string

	Borrow_book []Borrow_book `gorm:"foreignKey:Borrow_status_id"`
}

// -------------------------------------------------------------------------
type Review struct {
	gorm.Model
	Rate    string
	Comment string
	// ReviewDate time.Time

	Book_id *uint //fk
	Book    Book  `gorm:"references:id"`

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`
}

// -------------------------------------------------------------------------
