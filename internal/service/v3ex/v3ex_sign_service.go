package v3ex

type V3exRepo interface {
}

type V3exService struct {
	v3exRepo V3exRepo
}

func NewV3exService(
	v3exRepo V3exRepo,
) *V3exService {
	return &V3exService{
		v3exRepo: v3exRepo,
	}
}
