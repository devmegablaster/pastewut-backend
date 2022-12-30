package errors

import (
  "fmt"
)

var (
  InternalServerError = &Error{
    Code: "internal_server_error",
    Err: fmt.Errorf("Internal Server Error"),
  }

  InvalidPastewut = &Error{
    Code: "invalid_pastewut",
    Err: fmt.Errorf("Invalid Pastewut"),
  }

  InvalidPastewutCode = &Error{
    Code: "invalid_pastewut_code",
    Err: fmt.Errorf("Invalid Pastewut Code"),
  }
)
