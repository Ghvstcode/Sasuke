package model

import (
	"encoding/json"
	"time"
    "strconv"
)
// Cryptoresponse is
type Cryptoresponse struct {
	Markets []Markets `json:"Markets"`
}

// Markets is ...
type Markets struct {
	Label     string  `json:"Label"`
	Name      string  `json:"Name"`
	Price     float64 `json:"Price"`
	Volume24H float64 `json:"Volume_24h"`
	Timestamp int     `json:"Timestamp"`
}

// Result is ...
type Result struct {
	Label     string  `json:"Label"`
	Name      string  `json:"Name"`
	Price     float64 `json:"Price"`
	Volume    float64 `json:"Volume_24h"`
	Date      string   `json:"Timestamp"`
}

//FormattedDate returns string
func (m Markets) FormattedDate() string {
	i, err := strconv.ParseInt(string(m.Timestamp), 10, 64)

    if err != nil {
        panic(err)
	}
	tm := time.Unix(i, 0)
	return tm.String();
}

//JSON returns string
func (m Markets) JSON() string {
	mJSON, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(mJSON)
}