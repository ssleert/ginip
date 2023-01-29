package ginip

import (
	"strconv"
)

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

func (i *Ini) GetValueBool(sec, val string) (bool, error) {
	s, err := getValue(*i, sec, val)
	if err != nil {
		return false, err
	}
	sbool, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return sbool, nil
}
