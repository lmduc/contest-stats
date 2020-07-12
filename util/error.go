package util

// PanicIfError will panic if the err is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
