package tempconv

import "testing"

func TestCToF(t *testing.T) {
	var tests = []struct{
		input Celsius
		expected Fahrenheit
	}{
		{Celsius(0), Fahrenheit(32)},
		{Celsius(100), Fahrenheit(212)},
	}

	for _, test := range tests {
		if got := CToF(test.input); got != test.expected {
			t.Errorf("CToF(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct{
		input Celsius
		expected Kelvin
	}{
		{Celsius(0), Kelvin(273.15)},
		{Celsius(100), Kelvin(373.15)},
	}

	for _, test := range tests {
		if got := CToK(test.input); got != test.expected {
			t.Errorf("CToK(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct{
		input Fahrenheit
		expected Celsius
	}{
		{Fahrenheit(32), Celsius(0)},
		{Fahrenheit(212), Celsius(100)},
	}

	for _, test := range tests {
		if got := FToC(test.input); got != test.expected {
			t.Errorf("FToC(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct{
		input Fahrenheit
		expected Kelvin
	}{
		{Fahrenheit(32), Kelvin(273.15)},
		{Fahrenheit(212), Kelvin(373.15)},
	}

	for _, test := range tests {
		if got := FToK(test.input); got != test.expected {
			t.Errorf("FToK(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct{
		input Kelvin
		expected Celsius
	}{
		{Kelvin(273.15), Celsius(0)},
		{Kelvin(373.15), Celsius(100)},
	}

	for _, test := range tests {
		if got := KToC(test.input); got != test.expected {
			t.Errorf("KToC(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct{
		input Kelvin
		expected Fahrenheit
	}{
		{Kelvin(273.15), Fahrenheit(32)},
		{Kelvin(373.15), Fahrenheit(212)},
	}

	for _, test := range tests {
		if got := KToF(test.input); got != test.expected {
			t.Errorf("KToF(%g) = %g, expected %g", test.input, got, test.expected)
		}
	}
}
