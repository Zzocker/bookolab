package model

// ImageType : type of image
// eg : userprofile ,book cover, book_second
type ImageType uint8

// Image :
type Image struct {
	ID    string `json:"id" bson:"_id"`
	Owner string `json:"owner" bson:"owner"`
}
