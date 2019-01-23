package ticket

import (
	"time"
	"strconv"
)

// 主键生成器，按照时间顺序，排序用，后期加入idc信息
// string, 后期的长度会更长
// scheduler
func Generate() string {
	r := time.Now().UnixNano()
	return strconv.FormatInt(r, 10)
}
