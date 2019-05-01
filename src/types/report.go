package types

import (
	"encoding/json"
	"time"
)

// Report api call return definition
type Report struct {
	Time         time.Time
	Publications []string
	Message      string
}

// ToString return report as json format
func (r *Report) ToString() string {
	report, _ := json.Marshal(r)
	return string(report)
}
