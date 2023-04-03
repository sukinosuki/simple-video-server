package core

func PanicIfErr(err error) {

	if err != nil {
		panic(err)
	}
}

func MustBindForm[T any](c *Context) *T {
	var t T

	err := c.ShouldBind(&t)

	if err != nil {
		panic(err)
	}

	return &t
}
