package main

import (
	"encoding/json"
	"fmt"
)

type FileEntry struct {
	ScanId   string
	FileId   string
	Filename string
	IsDir    bool
	Date     int64
	Size     int64
	Path     string
	Dirname  string
	Hash     []byte
}

func (fe *FileEntry) ToCSVString(id string) string {
	return fmt.Sprintf("%s,%s,%s,%t,%s,%s,%d,%d,%x\n", id, fe.FileId, fe.Filename, fe.IsDir, fe.Dirname, fe.Path, fe.Date, fe.Size, fe.Hash)
}

func (fe *FileEntry) ToJSON(id string) []byte {
	fe.ScanId = id
	json, err := json.Marshal(fe)
	check(err)
	return json
}
