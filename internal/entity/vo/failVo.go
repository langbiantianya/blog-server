package vo

type FailVo struct {
	Message string `json:"message,omitempty"`
}

func Fail(err error) FailVo {
	return FailVo{Message: err.Error()}
}
