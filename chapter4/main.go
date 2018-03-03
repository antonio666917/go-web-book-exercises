package main

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func main() {

}
