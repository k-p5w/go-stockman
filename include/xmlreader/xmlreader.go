package xmlreader

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

// Xmlnode is XMLの読み込みのための定義(ns1)
type Xmlnode struct {
	Response      string `xml:"response"`
	Addupdatelist struct {
		Metadata []Metadata `xml:"meta_data"`
	} `xml:"add_update_list"`
}

// Metadata is メタデータ
type Metadata struct {
	Asin       string `xml:"ASIN"`
	Publishers struct {
		Publisher string `xml:"publisher"`
	} `xml:"publishers"`
}

// `xml:"meta_data"`

// ReadXML is XMLの読み込み
func ReadXML(xmlfile string) (map[string]int, []string) {

	var nodedata Xmlnode

	// XMLファイルの読み込み
	rawData, err := ioutil.ReadFile(xmlfile)
	if err != nil {
		log.Fatal(err)
	}
	// 変換 for NS1
	err = xml.Unmarshal(rawData, &nodedata)
	if err != nil {
		log.Fatal(err)
	}

	bookid := make([]string, 5000)
	publishersMap := make(map[string]int, 0)

	for _, v := range nodedata.Addupdatelist.Metadata {
		// fmt.Printf("%v:%v (%v) \n", k, v.Asin, v.Publishers.Publisher)
		publishersMap[v.Publishers.Publisher]++
		bookid = append(bookid, v.Asin)
	}
	return publishersMap, bookid
}
