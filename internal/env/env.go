package env

import "os"

var (
	FileTargetPath         = os.Getenv("TARGET_FULL_PATH")
	PostgresDataSourceName = os.Getenv("DB_POSTGRES_DATA_SOURCE_NAME")
	AwsRegion              = os.Getenv("AWS_REGION")
	AwsAccessKeyId         = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey     = os.Getenv("AWS_ACCESS_SECRET_KEY")
	AwsCredentialToken     = os.Getenv("AWS_CREDENTIAL_TOKEN")
	AwsS3Bucket            = os.Getenv("AWS_S3_BUCKET")
	AwsFullPath            = os.Getenv("AWS_FULL_PATH")
	ServiceEmail           = os.Getenv("SERVICE_EMAIL")
)
