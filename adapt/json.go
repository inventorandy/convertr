package adapt

import "encoding/json"

// UsingJSON converts from and to using the JSON struct tags.
func UsingJSON(from, to interface{}) error {
	b, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, to)
}
