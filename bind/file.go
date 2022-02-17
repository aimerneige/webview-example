package bind

import (
	"io/ioutil"
	"log"
)

func FileBind() interface{} {
	return func() []string {
		files, err := ioutil.ReadDir("/")
		result := make([]string, len(files))
		if err != nil {
			log.Fatal(err)
			return result
		}
		for _, file := range files {
			result = append(result, file.Name())
		}
		return result
	}
}
