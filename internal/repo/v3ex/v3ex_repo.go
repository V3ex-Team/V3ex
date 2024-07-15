package v3ex

import (
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/service/v3ex"
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
