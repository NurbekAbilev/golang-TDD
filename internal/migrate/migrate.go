package migrate

import "database/sql"

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS tasks (
				id UUID PRIMARY KEY,
				title TEXT NOT NULL,
				description TEXT,
				completed_at TIMESTAMP
			)
		`)
	if err != nil {
		return err
	}

	return nil
}
