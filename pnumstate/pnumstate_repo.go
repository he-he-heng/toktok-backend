package pnumstate

import "github.com/he-he-heng/toktok-backend/model"

type PnumStateRepo interface {
	Create(pnum string, pnumState *model.PnumStatus) (*model.PnumStatus, error)
	ReadByPnum(pnum string) (*model.PnumStatus, error)
	UpdateByPnum(pnum string, pnumState *model.PnumStatus) (*model.PnumStatus, error)
}
