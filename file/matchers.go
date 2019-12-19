package file

import "encoding/json"

// Matches represents the matches in a given file
type Matches struct {
	json.Marshaler
	Companies Words
}

// MarshalJSON implements the json.Marshaler interface
func (m *Matches) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"companies": m.Companies,
	})
}
