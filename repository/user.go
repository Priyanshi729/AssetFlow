package repository

import (
	"AssetFlow/database"
	"AssetFlow/models"
	"AssetFlow/utils"
)

func IsUserExists(email string) (bool, error) {
	var exists bool
	query := `
		SELECT COUNT(user_id) > 0
		FROM users
		WHERE email = TRIM($1)
		  AND archived_at IS NULL
	`

	err := database.DB.Get(&exists, query, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CreateUser(name, email, password, phoneNo, role, userType string) (string, error) {
	var userID string

	query := `
		INSERT INTO users (
			name,email,password,phone_no,role,user_type
		)VALUES (TRIM($1),TRIM($2),TRIM($3),$4,$5,$6)
		RETURNING user_id;
	`

	if err := database.DB.QueryRow(query, name, email, password, phoneNo, role, userType).Scan(&userID); err != nil {
		return "", err
	}

	return userID, nil
}

func GetUserByPassword(email, password string) (models.LoginResponse, error) {

	query := `
		SELECT
			user_id,password,role
		FROM users
		WHERE email = TRIM($1)
		  AND archived_at IS NULL
	`
	var user models.LoginResponse

	if err := database.DB.Get(&user, query, email); err != nil {
		return models.LoginResponse{}, err
	}

	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return models.LoginResponse{}, err
	}

	return user, nil
}

func GetUser(userID string) (*models.User, error) {

	query := `
		SELECT
			user_id,name,email,phone_no,role,user_type
		FROM users
		WHERE user_id = $1
		  AND archived_at IS NULL
	`
	var user models.User

	if err := database.DB.Get(&user, query, userID); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserAssets(userID string) ([]models.Asset, error) {

	query := `
	SELECT
		a.asset_id,a.brand,a.model,a.serial_number,a.asset_type,a.status,a.owner_type,a.warranty_start,a.warranty_end,a.created_at
	FROM assets a
	INNER JOIN asset_assignments aa
	ON a.asset_id = aa.asset_id
    WHERE aa.assigned_to = $1
	AND aa.returned_at IS NULL
	AND a.archived_at IS NULL
	AND aa.archived_at IS NULL
	`
	var assets []models.Asset
	if err := database.DB.Select(&assets, query, userID); err != nil {
		return nil, err
	}

	return assets, nil
}

func GetUserAssetByID(userID, assetID string) (*models.Asset, error) {

	query := `
	SELECT
	    a.asset_id,a.brand,a.model,a.serial_number,a.asset_type,a.status,a.owner_type,a.warranty_start,a.warranty_end,a.created_at
	FROM assets a
	INNER JOIN asset_assignments aa
		ON a.asset_id = aa.asset_id
	WHERE aa.assigned_to = $1
	  AND a.asset_id = $2
	  AND aa.returned_at IS NULL
	  AND aa.archived_at IS NULL
	  AND a.archived_at IS NULL
	`

	var asset models.Asset

	if err := database.DB.Get(&asset, query, userID, assetID); err != nil {
		return nil, err
	}

	return &asset, nil
}

func DeleteUser(userID string) error {

	query := `
		UPDATE users
		SET
			archived_at = NOW()
		WHERE user_id = $1
		  AND archived_at IS NULL
	`

	_, err := database.DB.Exec(query, userID)
	return err
}
