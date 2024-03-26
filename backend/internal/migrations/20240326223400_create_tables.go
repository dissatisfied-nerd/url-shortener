package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTables, downCreateTables)
}

func upCreateTables(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS urls
	(
		id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,

		original  text,
		shortened text
	);`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateTables(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table order_products`)
	if err != nil {
		return err
	}

	return nil
}
