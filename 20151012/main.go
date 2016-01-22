package main

import "C"
import (
	"bitbucket.org/binet/go-ffi/pkg/ffi"
	"log"
	//"net/http"
	//"unsafe"
)

func main() {
	//lib, err := ffi.NewLibrary("./libs/profile.so")
	lib, err := ffi.NewLibrary("./libs/hello.so")
	if err != nil {
		panic(err)
	}
	defer lib.Close()

	// I guess I am wrong type(ffi.Void) specified but I have no idea for this line.
	//fn, err := lib.Fct("Profile", ffi.Void, []ffi.Type{ffi.Void})
	fn, err := lib.Fct("Hello", ffi.Pointer, []ffi.Type{ffi.Void})
	if err != nil {
		panic(err)
	}
	i := 2
	log.Printf("&i: %#v\n", &i)
	fn(&i)
	// execute exported function
	//h := fn(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, World")
	//})
	//http.HandleFunc("/", h.Interface().(http.HandlerFunc))
	//http.ListenAndServe(":8080", nil)
}
