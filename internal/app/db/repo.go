package db

func BatchInsert(records any) error {
	res := DB().Create(records)
	return res.Error
}

func SelectById[T any](id uint) (*T, error) {
	result := new(T)
	res := DB().Where("id = ?", id).First(&result)
	return result, res.Error
}
