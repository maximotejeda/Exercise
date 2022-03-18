package parser

import "testing"

type ParseTest struct {
	str  string
	want bool
}

type RparseTest struct {
	str   string
	token []rune
	ret   bool
	want  bool
}

var ParseTests = []ParseTest{
	ParseTest{"{}", true},
	ParseTest{"{([])}", true},
	ParseTest{"(){}", true},
	ParseTest{"{[]{}()}", true},
	ParseTest{"{[}]", false},
	ParseTest{"{", false},
	ParseTest{"}", false},
	ParseTest{"{[({)]}}", false},
}

var RparseTests = []RparseTest{
	RparseTest{"{}", nil, true, true},
	RparseTest{"{([])}", nil, true, true},
	RparseTest{"(){}", nil, true, true},
	RparseTest{"{[]{}()}", nil, true, true},
	RparseTest{"{[}]", nil, false, false},
	RparseTest{"{", nil, false, false},
	RparseTest{"}", nil, false, false},
	RparseTest{"{[({)]}}", nil, false, false},
}

func TestParse(t *testing.T) {
	for _, test := range ParseTests {
		if output := ParseSTR(test.str); output != test.want {
			t.Errorf("Output %v different to expected %v", test.want, test.want)
		}
	}
}

func TestRparse(t *testing.T) {
	for _, test := range RparseTests {
		if output := Rparser(test.str, test.token, test.ret); output != test.want {
			t.Errorf("Output %v different to expected %v", test.want, test.want)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseSTR("{}")
	}
}
func BenchmarkRparse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rparser("{}", nil, true)
	}
}
