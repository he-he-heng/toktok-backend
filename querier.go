package main

import (
	"context"

	"github.com/he-he-heng/toktok-backend/ent"
	"github.com/he-he-heng/toktok-backend/ent/attempt"
)

type Querier struct {
	client *ent.Client
}

func (q *Querier) GetAttempt(ctx context.Context, p string) (*ent.Attempt, error) {
	return q.client.Attempt.Query().Where(attempt.Phone(p)).Only(ctx)
}

func (q *Querier) CreateAttempt(ctx context.Context, a *ent.Attempt) (*ent.Attempt, error) {
	return q.client.Attempt.Create().
		SetCnt(a.Cnt).
		SetBreak(a.Break).
		SetPhone(a.Phone).
		SetTimestamp(a.Timestamp).
		SetAuthcode(a.Authcode).
		Save(ctx)
}

func (q *Querier) UpdateAttempt(ctx context.Context, p string, a *ent.Attempt) (int, error) {
	return q.client.Attempt.Update().Where(attempt.Phone(p)).
		SetCnt(a.Cnt).
		SetBreak(a.Break).
		SetPhone(a.Phone).
		SetTimestamp(a.Timestamp).
		SetAuthcode(a.Authcode).
		Save(ctx)
}
