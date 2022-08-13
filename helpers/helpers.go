package helpers

func Filter(array []string, test func(string) bool) (ret []string) {
	for _, s := range array {
		if test(s) {
			ret = append(ret, s)
		}
	}

	return ret
}
