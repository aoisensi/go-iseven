package iseven

type ErrorResponse struct {
	Err string `json:"error"`
}

func (err *ErrorResponse) Error() string {
	return err.Err
}
