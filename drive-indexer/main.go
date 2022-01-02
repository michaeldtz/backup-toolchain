package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ctx := context.Background()

	client, err := GetClient(".secrets/clientsecret.json")
	check(err)

	drive, err := drive.NewService(ctx, option.WithHTTPClient(client))
	check(err)

	pageToken := ""
	cnt := 0

	file, err := os.Create("../index-store/drive-index.json")
	check(err)
	defer file.Close()

	for {

		r, err := drive.Files.List().
			PageToken(pageToken).
			PageSize(1000).
			Q("trashed = false").
			Fields("nextPageToken, files(id, name, kind, mimeType, parents, size, modifiedTime, md5Checksum)").Do()
		check(err)

		if len(r.Files) == 0 {
			fmt.Println("No files found.")
			break
		} else {
			for _, entry := range r.Files {

				cnt++
				if cnt%10000 == 0 {
					fmt.Printf("%d items in drive processed\n", cnt)
				}

				dirname := "/"
				if len(entry.Parents) > 0 {
					dirname += entry.Parents[len(entry.Parents)-1]
				}

				layout := "2006-01-02T15:04:05.000Z"
				timestamp, err := time.Parse(layout, entry.ModifiedTime)
				check(err)

				fe := &FileEntry{
					FileId:   entry.Id,
					Filename: entry.Name,
					IsDir:    entry.MimeType == "application/vnd.google-apps.folder",
					Date:     timestamp.Unix(),
					Size:     entry.Size,
					Path:     "/" + strings.Join(entry.Parents, "/"),
					Dirname:  dirname,
					Hash:     []byte(entry.Md5Checksum),
				}

				jsonLineToWrite := fe.ToJSON("drive-index")
				file.Write(jsonLineToWrite)
				file.Write([]byte("\n"))

			}

			pageToken = r.NextPageToken

			ptFile, err := os.Create(".secrets/pagetoken.json")
			check(err)
			ptFile.Write([]byte(pageToken))
			ptFile.Close()

			if pageToken == "" {
				break
			}

		}
	}

}
