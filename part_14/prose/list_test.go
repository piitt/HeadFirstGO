package prose

import "testing"

type testData struct {
	list []string // сегмент который передается функции JoinWithCommas
	want string   // строка которую должна вернуть функция JoinWithCommas
}

func TestJoinWithCommas(t *testing.T) {
	// создаем сегмент значений
	tests := []testData{
		{list: []string{}, want: ""},
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\" ,want \"%s\"", test.list, got, test.want)
		}
	}
}
