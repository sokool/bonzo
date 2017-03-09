package bonzo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sokool/bonzo"
)

type rangeTest struct {
	now      string
	from     string
	to       string
	expected bool
}

var n string = "2017-03-21"
var rangeTests = []rangeTest{
	{n, "2017-03-01", "2017-03-31", true},
	{n, "2017-03-20", "2017-03-22", true},
	{n, "", "2017-03-27", true},
	{n, "", n, false},
	{n, "2017-03-20", "", true},
	{n, n, "", true},
	{n, n, "2017-03-22", true},
	{n, "", "", true},
}

func TestRange(t *testing.T) {
	for _, tt := range rangeTests {
		var err error

		now, err := newTime(tt.now)
		if err != nil {
			t.Fatal(err)
		}

		from, err := newTime(tt.from)
		if err != nil {
			t.Fatal(err)
		}

		to, err := newTime(tt.to)
		if err != nil {
			t.Fatal(err)
		}

		actual := bonzo.Range(from, to).Express(now)
		if actual != tt.expected {
			t.Errorf("bonzo.Range(%s, %s) for %s: expected %v, actual %v",
				from.Format("2006-01-02 15:04:05"),
				to.Format("2006-01-02"),
				now.Format("2006-01-02"),
				tt.expected,
				actual)
		}
	}
}

func newTime(s string) (time.Time, error) {
	var l string
	if len(s) == 0 {
		return time.Time{}, nil
	}

	if len(s) == 10 {
		l = "2006-01-02"
	}

	if len(s) == 19 {
		l = "2006-01-02 15:04:05"
	}

	if l == "" {
		return time.Time{}, fmt.Errorf("wrong date format")
	}

	t, err := time.Parse(l, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("wrong date format of %s, try %s", s, l)
	}

	return t, nil
}
