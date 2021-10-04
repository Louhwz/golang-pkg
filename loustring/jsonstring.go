package loustring

func CreateJsonString(key, value interface{}) string {
	m := make(map[interface{}]interface{})
	m[key] = value
	return ToJsonString(m)
}
