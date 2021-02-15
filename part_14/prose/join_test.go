package prose

import (
	"fmt"
	"testing"
)

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	// ожидаемый результат
	want := "apple"
	// фактический
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	// ожидаемый результат
	want := "apple and orange"
	// фактический
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	// ожидаемый результат
	want := "apple, orange, and pear"
	// фактический
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\" ", list, got, want)
}
