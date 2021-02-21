package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/blog"
	"github.com/Zzocker/bookolab/pkg/code"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/pkg/util"
	"github.com/Zzocker/bookolab/ports"
)

type userCore struct {
	uStore ports.UserDatastore
	lg     blog.Logger
}

func (u *userCore) Register(ctx context.Context, in UserRegisterInput, password string) errors.E {
	lg := util.LoggerFromCtx(ctx, u.lg)
	lg.Debugf("validating user register request")
	err := in.validate(password)
	if err != nil {
		lg.Errorf("failed to validate user register request : %v", err.Error())
		return err
	}
	user := in.toUser(password)
	lg.Debugf("storing registered user")
	err = u.uStore.Store(ctx, user)
	if err != nil {
		lg.Errorf("failed to store user : %v", err.Error())
		return err
	}
	return nil
}
func (u *userCore) GetUser(ctx context.Context, username string) (*model.User, errors.E) {
	lg := util.LoggerFromCtx(ctx, u.lg)
	if username == "" {
		lg.Errorf("empty username")
		return nil, errors.Init(fmt.Errorf("empty username"), code.CodeInvalidArgument, "empty username")
	}
	lg.Debugf("getting user %s from user datastore", username)
	user, err := u.uStore.Get(ctx, username)
	if err != nil {
		lg.Errorf("failed to get user : %v", err.Error())
		return nil, err
	}
	return user, nil
}
func (u *userCore) UpdateUser(ctx context.Context, reader io.Reader) errors.E {
	lg := util.LoggerFromCtx(ctx, u.lg)
	username := unWrapUsername(ctx)
	lg.Debugf("getting older version of userprofile")
	user, err := u.uStore.Get(ctx, username)
	if err != nil {
		lg.Errorf("failed to get older version of userprofile : %v", err.Error())
		return err
	}
	lg.Debugf("updating userprofile")
	jErr := json.NewDecoder(reader).Decode(user)
	if jErr != nil {
		lg.Debugf("failed to update : %v", jErr)
	}
	lg.Debugf("storing updated userprofile")
	err = u.uStore.Update(ctx, *user)
	if err != nil {
		lg.Errorf("failed to store updated profile : %v", err.Error())
		return err
	}
	return nil
}
func (u *userCore) DeleteUser(ctx context.Context) errors.E {
	lg := util.LoggerFromCtx(ctx, u.lg)
	// TODO
	// firstly check can we delete this profile
	// check if this user has a book for which is not a owner
	lg.Debugf("deleteing userprofile from userstore")
	err := u.uStore.Delete(ctx, unWrapUsername(ctx))
	if err != nil {
		lg.Errorf("failed to delete userprofile %v", err)
		return err
	}
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
func (u *userCore) GetUserWithName(ctx context.Context, name string, pageNumber int64) ([]model.User, errors.E) {
	lg := util.LoggerFromCtx(ctx, u.lg)
	if pageNumber < 1 {
		pageNumber = 1
	}
	lg.Debugf("getting user with name=%s for pageNumber=%d", name, pageNumber)
	users, err := u.uStore.Query(ctx, "created_on", map[string]interface{}{
		"details.name": map[string]interface{}{
			"$regex":   fmt.Sprintf("%s*", name),
			"$options": "i",
		},
	}, pageNumber)
	if err != nil {
		lg.Errorf("%v", err.Error())
		return nil, err
	}
	return users, nil
}
func (u *userCore) Comment(ctx context.Context, username string, comment string) errors.E {
	lg := util.LoggerFromCtx(ctx, u.lg)
	lg.Debugf("checking if user=%s exists", username)
	_, err := u.uStore.Get(ctx, username)
	if err != nil {
		lg.Debugf("user %s doesn't exists %v", err.Error())
		return err
	}
	lg.Debugf("create a new comment on user=%s", username)
	cmt := model.NewComment(unWrapUsername(ctx), username, model.CommentTypeOnUser, comment)
	lg.Debugf("storing new created comment")
	err = GetCommentCore().Create(ctx, cmt)
	if err != nil {
		lg.Errorf("failed to store comment : %v", err.Error())
		return err
	}
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
