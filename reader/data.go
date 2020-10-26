package reader

var HttpData chan *Request

func init() {
	HttpData = make(chan *Request, 100)
}
