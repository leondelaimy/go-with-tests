package reflection

import (
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {

	t.Run("reflect float64", func(t *testing.T){
		var x float64 = 3.4
		got := Reflection(x)
		want := "type: float64"
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("reflect string", func(t *testing.T){
		var x string = "test"
		got := Reflection(x)
		want := "type: string"
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		var x float64 = 3.4
		Reflection(x)		
	}
}

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
			"struct with one string field",
			struct {
				Name string
			}{"Leon"},
			[]string{"Leon"},
		},
		{
			"struct with two string fields",
			struct {
					Name string
					City string
			}{"Leon", "Manchester"},
			[]string{"Leon", "Manchester"},
		},
		{
			"struct with non string field",
			struct {
					Name string
					Age  int
			}{"Leon", 30},
			[]string{"Leon"},
		},
		{
			"nested fields",
			struct {
					Name string
					Profile struct {
							Age  int
							City string
					}
			}{"Leon", struct {
					Age  int
					City string
			}{30, "Manchester"}},
			[]string{"Leon", "Manchester"},
		},
		{
			"nested fields",
			Person{
					"Leon",
					Profile{30, "Manchester"},
			},
			[]string{"Leon", "Manchester"},
		},
		{
			"pointers to things",
			&Person{
					"Leon",
					Profile{30, "Manchester"},
			},
			[]string{"Leon", "Manchester"},
		},
		{
			"slices",
			[]Profile {
					{33, "Manchester"},
					{34, "London"},
			},
			[]string{"Manchester", "London"},
		},
		{
			"arrays",
			[2]Profile {
					{33, "Manchester"},
					{34, "London"},
			},
			[]string{"Manchester", "London"},
		},
	}
	
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
	
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
	
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
	
		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Amsterdam"}
			close(aChannel)
		}()
	
		var got []string
		want := []string{"Berlin", "Amsterdam"}
	
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Amsterdam"}
		}
	
		var got []string
		want := []string{"Berlin", "Amsterdam"}
	
		walk(aFunction, func(input string) {
			got = append(got, input)
		})
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}