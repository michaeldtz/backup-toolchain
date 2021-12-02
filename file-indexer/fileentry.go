package main

import (
	"encoding/json"
	"fmt"
)

type FileEntry struct {
	Id       string
	Filename string
	Date     int64
	Size     int64
	Path     string
	Dirname  string
	Hash     []byte
}

func (fe *FileEntry) ToCSVString() string {
	return fmt.Sprintf("%s,%s,%s,%d,%d,%x\n", fe.Filename, fe.Dirname, fe.Path, fe.Date, fe.Size, fe.Hash)
}

func (fe *FileEntry) ToJSON(id string) []byte {
	fe.Id = id
	json, err := json.Marshal(fe)
	check(err)
	return json
}
