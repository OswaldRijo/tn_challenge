package utils

func Some[T comparable](mainArr []T, condition T) bool {
	for _, k := range mainArr {
		if k == condition {
			return true
		}
	}

	return false
}
