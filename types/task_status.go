package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
	Archive
	Pending
)

var _ fmt.Stringer = (*TaskStatus)(nil)
var _ json.Marshaler = (*TaskStatus)(nil)
var _ json.Unmarshaler = (*TaskStatus)(nil)
var _ driver.Valuer = (*TaskStatus)(nil)
var _ sql.Scanner = (*TaskStatus)(nil)

func (s TaskStatus) String() string {
	return [...]string{"Todo", "InProgress", "Done", "Archive", "Pending"}[s]
}

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TaskStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	for i, v := range [...]string{"Todo", "InProgress", "Done", "Archive", "Pending"} {
		if v == str {
			*s = TaskStatus(i)
			return nil
		}
	}
	return fmt.Errorf("invalid TaskStatus %q", str)
}

func (s TaskStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

func (s *TaskStatus) Scan(src any) error {
	i, ok := src.(int64)
	if !ok {
		return fmt.Errorf("cannot scan %T into TaskStatus", src)
	}
	*s = TaskStatus(i)
	return nil
}
