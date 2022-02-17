package bind

func GetStringBind() interface{} {
	return func() string {
		return "Hello, World!"
	}
}
