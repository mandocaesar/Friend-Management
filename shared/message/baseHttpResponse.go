package message

//BaseHTTPResponse : struct for http response
type BaseHTTPResponse struct {
	Success int    `json:"success"`
	Data    string `json:"data"`
}
