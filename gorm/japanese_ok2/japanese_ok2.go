package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// ビジネスコミック2 is comics table record struct.
type ビジネスコミック2 struct {
	ID         int64  `json:"id" gorm:"primary_key; column:id"  `
	Sタイトル      string `json:"title" gorm:"column:タイトル"`
	Iいいね       int64  `json:"good" gorm:"column:いいね"`
	S外部リンクURL  string `json:"outer_link_url" gorm:"column:外部リンクURL"`
	S外部イメージURL string `json:"outer_image_url" gorm:"column:外部イメージURL"`
	S要約        string `json:"summary" gorm:"column:要約"`
}

// TableName はビジネスコミック2構造体に対応するテーブル名を返します
func (m *ビジネスコミック2) TableName() string {
	return "ビジネスコミック2"
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
	db.AutoMigrate(&ビジネスコミック2{})
	var 本 = ビジネスコミック2{ID: 1, Sタイトル: "マンガでわかる回帰分析", Iいいね: 5, S外部リンクURL: "", S外部イメージURL: "", S要約: ""}
	fmt.Println("本 : ", 本)
	mTx := db.Begin()
	defer mTx.Close()
	本配列 := []ビジネスコミック2{}
	//fmt.Println("comic2 : ", comic2)
	//	mTx.First(&comic2, 2)
	mTx.Find(&本配列, "id = ?", 本.ID)
	fmt.Println("comics : ", 本配列)
	if 0 < len(本配列) {
		fmt.Println("存在するレコードです。")
		// レコードが存在する場合 updateする。
		本更新後情報 := 本配列[0]
		本更新後情報.Iいいね = 本更新後情報.Iいいね + 1
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
