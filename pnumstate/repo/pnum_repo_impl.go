package repo

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	myBolt "github.com/he-he-heng/toktok-backend/config/database/bolt"
	"github.com/he-he-heng/toktok-backend/model"
)

type PnumStateRepo struct {
	bolt *myBolt.Database
}

func NewPnumStateRepository(b *myBolt.Database) *PnumStateRepo {
	return &PnumStateRepo{
		bolt: b,
	}
}

func (r *PnumStateRepo) Create(pnum string, pnumState *model.PnumStatus) (*model.PnumStatus, error) {
	err := r.bolt.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(myBolt.PnumStatusBucket))
		if err != nil {
			return model.ErrUnableToCreateBucket
		}

		jsonBytes, err := json.Marshal(pnumState)
		if err != nil {
			return model.ErrJsonMarshalNotWorking
		}

		b.Put([]byte(pnum), jsonBytes)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pnumState, nil
}

func (r *PnumStateRepo) ReadByPnum(pnum string) (pnumState *model.PnumStatus, _ error) {
	err := r.bolt.DB.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(myBolt.PnumStatusBucket))
		if err != nil {
			return model.ErrUnableToCreateBucket
		}

		jsonBytes := b.Get([]byte(pnum))
		err = json.Unmarshal(jsonBytes, pnumState)
		if err != nil {
			return model.ErrJsonUnmarshalNotWorking
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pnumState, nil
}

func (r *PnumStateRepo) UpdateByPnum(pnum string, pnumState *model.PnumStatus) (*model.PnumStatus, error) {
	err := r.bolt.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(myBolt.PnumStatusBucket))
		if err != nil {
			return model.ErrUnableToCreateBucket
		}

		jsonBytes, err := json.Marshal(pnumState)
		if err != nil {
			return model.ErrJsonMarshalNotWorking
		}

		b.Put([]byte(pnum), jsonBytes)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pnumState, nil
}
