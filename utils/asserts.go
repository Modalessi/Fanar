package utils

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func ErrorAssert(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
