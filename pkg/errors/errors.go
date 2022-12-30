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

  InvalidDetails = &Error{
    Code: "invalid_details",
    Err: fmt.Errorf("Invalid Details"),
  }

  InvalidEmail = &Error{
    Code: "invalid_email",
    Err: fmt.Errorf("Invalid Email"),
  }

  InvalidPassword = &Error{
    Code: "invalid_password",
    Err: fmt.Errorf("Invalid Password"),
  }

  UserNotFound = &Error{
    Code: "user_not_found",
    Err: fmt.Errorf("User Not Found"),
  }

  UserAlreadyExists = &Error{
    Code: "user_already_exists",
    Err: fmt.Errorf("User Already Exists"),
  }

  Unauthorized = &Error{
    Code: "unauthorized",
    Err: fmt.Errorf("Unauthorized"),
  }

  InvalidToken = &Error{
    Code: "invalid_token",
    Err: fmt.Errorf("Invalid Token"),
  }
)
