package db

func Insert[T any](record T) error {
	res := DB().Create(&record)
	return res.Error
}

func BatchInsert[T any](records []T) error {
	res := DB().Create(&records)
	return res.Error
}
