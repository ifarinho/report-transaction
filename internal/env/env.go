package env

import "os"

var (
	PostgresDataSourceName = os.Getenv("DB_POSTGRES_DATA_SOURCE_NAME")
	FileTargetPath         = os.Getenv("TARGET_FULL_PATH")
)
