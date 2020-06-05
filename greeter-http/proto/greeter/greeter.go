package greeter

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Msg string `json:"msg"`
}
