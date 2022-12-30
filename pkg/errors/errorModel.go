package errors

type Error struct {
  Code string `json:"code"`
  Err error `json:"error"`
  Message string `json:"message"`
}

func (e *Error) Error() Error {
  e.Message = e.Err.Error()
  return *e
}
