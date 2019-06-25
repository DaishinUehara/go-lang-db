package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Comic is comics table record struct.
type Comic struct {
	ID            int64  `json:"id" gorm:"primary_key; column:id"  `
	Title         string `json:"title" gorm:"column:タイトル"`
	Good          int64  `json:"good" gorm:"column:いいね"`
	OuterLinkURL  string `json:"outer_link_url" gorm:"column:外部リンクURL"`
	OuterImageURL string `json:"outer_image_url" gorm:"column:外部イメージURL"`
	Summary       string `json:"summary" gorm:"column:要約"`
}

// TableName is interface of return table name of Comic stract.
func (m *Comic) TableName() string {
	return "ビジネスコミック"
}

func main() {
	// COMIC_DBCONNECTION="user=username password=pw dbname=testdb sslmode=disable"
	dbconnection := os.Getenv("COMIC_DBCONNECTION")
	db, err1 := gorm.Open("postgres", dbconnection)
	if err1 != nil {
		fmt.Println("err1 : ", err1)
		panic(err1)
	}
	db.Set("gorm:table_options", "charset=utf8")
	db.LogMode(true)
	defer db.Close()
	db.AutoMigrate(&Comic{})
	var 本 = Comic{ID: 1, Title: "マンガでわかる回帰分析", Good: 5, OuterLinkURL: "", OuterImageURL: "", Summary: ""}
	fmt.Println("本 : ", 本)
	mTx := db.Begin()
	defer mTx.Close()
	本配列 := []Comic{}
	//fmt.Println("comic2 : ", comic2)
	//	mTx.First(&comic2, 2)
	mTx.Find(&本配列, "id = ?", 本.ID)
	fmt.Println("comics : ", 本配列)
	if 0 < len(本配列) {
		fmt.Println("存在するレコードです。")
		// レコードが存在する場合 updateする。
		本更新後情報 := 本配列[0]
		本更新後情報.Good = 本更新後情報.Good + 1
		fmt.Println("comicAfter : ", 本更新後情報)
		mTx.Model(&本配列[0]).Update(&本更新後情報)
		mTx.Commit()
	} else {
		fmt.Println("存在ないレコードです。")
		// レコードが存在しない場合
		fmt.Println("comic : ", 本)
		err4 := mTx.Create(&本).Error
		if err4 != nil {
			fmt.Println("err4 : ", err4)
			mTx.Rollback()
		} else {
			mTx.Commit()
		}
	}

}
