package types

import (
	"encoding/json"
	"net/http"
	"time"
)

// Report api call return definition
type Report struct {
	Time         time.Time     `json:"time"`
	Publications []string      `json:"publications"`
	Errors       []http.Header `json:"errors"`
	Message      string        `json:"message"`
}

// ToString return report as json format
func (r *Report) ToString() string {
	report, _ := json.Marshal(r)
	return string(report)
}
