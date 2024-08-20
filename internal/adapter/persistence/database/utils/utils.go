package utils

import (
	"errors"
	"toktok-backend/internal/adapter/persistence/database/ent"
	"toktok-backend/internal/core/domain"
)

type changeable interface {
	string | []byte | domain.UserRoleType | domain.UserBanStateType | domain.AvatarSexType | domain.AvatarStateType | domain.RelationStateType | domain.RelationAlertStateType
}

func Changeable[T changeable](arg T) bool {
	return len(arg) == 0
}

type Candidates map[string]string

type FieldMap struct {
	DefaultField string
	Candidates   Candidates
}

func HandleOrderAndFilter(builder any, filter, order string, fieldMap FieldMap) error {
	orderFunc := ent.Asc
	if order == "desc" {
		orderFunc = ent.Desc
	}

	field, ok := fieldMap.Candidates[filter]
	if !ok {
		if filter == "" {
			field = fieldMap.DefaultField
		} else {
			return errors.New("the value of filter is not the value that is matched among the candidates")
		}
	}

	switch v := builder.(type) {
	case *ent.UserQuery:
		v.Order(orderFunc(field))

	case *ent.AvatarQuery:
		v.Order(orderFunc(field))

	default:
		return errors.New("the type of builder is invalid")
	}

	return nil
}
