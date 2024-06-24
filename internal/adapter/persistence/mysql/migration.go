package mysql

import (
	"context"
	_ "toktok-backend/internal/adapter/persistence/mysql/ent/runtime"
)

func (d *Database) AutoMigration(ctx context.Context) error {
	if err := d.Schema.Create(ctx); err != nil {
		return err
	}

	return nil
}
