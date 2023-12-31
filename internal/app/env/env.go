package env

import (
	"log"
	"os"
)

var (
	RunMode                 = getEnvOrFatal("ENV_RUN_MODE")
	PostgresDataSourceName  = getEnvOrFatal("ENV_POSTGRES_DATA_SOURCE_NAME")
	PostgresBatchCreateSize = getEnvOrFatal("ENV_POSTGRES_BATCH_CREATE_SIZE")
	ServiceEmail            = getEnvOrFatal("ENV_SERVICE_EMAIL")
	CorsOrigin              = getEnvOrFatal("ENV_CORS_ORIGIN")
	AllowedMethods          = getEnvOrFatal("ENV_ALLOWED_METHODS")
	AwsAccessKeyId          = getEnvOrFatal("ENV_AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey      = getEnvOrFatal("ENV_AWS_ACCESS_SECRET_KEY")
	AwsRegion               = getEnvOrFatal("ENV_AWS_REGION")
	AwsS3Bucket             = getEnvOrFatal("ENV_AWS_S3_BUCKET")
	AwsS3Prefix             = getEnvOrFatal("ENV_AWS_S3_PREFIX")
	AwsCredentialToken      = getEnv("ENV_AWS_CREDENTIAL_TOKEN")
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
