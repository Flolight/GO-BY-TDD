package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one field",
			struct {
				Name string
			}{"Chloe"},
			[]string{"Chloe"},
		},
		{
			"Struct with two fields",
			struct {
				Name string
				City string
			}{"Chloe", "Toulouse"},
			[]string{"Chloe", "Toulouse"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chloe", 28},
			[]string{"Chloe"},
		},
		{
			"Struct with nested structure",
			Person{
				"Chloe",
				Profile{28, "Toulouse"},
			},
			[]string{"Chloe", "Toulouse"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

}
