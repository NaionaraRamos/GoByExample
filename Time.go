package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	f := fmt.Printf

	now  := time.Now()
	p(now)

/* 	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))


	secs := now.Unix()
	nanos := now.UnixNano()
	p(now)

	millis := nanos /1000000
	p(secs)
	p(millis)
	p(nanos)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos)) */

	p(now.Format(time.RFC3339))

	t1, _:= time.Parse(
		time.RFC3339, "2012-11-01T22:08:41+00:00")
	
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"

	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	f("Representações numéricas: %d-%02d-%02dT%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

/* 	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e) */
	
}