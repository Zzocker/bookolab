package core

import (
	"context"
	"io"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/errors"
)

// UserCore : core business logic responsible for managaing user profile
type UserCore interface {
	// 1. Register called by top layer (http,grpc) of this project
	Register(ctx context.Context, in UserRegisterInput) errors.E

	// 2. GetUser : return userprofile of user with given username
	GetUser(ctx context.Context, username string) errors.E

	// 3. UpdateUser : will update userprofile
	// Get username of owner from ctx
	// older version of userprofile will first be loaded into memory
	// then updated userprofile reader will layered on older version
	UpdateUser(ctx context.Context, reader io.Reader) errors.E

	// 4. DeleteUser : will delete owners's profile
	// Get username of owner from ctx
	DeleteUser(ctx context.Context) errors.E

	// 5. CheckCred : check credential of user
	CheckCred(ctx context.Context, username, password string) errors.E

	// 6. GetUserWithName : returns userprofiles with matching name patterns
	// Decending orderer of rating
	GetUserWithName(ctx context.Context, name string) ([]model.User, errors.E)
}
