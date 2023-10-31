package decode

import "encoding/json"

func jsonUnmarshal[T any](b []byte) (v *T, err error) {
	return v, json.Unmarshal(b, &v)
}

func DeserializeJsonBytes[T any](b []byte) (v *T, err error) {
	return jsonUnmarshal[T](b)
}

func DeserializeJsonString[T any](s string) (v *T, err error) {
	return jsonUnmarshal[T]([]byte(s))
}
