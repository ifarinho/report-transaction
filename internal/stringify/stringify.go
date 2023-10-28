package stringify

import "strconv"

func Int64(i int64) string {
	return strconv.FormatInt(i, 10)
}
