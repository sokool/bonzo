package bonzo

import (
	"fmt"
	"strings"
)

const everyWeek = "every week at"

func ParseFake(s string) Expression {
	e, err := parseEveryWeek(s)
	if err != nil {
		fmt.Println(err)
	}
	return e
}

func parseEveryWeek(s string) (Expression, error) {
	i := strings.Index(s, everyWeek)
	if i == -1 {
		return nil, fmt.Errorf("niomo")

	}

	ws := strings.TrimSpace(s[len(everyWeek)+i:])

	w, ok := weekdays[ws]
	if !ok {
		return nil, fmt.Errorf("chujnia")

	}

	return Weekday(w), nil
}
