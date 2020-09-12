package xmlreader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const filetimestampformat = "20060102"

// CreateCSV is CSV作成用
func CreateCSV(val string, d map[string]int) {

	setName := func(filename string) string {
		dirname := "./csv/"

		t := time.Now()
		filename = fmt.Sprintf("%v_%v.csv", t.Format(filetimestampformat), filename)
		return filepath.Join(dirname, filename)
	}

	for maker, cnt := range d {
		// fmt.Printf("%v-%v[%v]\n", k, v, key)

		fullpath := setName(val)

		//書き込みファイル作成
		// file, err := os.Create(fullpath)
		file, err := os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		// WriteWrapper := func(w *csv.Writer, record []string) error {
		// 	for i, field := range record {
		// 		record[i] = `"` + field + `"`
		// 	}
		// 	return w.Write(record)
		// }

		writer := csv.NewWriter(transform.NewWriter(file, japanese.ShiftJIS.NewEncoder()))

		//　デリミタの設定
		writer.Comma = ','

		// "Subject","Start Date","Start Time","End Date","End Time","All Day Event","Description","Location","Private"

		csvdata := make([]string, 2)
		csvdata[0] = maker
		csvdata[1] = fmt.Sprintf("%v冊", cnt)
		// WriteWrapper(writer, csvdata)
		writer.Write(csvdata)
		writer.Flush()
	}

}
