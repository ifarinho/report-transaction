package env

import (
	"log"
	"os"
)

var (
	RunMode                = getEnvOrFail("RUN_MODE")
	PostgresDataSourceName = getEnvOrFail("DB_POSTGRES_DATA_SOURCE_NAME")
	ServiceEmail           = getEnvOrFail("SERVICE_EMAIL")
	AwsAccessKeyId         = getEnvOrFail("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey     = getEnvOrFail("AWS_ACCESS_SECRET_KEY")
	AwsRegion              = getEnvOrFail("AWS_REGION")
	AwsS3Bucket            = getEnvOrFail("AWS_S3_BUCKET")
	AwsS3Prefix            = getEnvOrFail("AWS_S3_PREFIX")
	AwsCredentialToken     = getEnv("AWS_CREDENTIAL_TOKEN")
	CorsOrigin             = getEnv("CORS_ORIGIN")
)

func getEnvOrFail(key string) string {
	v := getEnv(key)
	if v == "" {
		log.Fatalf("fatal: empty value for env variable: %v", key)
	}
	return v
}

func getEnv(key string) string {
	return os.Getenv(key)
}
