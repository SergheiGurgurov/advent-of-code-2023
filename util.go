package util

func CharToInt[T rune | byte](c T) int {
	return int(c - '0')
}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func First[T any](list []T, fn func(T) bool) T {
	for _, item := range list {
		if fn(item) {
			return item
		}
	}
	panic("No item found")
}

func Last[T any](list []T, fn func(T) bool) T {
	for i := len(list) - 1; i >= 0; i-- {
		if fn(list[i]) {
			return list[i]
		}
	}
	panic("No item found")
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Every[T any](ts []T, fn func(T) bool) bool {
	for i := range ts {
		if !fn(ts[i]) {
			return false
		}
	}
	return true
}
