package bonzo

import "time"

func Every(b time.Time, d time.Duration) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		return (t.Equal(b) || t.After(b)) && t.Sub(b)%d == 0
	})
}

func Weekday(ws ...time.Weekday) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		for _, w := range ws {
			if t.Weekday() == w {
				return true
			}
		}
		return false
	})
}

func Range(b, e time.Time) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		if e.IsZero() && (t.After(b) || t.Equal(b)) {
			return true
		}

		return (t.After(b) || t.Equal(b)) && t.Before(e)
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

func Differene(es ...Expression) Expression {
	return ExpressionFunc(func(t time.Time) bool {
		return false
	})
}
