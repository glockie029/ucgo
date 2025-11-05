package structpractice

import (
	"fmt"
	"time"
)

type Address struct {
	Street  string
	City    string
	ZipCode string
}
type VCard struct {
	Name        string
	AddressCode *Address
	Birthday    time.Time
	Picture     []byte //图像地址
}

func vcard() {
	homeaddress := Address{
		Street:  "XiangYang",
		City:    "China",
		ZipCode: "123456",
	}
	birthday := time.Date(2000, time.May, 1, 0, 0, 0, 0, time.UTC)
	myImage := []byte{0x89, 0x50, 0x4e, 0x47}
	myVcard := VCard{
		Name:        "s1ldyd",
		AddressCode: &homeaddress,
		Birthday:    birthday,
		Picture:     myImage,
	}
	fmt.Println(myVcard)
}
