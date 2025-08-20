package streval

import (
	"strings"
	"testing"
)

var tests = []string{
	"Hello, ${name}",
	"Expected value ${expected} but got ${value}",
	"This is not an escape: \\${helloo}, but this is: ${foo.bar}",
}

func TestParse(t *testing.T) {
	for n, test := range tests {
		result := ""
		if err := Parse(test, Handlers{
			Literal: func(str string) error {
				result += strings.ReplaceAll(str, "${", "\\${")
				return nil
			},
			Expression: func(str string) error {
				result += "${" + str + "}"
				return nil
			},
		}); err != nil {
			t.Errorf("got error from parse #%d: %s", n, err.Error())
		}
		if result != test {
			t.Errorf("test %d yielded \"%s\", when \"%s\" was expected", n, result, test)
		}
	}
}
