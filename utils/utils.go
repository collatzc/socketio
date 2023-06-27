package utils

import (
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var Json = func() jsoniter.API {
	// https://github.com/json-iterator/go/issues/244
	jsoniter.RegisterExtension(&BinaryAsStringExtension{})
	return jsoniter.ConfigCompatibleWithStandardLibrary
}()

func Debug(l ...interface{}) {
	if os.Getenv("DEBUG") == "1" {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)

		prefix := "[socketio]"

		log.Println(append([]interface{}{prefix}, l...)...)
	}
}
