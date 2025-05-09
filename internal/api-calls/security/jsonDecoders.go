package security

import (
	"encoding/json"
	"fmt"
)

func (s *SeverityField) UnmarshalJSON(data []byte) error {
	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		*s = arr
		return nil
	}

	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*s = []string{single}
		return nil
	}

	return fmt.Errorf("unexpected severity format: %s", string(data))
}

func (m *MsgField) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.Text); err == nil {
		return nil
	}

	if err := json.Unmarshal(data, &m.Other); err == nil {
		return nil
	}

	return fmt.Errorf("unexpected msg format: %s", string(data))
}
