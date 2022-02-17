package bind

import "webview-example/model"

func ComplexBind() interface{} {
	return func() model.Complex {
		return model.Complex{
			User: model.User{
				Name: "John Doe",
				Age:  42,
			},
			Data: []string{"foo", "bar"},
		}
	}
}
