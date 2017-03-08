package bonzo

import "time"

func WeekDay(ws ...time.Weekday) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		for _, w := range ws {
			if t.Weekday() == w {
				return true
			}
		}
		return false
	})
}

func TimeRange(b, e time.Time) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		return t.After(b) && t.Before(e)
	})
}

func Intersection(es ...Expression) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		for _, e := range es {
			if !e.Express(t) {
				return false
			}
		}
		return true
	})
}

func Union(es ...Expression) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		for _, e := range es {
			if e.Express(t) {
				return true
			}
		}
		return false
	})
}
