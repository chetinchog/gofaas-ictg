package main

import "testing"

// TestHelloEmpty function
func TestHelloEmpty(t *testing.T) {
	result := Hello("")
	if result != "Hello Dude!" {
		t.Errorf("Failed > Hello(\"\"): %v ? %v", "Hello Dude!", result)
	} else {
		t.Logf("Success > Hello(\"\"): %v ? %v", "Hello Dude!", result)
	}
}

// TestHelloChe function
func TestHelloChe(t *testing.T) {
	result := Hello("Che")
	if result != "Hello Che!" {
		t.Errorf("Failed > Hello(\"Che\"): %v ? %v", "Hello Che!", result)
	} else {
		t.Logf("Success > Hello(\"Che\"): %v ? %v", "Hello Che!", result)
	}
}
