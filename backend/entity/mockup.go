package entity

func MockupAdmin() {
	admin1 := Admin{
		Email:    "admin@gmail.com",
		Password: "admin1234",
	}
	db.Model(&Admin{}).Create(&admin1)
}

// --------------Room-----------------------
func MockupRoom_type() {
	Rtype1 := Room_type{
		Name:     "small",
		Capacity: 5,
	}
	db.Model(&Room_type{}).Create(&Rtype1)
	Rtype2 := Room_type{
		Name:     "medium",
		Capacity: 10,
	}
	db.Model(&Room_type{}).Create(&Rtype2)
	Rtype3 := Room_type{
		Name:     "large",
		Capacity: 20,
	}
	db.Model(&Room_type{}).Create(&Rtype3)
}
func MockupRoom_status() {
	Rstatus1 := Room_status{
		Status_name: "ไม่ว่าง",
	}
	db.Model(&Room_status{}).Create(&Rstatus1)

	Rstatus2 := Room_status{
		Status_name: "กำลังทำความสะอาด",
	}
	db.Model(&Room_status{}).Create(&Rstatus2)

	Rstatus3 := Room_status{
		Status_name: "ว่าง",
	}
	db.Model(&Room_status{}).Create(&Rstatus3)
}
func Mockup_Facilities() {
	Facilities1 := Facilities{
		Facilities: "โปรเจคเตอร์",
	}
	db.Model(&Facilities{}).Create(&Facilities1)

	Facilities2 := Facilities{
		Facilities: "ทีวี",
	}
	db.Model(&Facilities{}).Create(&Facilities2)

	Facilities3 := Facilities{
		Facilities: "ปลั๊กพ่วง",
	}
	db.Model(&Facilities{}).Create(&Facilities3)

	Facilities4 := Facilities{
		Facilities: "ปากกาไวท์บอร์ด",
	}
	db.Model(&Facilities{}).Create(&Facilities4)

}

// -----------------------Member-----------------------------
func MockupGender() {
	Gender1 := Gender{
		Gender: "Male",
	}
	db.Model(&Gender{}).Create(&Gender1)

	Gender2 := Gender{
		Gender: "Female",
	}
	db.Model(&Gender{}).Create(&Gender2)

	Gender3 := Gender{
		Gender: "unspecified",
	}
	db.Model(&Gender{}).Create(&Gender3)
}
func MockupReligion() {
	Religion1 := Religion{
		Religion: "พุธ",
	}
	db.Model(&Religion{}).Create(&Religion1)

	Religion2 := Religion{
		Religion: "คริสต์",
	}
	db.Model(&Religion{}).Create(&Religion2)

	Religion3 := Religion{
		Religion: "อิสลาม",
	}
	db.Model(&Religion{}).Create(&Religion3)

	Religion4 := Religion{
		Religion: "เบคอน",
	}
	db.Model(&Religion{}).Create(&Religion4)

	Religion5 := Religion{
		Religion: "ไม่ระบุศาสนา",
	}
	db.Model(&Religion{}).Create(&Religion5)
}
func MockupPrefix() {
	Prefix1 := Prefix{
		Prefix: "นาย",
	}
	db.Model(&Prefix{}).Create(&Prefix1)

	Prefix2 := Prefix{
		Prefix: "นาง",
	}
	db.Model(&Prefix{}).Create(&Prefix2)

	Prefix3 := Prefix{
		Prefix: "นางสาว",
	}
	db.Model(&Prefix{}).Create(&Prefix3)

	Prefix4 := Prefix{
		Prefix: "ไม่ระบุ",
	}
	db.Model(&Prefix{}).Create(&Prefix4)
}

// --------------------Reward--------------------------------------------
func MockupReward_cat() {
	ReCat1 := ReCat{
		Re_cat_name: "เครื่องเขียน",
	}
	db.Model(&ReCat{}).Create(&ReCat1)

	ReCat2 := ReCat{
		Re_cat_name: "เครื่องใช้ในบ้าน",
	}
	db.Model(&ReCat{}).Create(&ReCat2)

	ReCat3 := ReCat{
		Re_cat_name: "เครื่องใช้ในครัว",
	}
	db.Model(&ReCat{}).Create(&ReCat3)

	ReCat4 := ReCat{
		Re_cat_name: "เครื่องใช้ในทำความสะอาด",
	}
	db.Model(&ReCat{}).Create(&ReCat4)

	ReCat5 := ReCat{
		Re_cat_name: "เครื่องมือและเครื่องใช้ต่างๆ",
	}
	db.Model(&ReCat{}).Create(&ReCat5)

}

// ----------------------Book--------------------------
func MockupBookStatus() {
	Bstatus1 := Status{
		Status: "ยืมได้",
	}
	db.Model(&Status{}).Create(&Bstatus1)

	Bstatus2 := Status{
		Status: "ยืมไม่ได้",
	}
	db.Model(&Status{}).Create(&Bstatus2)
}
func MockupPublisher() {
	Pub1 := Publisher{
		Publisher: "SalmonBook",
	}
	db.Model(&Publisher{}).Create(&Pub1)

	Pub2 := Publisher{
		Publisher: "Abook",
	}
	db.Model(&Publisher{}).Create(&Pub2)

	Pub3 := Publisher{
		Publisher: "AvocadoBook",
	}
	db.Model(&Publisher{}).Create(&Pub3)

	Pub4 := Publisher{
		Publisher: "SiaminterComic",
	}
	db.Model(&Publisher{}).Create(&Pub4)
}
func MockupBookCat() {
	BookCat1 := Catagory{
		Cat_name: "สืบสวนสอบสวน",
	}
	db.Model(&Catagory{}).Create(&BookCat1)

	BookCat2 := Catagory{
		Cat_name: "โรแมนติก",
	}
	db.Model(&Catagory{}).Create(&BookCat2)

	BookCat3 := Catagory{
		Cat_name: "จิตวิทยา",
	}
	db.Model(&Catagory{}).Create(&BookCat3)

	BookCat4 := Catagory{
		Cat_name: "มังงะ",
	}
	db.Model(&Catagory{}).Create(&BookCat4)

	BookCat5 := Catagory{
		Cat_name: "พัฒนาตนเอง",
	}
	db.Model(&Catagory{}).Create(&BookCat5)

	BookCat6 := Catagory{
		Cat_name: "การเงินการลงทุน",
	}
	db.Model(&Catagory{}).Create(&BookCat6)

	BookCat7 := Catagory{
		Cat_name: "สยองขวัญ",
	}
	db.Model(&Catagory{}).Create(&BookCat7)
}
func MockupBook_re_status() {
	B_re_status1 := Book_request_status{
		Status_name: "accept",
	}
	db.Model(&Book_request_status{}).Create(B_re_status1)

	B_re_status2 := Book_request_status{
		Status_name: "unaccept",
	}
	db.Model(&Book_request_status{}).Create(B_re_status2)
}
