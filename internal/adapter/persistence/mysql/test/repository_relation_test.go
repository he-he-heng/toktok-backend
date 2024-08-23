package test

import (
	"testing"
	"toktok-backend/internal/adapter/persistence/mysql/repository"
	"toktok-backend/internal/adapter/persistence/mysql/test/provider"
)

var userRelation = repository.NewRelationRepository(provider.Get().Client)

func TestCreateRelation(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)
}
func TestGetRelation(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

}

func TestListRelation(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

}
func TestUpdateRelation(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

}

func TestDeleteRelation(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

}
