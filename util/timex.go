package util

import (
	"log"
	"strconv"
	"time"
)

func ParseUnix(s string) time.Time {
	if len(s) > 10 {
		s = s[:10]
	}
	i, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		log.Println(e)
		return time.Time{}
	}
	return time.Unix(i, 0)
}
