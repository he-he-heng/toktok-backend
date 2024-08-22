package database

import "context"

func AutoMigration(ctx context.Context, client *Client) error {
	err := client.Schema.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}
