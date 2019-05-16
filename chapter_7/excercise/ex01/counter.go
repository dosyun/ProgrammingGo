/*
ex7-1
ByteCounterの考え方を利用して、ワードと行に対するカウンタを実装しなさい。
bufio.ScanWordが役立つでしょう。
*/
package counter

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}
