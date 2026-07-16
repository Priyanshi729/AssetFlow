package repository

import (
	"AssetFlow/database"
	"AssetFlow/models"

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
func GetAllAssetAssignmentHistory() ([]models.AssetAssignmentHistory, error) {

	var history []models.AssetAssignmentHistory

	query := `
	SELECT
		aa.assignment_id,
		a.asset_id,
		a.brand,
		a.model,
		a.asset_type,
		u.user_id,
		u.name,
		a.status,
		aa.assigned_on,
		aa.returned_at
	FROM asset_assignments aa
	JOIN users u
		ON aa.assigned_to = u.user_id
	JOIN assets a
		ON aa.asset_id = a.asset_id
	WHERE aa.archived_at IS NULL
	ORDER BY aa.assigned_on DESC;
	`

	if err := database.DB.Select(&history, query); err != nil {
		return nil, err
	}

	return history, nil
}

func GetAllAssignedAsset() ([]models.AssetAssignmentHistory, error) {

	var history []models.AssetAssignmentHistory

	query := `
	SELECT
		aa.assignment_id,
		a.asset_id,
		a.brand,
		a.model,
		a.asset_type,
		u.user_id,
		u.name,
		a.status,
		aa.assigned_on,
		aa.returned_at
	FROM asset_assignments aa
	JOIN users u
		ON aa.assigned_to = u.user_id
	JOIN assets a
		ON aa.asset_id = a.asset_id
	WHERE aa.archived_at IS NULL
	      AND aa.returned_at IS NULL
	ORDER BY aa.assigned_on DESC;
	`

	if err := database.DB.Select(&history, query); err != nil {
		return nil, err
	}

	return history, nil
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
