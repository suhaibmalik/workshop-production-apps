package jsend

import "net/http"

func Ok(data interface{}) JSend {
	return JSend{
		StatusCode: http.StatusOK,
		Status:     "success",
		Data:       data,
	}
}

func Created(data interface{}) JSend {
	return JSend{
		StatusCode: http.StatusCreated,
		Status:     "success",
		Data:       data,
	}
}

func Malformed() JSend {
	return JSend{
		StatusCode: http.StatusBadRequest,
		Status:     "fail",
		Data:       "malformed syntax",
	}
}

func ServerError() JSend {
	return JSend{
		StatusCode: http.StatusInternalServerError,
		Status:     "error",
		Message:    "internal server error",
	}
}
