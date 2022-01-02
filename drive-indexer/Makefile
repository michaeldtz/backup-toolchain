
index_gdrive:
	go run *.go 
	gsutil cp ../index-store/drive-index.json gs://midietz-backup-bucket-standard/admin/
	bq --location=europe-west3 load --replace=true --source_format=NEWLINE_DELIMITED_JSON --autodetect backup_file_index.drive_index gs://midietz-backup-bucket-standard/admin/drive-index.json
	