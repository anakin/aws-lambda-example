package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"

)

type Req struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}
type Resp struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: make(map[string]string),
	}
	r, _ := json.Marshal(Resp{
		Success: false,
		Message: "You can not accessed my APIs!",
	})
	resp.Body = string(r)
	resp.StatusCode = http.StatusOK

	rb := Req{}
	if err := json.Unmarshal([]byte(request.Body), &rb); err != nil {
		resp.Body = setError("get request error")
		return resp, nil
	}

	response, _ := json.Marshal(Resp{
		Success:   true,
		Message:   "You accessed my APIs!",
	})

	resp.Body = string(response)
	return resp, nil
}

func setError(msg string) string {

	body, _ := json.Marshal(Resp{
		Success: false,
		Message: msg,
	})
	return string(body)
}

func main() {
	lambda.Start(Handler)
}
