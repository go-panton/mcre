package users

import "time"

type BadRequestError error

//SignUpRequest is a request struct
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SignUpResponse is a response struct
type SignUpResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	//UserID       int       `json:"userid"`
	//SessionToken string    `json:"sessiontoken"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
