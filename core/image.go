package core

import (
	"context"

	"github.com/Zzocker/bookolab/model"
	"github.com/Zzocker/bookolab/pkg/errors"
	"github.com/Zzocker/bookolab/ports"
)

type imageCore struct {
	iStore ports.ImageStore
}

func (i *imageCore) Create(ctx context.Context, img model.Image) errors.E {
	return nil // TODO
}
func (i *imageCore) Get(ctx context.Context, imgType model.ImageType, identifier string) (*model.Image, errors.E) {
	return nil, nil // TODO
}
func (i *imageCore) Update(ctx context.Context, imgType model.ImageType, img model.Image) errors.E {
	return nil // TODO
}
func (i *imageCore) Delete(ctx context.Context, imgType model.ImageType, identifier string) errors.E {
	return nil // TODO
}
func (i *imageCore) DeleteAll(ctx context.Context, identifier string) errors.E {
	return nil // TODO
}
