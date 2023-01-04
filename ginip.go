package ginip

import (
	"errors"
	"strings"
)

func notAComment(l string) bool {
	if len(l) > 1 &&
		l[0] == '#' ||
		l[0] == ';' {
		return false
	}
	return true
}

func isSection(l string) bool {
	if len(l) > 2 &&
		l[0] == '[' &&
		l[len(l)-1] == ']' {
		return true
	}
	return false
}

func findKey(lines []string, val string) (string, error) {
	var key string
	for _, line := range lines {
		if strings.ContainsRune(line, '=') && notAComment(line) {
			kv := strings.Split(line, "=")
			if strings.TrimSpace(kv[0]) == val {
				key = strings.TrimSpace(kv[1])
			}
		}
		if key == "" {
			continue
		}
		return key, nil
	}
	return "", errors.New("value doesn't finded")
}

func getValue(ini Ini, sec, val string) (string, error) {
	if len(ini) == 0 {
		return "", errors.New("ini file is empty")
	} else if val == "" {
		return "", errors.New("value is \"\"")
	} else if sec == "" {
		k, err := findKey(ini, val)
		return k, err
	}

	var ss bool
	for i, line := range ini {
		if ss && isSection(line) {
			ini = ini[:i]
		}
		if !ss && isSection(line) {
			if line[1:len(line)-1] == sec {
				ini = ini[i+1:]
				ss = true
			}
		}
	}

	if ss {
		k, err := findKey(ini, val)
		return k, err
	}

	panic("unreachable")
}
