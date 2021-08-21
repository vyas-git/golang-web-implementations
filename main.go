package main

import (
	ex "goweb-examples/examples"
)

func main() {
	//ex.Helloworld()
	//ex.CustomRouter()
	//ex.Viewtemplate()

	//obj := ex.SingleMiddleware{}

	//obj := ex.MultiMiddleware{}

	//obj := ex.AuthSession{}

	//obj := ex.JsonImplentation{}

	//obj := ex.SocketsImplementation{}

	//obj := ex.PasswordHashing{}
	obj := ex.UrlShortner{}

	obj.Run()
}
