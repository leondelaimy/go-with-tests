package reflection

import "testing"

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