package decode

import "encoding/json"

func DeserializeJson[T any](b []byte) (v *T, err error) {
	return v, json.Unmarshal(b, &v)
}
