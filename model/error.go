package model

type MyError struct {
	Err string `json:"error"`
}

func (me MyError) Error() string {
	return me.Err
}

var (
	ErrorForbiddenAccess = MyError{
		Err: "Forbidden Access!",
	}

	ErrorNotFound = MyError{
		Err: "Not Found!",
	}
)
