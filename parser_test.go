package bonzo_test

import (
	"testing"

	"github.com/sokool/bonzo"
)

type parseTest struct {
	text string
	//expects error
}

var parseTests = []parseTest{
	//{"every Friday"},
	{"  every Monday in March but not Monday on 20 March"},
	{"every Monday in March 2017 and in every April"},
	{"every Friday"},
	{"every Tuesday from 1 April 2017"},
	{"every Tuesday, Wednesday from 1 Jan to 3 Feb"},
	{"every 5 Days"},
	{"every 5 Days after first week of every month"},
	{"every Day and n	ot Friday"},
	{"on Today to 9 Jan"},
	{"on 17 March"},
	{"in every April without Saturday and Sunday"},
	{"on 27 April or 27 November but not 13 December"},
	{"in April and May but not between 30 April and 5 May"},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		bonzo.Parse(tt.text)
	}
}
