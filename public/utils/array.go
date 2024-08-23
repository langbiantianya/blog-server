package utils

func Map[T1, T2 any](list []T1, transform func(index int, item T1) (T2, error)) []T2 {
	arr := make([]T2, 0)
	for i, i2 := range list {
		res, err := transform(i, i2)
		if err == nil {
			arr = append(arr, res)
		}
	}
	return arr
}

func Filter[T any](list []T, filter func(item T) bool) []T {
	arr := make([]T, 0)
	for _, t := range list {
		if filter(t) {
			arr = append(arr, t)
		}
	}
	return arr
}

func Flatten[T any](arr [][]T) []T {
	var res []T
	for _, subArr := range arr {
		res = append(res, subArr...)
	}
	return res
}

func DistinctBy[T any](list []T, filter func(item T) any) []T {
	uniqueMap := make(map[any]bool)
	var uniqueArr []T

	for _, item := range list {
		if !uniqueMap[filter(item)] {
			uniqueMap[filter(item)] = true
			uniqueArr = append(uniqueArr, item)
		}
	}
	return uniqueArr
}
