package ticket

import "time"

// 主键生成器
func Generate() string {
	return time.Now().String()
}
