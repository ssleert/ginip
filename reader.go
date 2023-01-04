package ginip

import (
	"strings"

	"github.com/ssleert/sfolib"
)

func Load(f string) (Ini, error) {
	ini, err := sfolib.LoadFile(f)
	if err != nil {
		return Ini{}, err
	}

	for i, e := range ini {
		ini[i] = strings.TrimSpace(e)
	}
	return ini, nil
}
