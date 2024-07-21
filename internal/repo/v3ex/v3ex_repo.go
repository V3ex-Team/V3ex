package v3ex

import (
	"context"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/service/v3ex"
	"github.com/segmentfault/pacman/errors"
)

type v3exRepo struct {
	data *data.Data
}

// NewV3exRepo new repository
func NewV3exRepo(data *data.Data) v3ex.V3exRepo {
	return &v3exRepo{
		data: data,
	}
}

func (ur *v3exRepo) ListToken(ctx context.Context) (tokens []*entity.ERC20Token, err error) {
	tokens = make([]*entity.ERC20Token, 0)
	err = ur.data.DB.Context(ctx).Find(&tokens)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil, nil
}
func (ur *v3exRepo) AddToken(ctx context.Context, req *entity.ERC20Token) (err error) {
	_, err = ur.data.DB.Context(ctx).Insert(req)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (ur *v3exRepo) DeleteToken(ctx context.Context, id int64) (err error) {
	_, err = ur.data.DB.Context(ctx).Where("id = ?", id).Delete(&entity.ERC20Token{})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return err
}
func (ur *v3exRepo) UpdateToken(ctx context.Context, req *entity.ERC20Token) (err error) {
	_, err = ur.data.DB.Context(ctx).Cols("name", "symbol", "address", "decimals", "logo_uri", "description").Update(req, &entity.ERC20Token{ID: req.ID})
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}
