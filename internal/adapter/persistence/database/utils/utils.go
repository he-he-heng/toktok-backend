package utils

import (
	"toktok-backend/internal/core/domain"
)

type changeable interface {
	string | []byte | domain.UserRoleType | domain.UserBanStateType | domain.AvatarSexType | domain.AvatarStateType | domain.RelationStateType | domain.RelationAlertStateType
}

func Changeable[T changeable](arg T) bool {
	return len(arg) == 0
}

type Candidates map[string]string

type Criterions struct {
	DefaultField string
	Candidates   Candidates
}
