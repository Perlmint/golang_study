package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (self rot13Reader) Read(p []byte) (n int, err error) {
	count := 0
	self.r.Read(p)
	for index, val := range p {
		if (val == 0) {
			return count, io.EOF
		}
		if (val > 64 && val < 91) {
			val = val + 13
			if val > 90 {
				val -= 26
			}
		} else if (val > 96 && val < 123) {
			val = val + 13
			if val > 122 {
				val -= 26
			}
		}
		
		p[index] = val
		
		count++
	}
	return count, nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

