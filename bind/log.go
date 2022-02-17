package bind

import "log"

func GoLogBind() interface{} {
	return func(data string) {
		log.Println(data)
	}
}
