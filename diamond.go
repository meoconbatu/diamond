package diamond

import "errors"

const testVersion = 1

func Gen(c byte) (d string, err error) {
	if c < 65 || c > 90 {
		return "", errors.New("Char out of range")
	}
	delta := 0
	l := int((c-65)*2 + 1)
	s := make([]byte, l)
	for i := c; i >= 65; i-- {
		for j := 0; j < l; j++ {
			s[j] = byte(' ')
		}
		s[0+delta], s[l-delta-1] = c, c
		delta += 1
		if d != "" {
			d += string(s) + "\n"
			d = string(s) + "\n" + d
		} else {
			d = string(s) + "\n"
		}
		c -= 1
	}
	return d, nil
}
