package model

import (
	"fmt"
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
func (r Result) JSON() string {
	rJSON, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(rJSON)
}

func (m Markets) Result() Result {
	return Result{   
		Name: m.Name,      
		Price: m.Price,      
		Volume: m.Volume24H,    
		Date: m.FormattedDate(),      
	}
}

func (r Result) TextOutput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice : %d\nVolume: %s\nDate: %s\n",
		 r.Name, r.Price, r.Volume, r.Date)
	return p
}