package env

import (
	"log"
	"os"
)

var (
	RunMode                = getEnvOrFatal("RUN_MODE")
	PostgresDataSourceName = getEnvOrFatal("DB_POSTGRES_DATA_SOURCE_NAME")
	ServiceEmail           = getEnvOrFatal("SERVICE_EMAIL")
	AwsAccessKeyId         = getEnvOrFatal("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey     = getEnvOrFatal("AWS_ACCESS_SECRET_KEY")
	AwsRegion              = getEnvOrFatal("AWS_REGION")
	AwsS3Bucket            = getEnvOrFatal("AWS_S3_BUCKET")
	AwsS3Prefix            = getEnvOrFatal("AWS_S3_PREFIX")
	AwsCredentialToken     = getEnv("AWS_CREDENTIAL_TOKEN")
	CorsOrigin             = getEnv("CORS_ORIGIN")
)

func getEnvOrFatal(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("fatal: empty value for env variable: %v", key)
	}
	return val
}

func getEnv(key string) string {
	return os.Getenv(key)
}
