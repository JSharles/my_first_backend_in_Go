package error

// HandleError : Function to handle error case everywhere its needed
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
