package utils

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

var capyAlts map[string]string

func LoadCapyAlts(fn string) error {
	content, err := os.ReadFile(fn)
	if err != nil {
		return errors.Wrap(err, "reading alt.json")
	}
	if err := json.Unmarshal(content, &capyAlts); err != nil {
		return errors.Wrap(err, "unmarshaling alt.json")
	}
	return nil
}

func GetAlt(index string) string {
	alt, ok := capyAlts[index]
	if !ok {
		return "a capybara"
	}
	return alt
}

func GetAlti(index int) string {
	return GetAlt(strconv.Itoa(index))
}
