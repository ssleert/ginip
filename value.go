package ginip

import (
	"errors"
	"strconv"
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
		if isSection(line) {
			if line[1:len(line)-1] == sec {
				ini = ini[i+1:]
				ss = true
				break
			}
		}
	}
	for i, line := range ini {
		if isSection(line) {
			ini = ini[:i]
		}
	}

	if ss {
		k, err := findKey(ini, val)
		return k, err
	}

	panic("unreachable")
}

func (i *Ini) GetValueInt(sec, val string) (int, error) {
	s, err := getValue(*i, sec, val)
	if err != nil {
		return 0, err
	}
	sint, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return sint, nil
}

func (i *Ini) GetValueString(sec, val string) (string, error) {
	s, err := getValue(*i, sec, val)
	return s, err
}

func (i *Ini) GetValueFloat(sec, val string) (float64, error) {
	s, err := getValue(*i, sec, val)
	if err != nil {
		return 0, err
	}
	sfloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return sfloat, nil
}
