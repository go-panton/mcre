package users

import "time"

//BadRequestError store the error in the struct
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

//ErrorResponse stores the code and message of the error in the struct
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
