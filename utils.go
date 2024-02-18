package cowww

func getHeaderByKey(m map[string]string, key string) string {
	val, ok := m[key]
	if !ok {
		return ""
	}

	return val
}
