package main

import (
	"fmt"
	"log"
	"time"

	"github.com/k-p5w/go-stockman/include/xmlreader"
	"github.com/sclevine/agouti"
)

func main() {
	fmt.Println("start.")
	fname := "book.xml"
	data, alldata := xmlreader.ReadXML(fname)
	// 出版社別のcsvファイルを作成する
	xmlreader.CreateCSV("出版社別サマリ", data)

	myGoblin(alldata)
}

// myGoblin is メッセージＸＭＬをinput.xmlに変換して投入する
func myGoblin(items []string) {

	// クロームの場合
	agoutiDriver := agouti.ChromeDriver(agouti.Browser("chrome"))

	if err := agoutiDriver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer agoutiDriver.Stop()

	page, err := agoutiDriver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	// page, _ := agoutiDriver.NewPage()
	openurl := "https://tsundoku.site/login"
	// URLの表示
	// page.Navigate(loginurl)
	// ログインページに遷移
	if err := page.Navigate(openurl); err != nil {
		log.Fatalf("URL展開に失敗:%v", err)
	}

	// 表示を待つ
	time.Sleep(10 * time.Second)
	// ログインする
	page.FindByXPath(`//*[@id="firebaseui-auth-container"]/div/div[1]/form/ul/li[1]/button`).Click()
	time.Sleep(10 * time.Second)
	page.FindByID("username_or_email").Fill("twitter_id")
	page.FindByID("password").Fill("twitter_password")

	// submitする
	page.FindByButton("ログイン").Submit()
	time.Sleep(10 * time.Second)

	goFind := func() {
		// Amazon本を選択する
		openurl = "https://tsundoku.site/search"
		if err := page.Navigate(openurl); err != nil {
			log.Fatalf("URL展開に失敗:%v", err)
		}

		time.Sleep(3 * time.Second)
		// 検索する
		page.FindByXPath(`//*[@id="main-content"]/div/div[1]/div[1]/div[1]/div[1]/button`).Click()
		time.Sleep(3 * time.Second)
		//
		page.FindByXPath(`//*[@id="main-content"]/div/div[1]/div[1]/div[1]/div[3]/div/a[3]`).Click()

	}

	// 対象数分繰り返す
	for key, val := range items {
		if len(val) == 0 {
			continue
		}
		goFind()
		fmt.Printf("%v:%v　\n", key, val)

		// 検索欄に値を入力
		page.FindByXPath(`//*[@id="main-content"]/div/div[1]/div[2]/div[1]/p[1]/input`).Fill(val)
		// 検索する
		page.FindByXPath(`//*[@id="main-content"]/div/div[1]/div[2]/div[1]/p[2]/button`).Click()
		time.Sleep(7 * time.Second)
		// これを積む押下
		page.FindByXPath(`//*[@id="main-content"]/div/div[2]/div[2]/nav/div[1]/div[4]/div[2]/div[1]/div/p/a`).Click()
		time.Sleep(2 * time.Second)
	}

	log.Printf("end..")

}
