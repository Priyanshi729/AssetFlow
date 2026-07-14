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
