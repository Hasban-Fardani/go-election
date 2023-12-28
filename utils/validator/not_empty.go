package validator

func EmptyStr(v ...string) bool {
	for _, v := range v {
		if v == "" {
			return true
		}
	}
	return false
}
