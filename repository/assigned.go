package repository

import (
	"AssetFlow/database"

	"github.com/jmoiron/sqlx"
)

func CreateAssetAssignment(db sqlx.Ext, assetID string, userID string) error {
	query := `INSERT INTO asset_assignments
	(
		asset_id,assigned_to
	)
	VALUES
	(
		$1,$2
	)
	`
	_, err := db.Exec(query, assetID, userID)

	return err
}

func UpdateAssetStatus(assetID, status string) error {
	query := `
		UPDATE assets
		SET
			status = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
		  AND archived_at IS NULL
	`
	_, err := database.DB.Exec(query, assetID, status)
	
	return err
}

func ReturnAsset(assetID string) error {
	query := `
		UPDATE asset_assignments
		SET
			returned_at = CURRENT_TIMESTAMP,updated_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
		  AND returned_at IS NULL
		  AND archived_at IS NULL
	`

	_, err := database.DB.Exec(query, assetID)
	if err != nil {
		return err
	}

	return nil
}
