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
	GetUser(ctx context.Context, username string) (*model.User, errors.E)

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

	// 13. GetAllComment : returns all comment on owners' profile
	// get owner's username from ctx
	GetAllComment(ctx context.Context) ([]model.Comment, errors.E)

	// 15. GetAllCurrentBook : returns all currently owned books
	// get owner's username from ctx
	GetAllCurrentBook(ctx context.Context) ([]model.Book, errors.E)

	// 16. GetAllOwnedBook : returns all owned books
	// get owner's username from ctx
	GetAllOwnedBook(ctx context.Context) ([]model.Book, errors.E)
}

// BookCore : core business logic responsible for managing book
type BookCore interface {

	// 1. Add will add book to the system
	// with owner and current owner to user making this request
	// get username from ctx
	Add(ctx context.Context, in AddBookInput) errors.E

	// 2. Get will return book with given isbn
	// isbn : of book
	Get(ctx context.Context, isbn string) (*model.Book, errors.E)

	// 3. Update: will update book
	Update(ctx context.Context, book model.Book) errors.E

	// 4. Delete : only owner of the book can delete
	// get username of owner from ctx
	Delete(ctx context.Context, isbn string) errors.E

	// 5. DeleteAll will delete all book whose owner and current owner is this user
	// get username from ctx
	DeleteAll(ctx context.Context) errors.E

	// 6. Comment will make comment on this book
	// isbn of book on which comment with content is made
	// get username of user makeing this comment from ctx
	Comment(ctx context.Context, isbn string, content string) errors.E

	// 7. Rate will rate book with isbn out of 10
	// get username of user making this rateing from ctx
	Rate(ctx context.Context, isbn string, rating uint) errors.E

	// 8. SearchBookByName : returns all book matching a name
	SearchBookByName(ctx context.Context, name string, pageNumber uint) ([]model.Book, errors.E)

	// 9. SearchBookByAuthor : returns all book by a author
	SearchBookByAuthor(ctx context.Context, author string, pageNumber uint) ([]model.Book, errors.E)

	// 10. SearchBookByGenere : returns all book with given genere
	SearchBookByGenere(ctx context.Context, genere []string, pageNumber uint) ([]model.Book, errors.E)

	// 11. GetAllOwned will give all the book owned by a user
	GetAllOwned(ctx context.Context, username string) ([]model.Book, errors.E)

	// 12. GetAllCurrent will give all the book currently owned by a user
	GetAllCurrent(ctx context.Context, username string) ([]model.Book, errors.E)

	// 13. GetAllComment returns all comment made on a book
	GetAllComment(ctx context.Context, isbn string) ([]model.Comment, errors.E)
}

// ImageCore : core business logic responsible for managing image
// not not directly called by client
// to be used by other cores
type ImageCore interface {
	// Create : store image to image database
	Create(ctx context.Context, img model.Image) errors.E

	// Get:
	// imgType : type of image to get
	// identifer : to whom this image is related
	// eg: imgType : profile and identifer can be usernam of a user
	Get(ctx context.Context, imgType model.ImageType, identifier string) (*model.Image, errors.E)
	Update(ctx context.Context, imgType model.ImageType, img model.Image) errors.E
	Delete(ctx context.Context, imgType model.ImageType, identifier string) errors.E

	// DeleteAll : delete all image which can be identified by given identifier
	DeleteAll(ctx context.Context, identifier string) errors.E
}

// CommentCore : core business logic responsible for managing comment
type CommentCore interface {
	Create(ctx context.Context, comment model.Comment) errors.E
	// Get:
	// commentType : type of comment to get
	// identifer : to whom this comment is related
	// eg: commentType : commentOnComment and identifer can id of coment
	// this will give list of comment made on a comment
	Get(ctx context.Context, commentType model.CommentType, identifer string) ([]model.Comment, errors.E)
	Update(ctx context.Context, commentType model.CommentType, identifer string, comment model.Comment) errors.E
	Delete(ctx context.Context, commentType model.CommentType, identifer string) errors.E

	// DeleteAll : delete all comment which can be identified by given identifier
	DeleteAll(ctx context.Context, identifer string) errors.E
}
