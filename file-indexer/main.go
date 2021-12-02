package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func executeOnDir(chn chan<- FileEntry, path string, doClose bool) {
	dir, err := os.ReadDir(path)
	check(err)

	for _, entry := range dir {
		if entry.IsDir() {
			nextPath := path + "/" + entry.Name()
			executeOnDir(chn, nextPath, false)
		} else {
			fileInfo, err := entry.Info()
			check(err)

			//Calc hash
			fileReader, err := os.Open(path + "/" + entry.Name())
			check(err)
			defer fileReader.Close()

			hasher := NewHasher()
			io.Copy(hasher, fileReader)
			hash := hasher.Sum()

			pathParts := strings.Split(path, "/")
			dirName := pathParts[len(pathParts)-1]

			fileEntry := FileEntry{
				Filename: entry.Name(),
				Date:     fileInfo.ModTime().Unix(),
				Size:     fileInfo.Size(),
				Path:     path,
				Dirname:  dirName,
				Hash:     hash,
			}
			chn <- fileEntry
		}

	}

	if doClose {
		close(chn)
	}
}

func initFileWriter(chn <-chan FileEntry, outFile string, scanId string) {
	file, err := os.Create(outFile)
	check(err)
	defer file.Close()

	if strings.HasSuffix(outFile, ".csv") {
		file.Write([]byte("Id,Filename,Directory,FullPath,Date,Size,Hash\n"))
	}

	cnt := 0

	for entry := range chn {

		cnt++
		if cnt%1000 == 0 {
			fmt.Printf("%d files processed\n", cnt)
		}

		if strings.HasSuffix(outFile, ".csv") {
			csvLineToWrite := fmt.Sprint(scanId, ",", entry.toCSVString())
			file.Write([]byte(csvLineToWrite))
		} else if strings.HasSuffix(outFile, ".json") {
			jsonLineToWrite := entry.toJSON(scanId)
			file.Write(jsonLineToWrite)
			file.Write([]byte("\n"))
		}
	}
}

func main() {

	outFilePtr := flag.String("out", "", "output csv file with the index")
	scanIdPtr := flag.String("id", "", "a id for the dir that is scanned")
	flag.Parse()

	tail := flag.Args()
	dirName := tail[0]

	writeChan := make(chan FileEntry)

	go executeOnDir(writeChan, dirName, true)

	fullOutFile := *outFilePtr
	if fullOutFile == "" {
		fullOutFile = dirName + "/fileindex.csv"
	}

	initFileWriter(writeChan, fullOutFile, *scanIdPtr)

	fmt.Println("done")
}
