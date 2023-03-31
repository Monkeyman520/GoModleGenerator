package set

type SET[KEY ~string | string | ~int | int, VALUE interface{}] map[KEY]VALUE

// DiffArray returns the difference of two arrays
func DiffArray[T ~string | string | ~int | int](a []T, b []T) []T {
	var diffArray []T
	var temp SET[T, struct{}] = map[T]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// IntersectArray returns the intersection of two arrays
func IntersectArray[T ~string | string | ~int | int](a []T, b []T) []T {
	var inter []T
	var mp SET[T, bool] = map[T]bool{}

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// RemoveRepByMap remove the repeat element
func RemoveRepByMap[T ~string | string | ~int | int](slc []T) []T {
	var result []T
	var tempMap SET[T, byte] = map[T]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
