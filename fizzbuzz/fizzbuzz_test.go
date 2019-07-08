package fizzbuzz

import "testing"

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := FizzBuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}
func TestInput3ShouldBeDisplay3(t *testing.T) {
	v := FizzBuzz(3)
	if "Fizz" != v {
		t.Error("fizzbuzz of 3 should be '3' but have", v)
	}
}
func TestFizzBuzz(t *testing.T) {
	v := FizzBuzz(15)
	if "FizzBuzz" != v {
		t.Error("error fizzbuzz :", v)
	}
}
