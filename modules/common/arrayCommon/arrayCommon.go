package arrayCommon

func Map[T interface{}, V interface{}](arr []T, callback func(val T) V) []V {
	var returnArr []V

	for _, val := range arr {
		returnArr = append(returnArr, callback(val))
	}

	return returnArr
}

func Filter[T interface{}](arr []T, callback func(val T) bool) []T {
	var returnArr []T

	for _, val := range arr {
		if callback(val) {
			returnArr = append(returnArr, val)
		}
	}
	return returnArr
}

func Reduce[T interface{}](arr []T, init int, callback func(val T) int) int {
	returnNum := init

	for _, val := range arr {
		returnNum += callback(val)
	}

	return returnNum
}

func ForEach[T interface{}](arr []T, callback func(val T) bool) {
	for _, val := range arr {
		if !callback(val) {
			break
		}
	}
}

func KeyBy[T interface{}, V comparable](arr []T, callback func(val T) V) map[V]T {
	keyByArr := map[V]T{}
	for _, val := range arr {
		keyByArr[callback(val)] = val
	}
	return keyByArr
}

func GroupBy[T interface{}, V comparable](arr []T, callback func(val T) V) map[V][]T {
	groupByArr := map[V][]T{}
	for _, val := range arr {
		if getArr, ok := groupByArr[callback((val))]; ok {
			getArr = append(getArr, val)
			groupByArr[callback((val))] = getArr
		} else {
			groupByArr[callback((val))] = []T{val}
		}
	}
	return groupByArr
}
