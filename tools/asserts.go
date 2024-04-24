package tools

func Assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}
