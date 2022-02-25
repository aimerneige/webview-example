package bind

func UrlImageBind() interface{} {
	return func() string {
		return "https://go.dev/doc/gopher/gophercolor.png"
	}
}

func LocalImageBind() interface{} {
	return func() string {
		return "img/gophercolor.png"
	}
}
