package types

import "encoding/json"

func UnmarshalFacebookTimeline(data []byte) (FacebookTimeline, error) {
	var r FacebookTimeline
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FacebookTimeline) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FacebookTimeline struct {
	Data   []FacebookPost `json:"data"`
	Paging Paging         `json:"paging"`
}

type FacebookPost struct {
	CreatedTime string `json:"created_time"`
	Message     string `json:"message"`
	ID          string `json:"id"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
	Next    string  `json:"next"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
