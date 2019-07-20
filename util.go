package BplusTree

func generateKeyBinarySearchFunc(compareFunc func(a, b interface{}) int, keyExample interface{}) func(data []interface{}, key interface{}, size int) int {
	if compareFunc == nil {
		switch keyExample.(type) {
		case int:
			compareFunc = func(a, b interface{}) int {
				if a.(int) < b.(int) {
					return -1
				} else if a.(int) > b.(int) {
					return 1
				}
				return 0
			}
		case float64:
			compareFunc = func(a, b interface{}) int {
				if a.(float64) < b.(float64) {
					return -1
				} else if a.(float64) > b.(float64) {
					return 1
				}
				return 0
			}
		case string:
			compareFunc = func(a, b interface{}) int {
				if a.(string) < b.(string) {
					return -1
				} else if a.(string) > b.(string) {
					return 1
				}
				return 0
			}
		default:
			panic("请定义key比较规则")
		}
	}
	return func(keys []interface{}, key interface{}, size int) int {
		low := 0
		high := size - 1
		mid := low + (high-low)/2
		var res int
		for low <= high {
			res = compareFunc(keys[mid], key)
			if res == 0 {
				//找到了
				return mid
			} else if res > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
			mid = low + (high-low)/2
		}
		//没找到 但是返回 最接近key且大于key的 位置下标 的相反数
		return -low - 1
	}
}

