package ginip

import (
	"errors"
	"strconv"
	"strings"
)

func parseValueKey(line, val string) string {
	if strings.ContainsRune(line, '=') {
		kv := strings.Split(line, "=")
		kv[0] = strings.TrimSpace(kv[0])
		if kv[0] == val {
			kv[1] = strings.TrimSpace(kv[1])
			return kv[1]
		}
	}
	return ""
}

func findKey(lines []string, val string) (string, error) {
	for _, line := range lines {
		k := parseValueKey(line, val)
		if k == "" {
			continue
		}
		return k, nil
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
		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {
			if line[1:len(line)-1] == sec {
				ini = ini[i:]
				ss = true
				break
			}
		}
	}
	if ss {
		k, err := findKey(ini, val)
		return k, err
	}

	panic("unreachable")
}

func (i Ini) GetValueInt(sec, val string) (int, error) {
	s, err := getValue(i, sec, val)
	if err != nil {
		return 0, err
	}
	sint, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return sint, nil
}

func (i Ini) GetValueString(sec, val string) (string, error) {
	s, err := getValue(i, sec, val)
	return s, err
}

func (i Ini) GetValueFloat(sec, val string) (float64, error) {
	s, err := getValue(i, sec, val)
	if err != nil {
		return 0, err
	}
	sfloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return sfloat, nil

}
