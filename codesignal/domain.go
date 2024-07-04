package codesignal

import "strings"

func domainMapper(domains []string) []string {
	var result []string
	m := map[string]string{
		"com":  "commercial",
		"org":  "organization",
		"net":  "network",
		"info": "information",
	}
	for _, v := range domains {
		data := strings.Split(v, ".")
		k := data[len(data)-1]
		if val, ok := m[k]; ok {
			result = append(result, val)
		}
	}
	return result
}
