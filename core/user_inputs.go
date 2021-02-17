package core

import (
	"fmt"
	"time"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/go-playground/validator"
)

// user_inputs.go
// contains all inputs for userCore methods

// UserRegisterInput : input for register method
type UserRegisterInput struct {
	Username string           `json:"username"`
	Name     string           `json:"name"`
	DOB      int64            `json:"dob"`
	Gender   model.GenderType `json:"gender"`
	EmailID  string           `json:"email_id"`
}

func (u UserRegisterInput) validate(password string) errors.E {
	var err errors.E
	if u.Username == "" {
		err = errors.Init(fmt.Errorf("empty username"), code.CodeInvalidArgument, "empty username")
	} else if u.Name == "" {
		err = errors.Init(fmt.Errorf("empty name"), code.CodeInvalidArgument, "empty name")
	} else if u.DOB <= 0 {
		err = errors.Init(fmt.Errorf("invalid dob"), code.CodeInvalidArgument, "invalid dob")
	} else if password == "" {
		err = errors.Init(fmt.Errorf("empty password"), code.CodeInvalidArgument, "empty password")
	} else if u.EmailID == "" {
		err = errors.Init(fmt.Errorf("empty email ID"), code.CodeInvalidArgument, "empty email ID")
	} else if vErr := validator.New().Var(u.EmailID, "email"); vErr != nil {
		err = errors.Init(fmt.Errorf("invalid email ID"), code.CodeInvalidArgument, "invalid email ID")
	}

	if u.Gender == model.GenderTypeMale {
	} else if u.Gender == model.GenderTypeMale {
	} else {
		err = errors.Init(fmt.Errorf("invalid gender type"), code.CodeInvalidArgument, "invalid gender type")
	}
	return err
}

func (u UserRegisterInput) toUser(password string) model.User {
	return model.User{
		Username: u.Username,
		Details: model.UserDetails{
			Name:   u.Name,
			DOB:    u.DOB,
			Gender: u.Gender,
		},
		Contact: model.UserContact{
			EmailID: u.EmailID,
		},
		CreatedOn: time.Now().Unix(),
		Password:  hash(password),
	}
}
