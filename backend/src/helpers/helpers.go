package helpers

func In[T comparable](arr []T, val T) bool {
	for _, arrVal := range arr {
		if arrVal == val {
			return true
		}
	}

	return false
}
