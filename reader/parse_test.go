package reader

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	str := `{"name":"abt.xlsx","sheets":[{"name":"Sheet1","titles":[{"name":"testA","style":null},{"name":"testB","style":null}],"rows":[{"columns":["t1","t2"]},{"columns":["T1","T2"]}]}]}`

	fileData, e := ParseFileData(str)
	if e != nil {
		t.Error(e)
	}
	if fileData.Name == "abt.xlsx" && fileData.Sheets[0].Name == "Sheet1" && fileData.Sheets[0].Titles[0].Name == "testA" && fileData.Sheets[0].Rows[0].Columns[0] == "T1" {
		t.Log("TestParse success")
	} else {
		t.Error()
	}
	fmt.Println(fileData)
}
