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

	// 7. Comment : to comment on user
	// username : of user on which is comment is made
	// comment : content of comment
	// get username of user making this comment from ctx
	Comment(ctx context.Context, username string, comment string) errors.E

	// 8. RateAsSeller : rate on user considering that user as seller
	// username : of rated user
	// rating : value out of 10
	// get username of user making this rating from ctx
	RateAsSeller(ctx context.Context, username string, rating uint) errors.E

	// 9. RateAsBorrower : rate on user considering that user as bollower
	// username : of rated user
	// rating : value out of 10
	// get username of user making this rating from ctx
	RateAsBorrower(ctx context.Context, username string, rating uint) errors.E

	// 10. UpdateProfile : updates profile of user
	// f : reader interface to reade image bytevalue
	// contentType : of image
	// sizeBytes : size of image in bytes
	// get owner's username from ctx
	UpdateProfile(ctx context.Context, f io.Reader, contentType string, sizeBytes int64) errors.E

	// 11. GetUserProfile : return profile image of user
	// io.Reader : raw bytes reader
	// string : content-type
	GetUserProfile(ctx context.Context, username string) (io.Reader, string, errors.E)

	// 12. UpdatePassword : update password of owner's userprofile
	// get owner's username from ctx 
	UpdatePassword(ctx context.Context, newPassword string) errors.E
}
