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

func UpdateAssetStatus(assetID string) error {
	query := `
	UPDATE assets
	SET
		status='assigned',
		updated_at=CURRENT_TIMESTAMP
	WHERE asset_id=$1
	`

	_, err := database.DB.Exec(query, assetID)
	
	return err
}
