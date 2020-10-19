package reader

import "encoding/json"

type RequestInterface interface {
}
type Request struct {
	File *FileData
}

type FileData struct {
	Sheets []*Sheet `json:"sheets"`
	Name   string   `json:"name"`
}

type Sheet struct {
	Name   string   `json:"name"`
	Titles []*Title `json:"titles"`
	Rows   []*Rows  `json:"rows"`
}

type Title struct {
	Name  string `json:"name"`
	Style *Style `json:"style"`
}

type Rows struct {
	Columns []Column `json:"columns"`
}

type Column string

type Style struct {
	//	对齐方式
	Align string `json:"align"`
	//	行宽
	RowWidth float64 `json:"row_width"`
	//	行高
	RowHeight float64 `json:"row_height"`
}

func ParseFileData(jsonStr string) (file *FileData, e error) {
	file = &FileData{}
	e = json.Unmarshal([]byte(jsonStr), file)
	if e != nil {
		return nil, e
	}
	return file, e
}
