package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Comic is comics table record struct.
type Comic struct {
	ID            int64  `gorm:"primary_key" json:"id"`
	Title         string `json:"title"`
	Star          int64  `json:"star"`
	OuterLinkURL  string `json:"outer_link_url"`
	OuterImageURL string `json:"outer_image_url"`
	Summary       string `json:"summary"`
}

// Comics is table, has comic records.
type Comics []Comic

func main() {
	// COMIC_DBCONNECTION="user=username password=pw dbname=testdb sslmode=disable"
	dbconnection := os.Getenv("COMIC_DBCONNECTION")
	db, err1 := gorm.Open("postgres", dbconnection)
	if err1 != nil {
		fmt.Println("err1 : ", err1)
		panic(err1)
	}
	db.LogMode(true)
	defer db.Close()
	db.AutoMigrate(&Comic{})
	var comic = Comic{ID: 1, Title: "マンガでわかる回帰分析", Star: 5, OuterLinkURL: "", OuterImageURL: "", Summary: ""}
	fmt.Println("comic : ", comic)
	mTx := db.Begin()
	defer mTx.Close()
	comics := []Comic{}
	//fmt.Println("comic2 : ", comic2)
	//	mTx.First(&comic2, 2)
	mTx.Find(&comics, "id = ?", comic.ID)
	fmt.Println("comics : ", comics)
	if 0 < len(comics) {
		fmt.Println("存在するレコードです。")
		// レコードが存在する場合 updateする。
		comicAfter := comics[0]
		fmt.Println("comicAfter : ", comicAfter)
		comicAfter.Star = comicAfter.Star + 1
		mTx.Model(&comics[0]).Update(&comicAfter)
		mTx.Commit()
	} else {
		fmt.Println("存在ないレコードです。")
		// レコードが存在しない場合
		fmt.Println("comic : ", comic)
		err4 := mTx.Create(&comic).Error
		if err4 != nil {
			fmt.Println("err4 : ", err4)
			mTx.Rollback()
		} else {
			mTx.Commit()
		}
	}
	//fmt.Println("comic2 : ", comic2)
	/*
		err4 := mTx.Create(&comic).Error
		if err4 != nil {
			fmt.Println("err4 : ", err4)
			mTx.Rollback()
		} else {
			mTx.Commit()
		}
	*/
	//	fmt.Printf("Hello world\n")

}
