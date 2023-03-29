package internetarchive

import (
	"encoding/json"
	"strings"
)

type stringOrArray string

func (sa *stringOrArray) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		switch data[0] {
		case '"':
			var s string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*sa = stringOrArray(s)
		case '[':
			var s []string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*sa = stringOrArray(strings.Join(s, ", "))
		}
	}
	return nil
}
