package arr

func Map[T any, T2 any](list []T, f func(item T, index int) T2) []T2 {

	var arr []T2
	for k, v := range list {
		newItem := f(v, k)
		arr = append(arr, newItem)
	}

	return arr
}

func ForEach[T any](list []T, f func(item T, index int)) {

	for k, v := range list {
		f(v, k)
	}
}
