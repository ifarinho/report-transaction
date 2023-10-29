package env

import (
	"log"
	"os"
)

var (
	PostgresDataSourceName = getOrFail("DB_POSTGRES_DATA_SOURCE_NAME")
	AwsRegion              = getOrFail("AWS_REGION")
	AwsAccessKeyId         = getOrFail("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey     = getOrFail("AWS_ACCESS_SECRET_KEY")
	AwsCredentialToken     = getOrFail("AWS_CREDENTIAL_TOKEN")
	AwsS3Bucket            = getOrFail("AWS_S3_BUCKET")
	AwsS3Prefix            = getOrFail("AWS_S3_PREFIX")
	ServiceEmail           = getOrFail("SERVICE_EMAIL")
	CorsOrigin             = getOrFail("CORS_ORIGIN")
)

func getOrFail(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("fatal: empty value for env variable: %v", v)
	}
	return v
}
