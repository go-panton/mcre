package users

import "time"

type BadRequestError struct {
	Err error
}

func (err BadRequestError) Error() string {
	return err.Err.Error()
}

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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID      int    `json:"userid"`
	TokenString string `json:"tokenString"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
