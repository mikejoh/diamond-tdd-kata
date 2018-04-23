package diamond

/*
Diamond TDD kata.

Printed Diamond:

   A
  B B
 C   C
  B B
   A

*/

import (
	"strings"
	"testing"
)

/*
Helper functions
*/
func isEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("%s is not %s", actual, expected)
	}
}

func Contains(t *testing.T, subject string, s string) {
	if !strings.Contains(subject, s) {
		t.Errorf("%#v does not contain %#v", subject, s)
	}
}

func EndsWith(t *testing.T, subject string, s string) {
	start := len(subject) - len(s)
	if subject[start:] != s {
		t.Errorf("%#v does not end with %#v", subject, s)
	}
}

func ContainsNumOfChars(t *testing.T, subject string, s string, n int) {
	num := strings.Count(subject, s)
	if num != n {
		t.Errorf("Got %d %#v expected %d in %#v", num, s, n, subject)
	}
}

/*
Tests
*/
func TestDiamondA(t *testing.T) {
	isEquals(t, Diamond('A'), "A")
}
func TestDiamondB(t *testing.T) {
	Contains(t, Diamond('B'), "B")
	ContainsNumOfChars(t, Diamond('B'), "B", 2)
	ContainsNumOfChars(t, Diamond('B'), "A", 2)
	Contains(t, Diamond('B'), " ")
	EndsWith(t, Diamond('B'), " A")
}
func TestDiamondC(t *testing.T) {
	ContainsNumOfChars(t, Diamond('C'), "B", 4)
	Contains(t, Diamond('C'), "C")
}
func TestThatDimondsHaveDoubleAmountOfNewlinesMinus2(t *testing.T) {
	ContainsNumOfChars(t, Diamond('B'), "\n", 2)
	ContainsNumOfChars(t, Diamond('C'), "\n", 4)
}
