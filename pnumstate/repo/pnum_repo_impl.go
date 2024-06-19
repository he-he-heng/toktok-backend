package repo

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	mybolt "github.com/he-he-heng/toktok-backend/config/database/bolt"
	"github.com/he-he-heng/toktok-backend/model"
)

type PnumStateRepo struct {
	bolt *mybolt.Database
}

func NewPnumStateRepository(b *mybolt.Database) *PnumStateRepo {
	return &PnumStateRepo{
		bolt: b,
	}
}

func (r *PnumStateRepo) Create(pnum string, pnumState *model.PnumStatus) (*model.PnumStatus, error) {
	err := r.bolt.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(mybolt.PnumStatusBucket))

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
		b := tx.Bucket([]byte(mybolt.PnumStatusBucket))

		jsonBytes := b.Get([]byte(pnum))
		pnumState = &model.PnumStatus{}
		err := json.Unmarshal(jsonBytes, pnumState)
		if err != nil {
			fmt.Println(err)
			return model.ErrJsonUnmarshalNotWorking
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pnumState, nil
}
