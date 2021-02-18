package core

import (
	"context"
	"fmt"
	"io"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type userCore struct {
	uStore ports.UserDatastore
}

func (u *userCore) Register(ctx context.Context, in UserRegisterInput, password string) errors.E {
	err := in.validate(password)
	if err != nil {
		return err
	}
	user := in.toUser(password)
	return u.uStore.Store(ctx, user)
}
func (u *userCore) GetUser(ctx context.Context, username string) (*model.User, errors.E) {
	if username == "" {
		return nil, errors.Init(fmt.Errorf("empty username"), code.CodeInvalidArgument, "empty username")
	}
	return u.uStore.Get(ctx, username)
}
func (u *userCore) UpdateUser(ctx context.Context, reader io.Reader) errors.E {
	return nil
}
func (u *userCore) DeleteUser(ctx context.Context) errors.E {
	return nil
}
func (u *userCore) CheckCred(ctx context.Context, username, password string) errors.E {
	user, err := u.uStore.Get(ctx, username)
	if err != nil {
		return err
	}
	if user.Password != hash(password) {
		return errors.Init(fmt.Errorf("invalid password"), code.CodeUnauthorized, "invalid password")
	}
	return nil
}
func (u *userCore) GetUserWithName(ctx context.Context, name string) ([]model.User, errors.E) {
	return nil, nil
}
func (u *userCore) Comment(ctx context.Context, username string, comment string) errors.E {
	return nil
}
func (u *userCore) RateAsSeller(ctx context.Context, username string, rating uint) errors.E {
	return nil
}
func (u *userCore) RateAsBorrower(ctx context.Context, username string, rating uint) errors.E {
	return nil
}
func (u *userCore) UpdateProfile(ctx context.Context, f io.Reader, contentType string, sizeBytes int64) errors.E {
	return nil
}
func (u *userCore) GetUserProfile(ctx context.Context, username string) (io.Reader, string, errors.E) {
	return nil, "", nil
}
func (u *userCore) UpdatePassword(ctx context.Context, newPassword string) errors.E {
	return nil
}
func (u *userCore) GetAllComment(ctx context.Context) ([]model.Comment, errors.E) {
	return nil, nil
}
func (u *userCore) GetAllCurrentBook(ctx context.Context) ([]model.Book, errors.E) {
	return nil, nil
}
func (u *userCore) GetAllOwnedBook(ctx context.Context) ([]model.Book, errors.E) {
	return nil, nil
}
