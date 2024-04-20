package helper

// GroupBy 使用泛型实现列表分组
func GroupBy[K comparable, T any](list []T, keyFunc func(T) K) map[K][]T {
	group := make(map[K][]T)
	for _, item := range list {
		key := keyFunc(item)
		group[key] = append(group[key], item)
	}
	return group
}
