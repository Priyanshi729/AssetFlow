package repository

import "AssetFlow/database"

func CreateRepair(assetID string) error {
	query := `INSERT INTO asset_repairs
		(asset_id)
		VALUES
		($1)
	`
	_, err := database.DB.Exec(query, assetID)

	return err
}

func CompleteRepair(assetID string) error {
	query := `
		UPDATE asset_repairs
		SET
			repaired_at = CURRENT_TIMESTAMP,
			updated_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
		  AND repaired_at IS NULL
		  AND archived_at IS NULL
	`

	_, err := database.DB.Exec(query, assetID)
	if err != nil {
		return err
	}

	return nil
}
