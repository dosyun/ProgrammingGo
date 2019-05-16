package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		str string
		wc  WordCounter
		lc  LineCounter
	}{
		{"", 0, 0},
		{"あいうえお あいうえお  あいうえお", 3, 1},
		{`あいうえお　あいうえ あいうえ
お`, 4, 2},
		{`あいうえ
あいうえ
あいうえお`, 3, 3},
	}

	for _, test := range tests {
		var wc WordCounter
		fmt.Fprintf(&wc, test.str)
		if wc != test.wc {
			t.Errorf("WordCounter error: res[%d] testdata[%s] expected[%v]", wc, test.str, test.wc)
		}

		var lc LineCounter
		fmt.Fprintf(&lc, test.str)
		if lc != test.lc {
			t.Errorf("LineCounter error: res[%d] testdata[%s] expected[%v]", lc, test.str, test.lc)
		}
	}
}
