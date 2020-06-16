package main

import "context"

var cancelBefore = false

func main() {
	c, cCancel := context.WithCancel(context.Background())

	c1, cf1 := context.WithCancel(c)
	defer cf1()

	c2, cf2 := context.WithCancel(c)
	defer cf2()

	c11, cf11 := context.WithCancel(c1)
	defer cf11()

	c12, cf12 := context.WithCancel(c1)
	defer cf12()

	if cancelBefore {
		cCancel()
	}

	for k, c := range map[string]context.Context{`c1`: c1, `c11`: c11, `c12`: c12, `c2`: c2} {
		var s string
		if c.Err() != nil {
			s = `cancelled`
		} else {
			s = `not cancelled`
		}
		println(k + ` is ` + s)
	}

	if !cancelBefore {
		cCancel()
	}
}
