package bonzo

import "time"

type Schedule struct {
	events []Event
}

type Event struct {
	Value      interface{}
	Expression Expression
}

type Expression interface {
	Express(time.Time) bool
}

type ExpressionFunc func(time.Time) bool

func (f ExpressionFunc) Express(t time.Time) bool {
	return f(t)
}

func (s *Schedule) Add(v interface{}, t string) {
	s.events = append(s.events, Event{
		Expression: ParseFake(t),
		Value:      v,
	})
}

func (s *Schedule) Occur(t time.Time) []interface{} {
	var vs []interface{}
	for _, e := range s.events {
		if e.Expression.Express(t) {
			vs = append(vs, e.Value)
		}
	}

	return vs
}

func (s *Schedule) Schedule(b, e time.Time, i time.Duration) map[time.Time]interface{} {
	es := make(map[time.Time]interface{})
	d := b

	for d.Before(e) {
		vs := s.Occur(d)
		if vs != nil {
			es[d] = vs[0]
		}
		d = d.Add(i)
	}

	return es
}

func NewScheduler(es ...Event) *Schedule {
	return &Schedule{
		events: es,
	}
}
