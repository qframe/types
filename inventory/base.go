package inventory

import (
	"time"
	"encoding/json"
)
const nanoX = 1000000000

func SplitUnixNano(t int64) (sec, nano int64) {
	sec = t/nanoX
	nano = t - sec*nanoX
	return
}

type Base struct {
	Time 			time.Time
	TimeUnixNano	int64				`json:"time"`
	Subject			string 				`json:"subject"`	// Subject of what is going on (e.g. container)
	Action			string				`json:"action"`
	Object  		string        	 	`json:"object"` 	// Passive object
	Tags 			map[string]string 	`json:"tags"` 		// Tags that should be applied to the action
}


func NewBaseFromJson(str string) (b Base, err error) {
	err = json.Unmarshal([]byte(str), &b)
	s,n := SplitUnixNano(b.TimeUnixNano)
	b.Time = time.Unix(s, n)
	return
}
