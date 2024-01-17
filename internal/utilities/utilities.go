package utilities

// validateFields is a generic fields validator of keys in a map
func ValidateFields(m map[string]any, fields ...string) bool {
	for _, field := range fields {
		_, ok := m[field]
		if !ok {
			return false
		}
	}
	return true
}
