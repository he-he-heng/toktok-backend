package test

import (
	"toktok-backend/internal/adapter/persistence/mysql/repository"
	"toktok-backend/internal/adapter/persistence/mysql/test/provider"
)

var userRelation = repository.NewRelationRepository(provider.Get().Client)
