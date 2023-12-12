package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("bookhouse.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect bookhouse database")
	}

	database.AutoMigrate(
		&Admin{},
		&Member{},
		&Gender{},
		&Religion{},
		&Prefix{},
		&Room{},
		&Room_type{},
		&Room_status{},
		&Room_booking{},
		&Facilities{},
		&Rewards_history{},
		&Point_history{},
		&Rewards{},
		&ReCat{},
		&Donate_book{},
		&Book_request_status{},
		&Book_request{},
		&Catagory{},
		&Book{},
		&Status{},
		&Publisher{},
		&Borrow_book{},
		&Borrow_status{},
		&Review{},
	)
	db = database
	MockupAdmin()
	MockupRoom_status()
	MockupRoom_type()
	Mockup_Facilities()
	MockupGender()
	MockupReligion()
	MockupPrefix()
	MockupReward_cat()
	MockupBookStatus()
	MockupPublisher()
	MockupBookCat()
	//MockupBook_re_status()

}
