/*
 * Copyright (c) 2021 武汉美宸时科科技有限公司 fashiontech.top
 */

package strings

// 指定的数组中是否包含某个元素
func Contains(s []string, substr ...string) bool {
	for _, si := range s {
		for _, subsi := range substr {
			if si == subsi {
				return true
			}
		}
	}
	return false
}
