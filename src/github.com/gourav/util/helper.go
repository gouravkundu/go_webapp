package util

import "bytes"

type Name string

func (s Name) ToCapitalize() string {
	// strings.Title() with do the same as the following func implementations
	if len(s) > 0 {
		chars := []byte(string(s))
		start := bytes.ToUpper([]byte{chars[0]})
		rest := bytes.ToLower(chars[1:])
		result := bytes.Join([][]byte{start, rest}, nil)
		return string(result)
	} else {
		return string(s)
	}
}
