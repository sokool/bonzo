package bonzo

import (
	"time"

	"strings"

	"sort"

	"fmt"
)

// ----------------
// every {Year, Month, Weekday, Day} ==> Every, Weekday
// in {Month, Year} ==> Range
// from, to, on {Date}, ==> Range

//
// ,|and|with {Intersection)
// or| {Union}
// not|without {Difference)

// X every (Monday in March but) not ( Monday on 20 March)
// X every Monday in March 2017 and in every April
// X every Friday
// X every Tuesday from 1 April 2017
// X (every Tuesday, Wednesday from 1 Jan to 3 Feb
// X every 5 Days
// X every 5 Days after first week of every month
// X every Day and
// X on Today to 9 Jan
// X on 17 March
// X in every April without Saturday and Sunday [april not sat and sun]
// X (on 27 April) or (27 November but) not (13 December)

var weekdays = map[string]time.Weekday{
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Sunday":    time.Sunday,
	"Saturday":  time.Saturday,
}

var month = map[string]time.Month{
	"January":   time.January,
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"May":       time.May,
	"June":      time.June,
	"July":      time.July,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

const (
	DateDMYNormal = "2 January 2006"
	DateDMYShort  = "2 Jan 2006"
	DateDMNormal  = "2 January"
	DateDMShort   = "2 Jan"
)

var operators = map[string]string{
	"every":   "every (Month|Weekday|XDay|XDays)",
	"from":    "from Date (to Date)",
	"to":      "Range",
	"on":      "Range",
	"in":      "Range",
	"and":     "Intersection",
	"not":     "Difference",
	"without": "",
	"or":      "Union",
	"between": "",
}

type token struct {
	kind  string
	start int
	end   int
	value string
}

type byStart []token

func (a byStart) Len() int           { return len(a) }
func (a byStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStart) Less(i, j int) bool { return a[i].start < a[j].start }

type tree struct {
	l *tree
	v interface{}
	r *tree
}

func (t *tree) insertLeft() {

}

func (t *tree) insertRight() {

}
func (t *tree) Left() {

}
func (t *tree) Right() {

}

func Parse(s string) {

	parse(fmt.Sprintf(" %s", s))
	//pretty.Println(s, o)

}

func parse(s string) []token {
	o := []token{}
	for op := range operators {
		n := 0
		x := 0
		op = " " + op + " "
		for {
			n = strings.Index(s[x:], op)

			if n == -1 {
				break
			}

			x = n + len(op) + x
			o = append(o, token{
				kind:  op,
				start: x,
			})
		}

	}
	sort.Sort(byStart(o))

	for i, t := range o {
		o[i].kind = strings.TrimSpace(o[i].kind)
		if i == 0 {
			continue
		}

		p := &o[i-1]
		p.end = t.start - len(t.kind)

		if p.start > p.end {
			continue
		}
		p.value = s[p.start:p.end]

	}

	l := len(o)
	if l == 0 {
		return o
	}

	o[l-1].end = len(s)
	o[l-1].value = s[o[l-1].start:o[l-1].end]

	return o
}
