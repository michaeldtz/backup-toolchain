
index_gdrive:
	go run *.go -id gdrive "/Volumes/GoogleDrive/Meine Ablage"

on_download:
	go run *.go -out ../index-store/test-fileindex.json -id test ~/Downloads

index_mdbackup:
	go run *.go -out ../index-store/mdbackup-fileindex.json -id mdbackup_ssd "/Volumes/MD Backup/Backups/BackupPrivate"
	gsutil cp ../index-store/mdbackup-fileindex.json gs://midietz-backup-bucket-standard/admin/
	bq --location=europe-west3 load --replace=true --source_format=NEWLINE_DELIMITED_JSON --autodetect backup_file_index.mdbackup_index gs://midietz-backup-bucket-standard/admin/mdbackup-fileindex.json

some_md5_test:
	go run *.go -out ../index-store/test-fileindex.json -id test "/Volumes/Google Drive/Meine Ablage/100 Michael/Bewerbung/SoftwareAG"
