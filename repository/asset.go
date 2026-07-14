package repository

import (
	"AssetFlow/database"
	"AssetFlow/models"

	"github.com/jmoiron/sqlx"
)

func CreateAsset(db sqlx.Ext, brand, model, serialNumber, assetType, status, ownerType, warrantyStart, warrantyEnd string) (string, error) {

	var assetID string

	query := `
		INSERT INTO assets (brand,model,serial_number,asset_type,status,owner_type,warranty_start,warranty_end)
		VALUES (TRIM($1),TRIM($2),TRIM($3),$4,$5,$6,$7,$8)
		RETURNING asset_id`

	err := sqlx.Get(db, &assetID, query, brand, model, serialNumber, assetType, status, ownerType, warrantyStart, warrantyEnd)
	if err != nil {
		return "", err
	}

	return assetID, nil
}

func CreateLaptop(db sqlx.Ext, assetID, processor, ram, storage, operatingSystem string, charger bool, devicePassword string) error {

	query := `
		INSERT INTO laptops (asset_id,processor,ram,storage,operating_system,charger,device_password)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := db.Exec(query, assetID, processor, ram, storage, operatingSystem, charger, devicePassword)

	return err
}

func CreateMobile(db sqlx.Ext, assetID, operatingSystem, ram, storage string, charger bool, devicePassword string) error {

	query := `
		INSERT INTO mobiles (asset_id,operating_system,ram,storage,charger,device_password)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := db.Exec(query, assetID, operatingSystem, ram, storage, charger, devicePassword)

	return err
}

func CreateKeyboard(db sqlx.Ext, assetID, layout, connectivity string) error {

	query := `
		INSERT INTO keyboards (asset_id,layout,connectivity)
		VALUES ($1,$2,$3)
	`

	_, err := db.Exec(query, assetID, layout, connectivity)

	return err
}

func CreateMouse(db sqlx.Ext, assetID string, dpi int, connectivity string) error {

	query := `
		INSERT INTO mouses (asset_id,dpi,connectivity)
		VALUES ($1,$2,$3)
	`

	_, err := db.Exec(query, assetID, dpi, connectivity)

	return err
}

func GetAssets() ([]models.Asset, error) {

	query := `
		SELECT
			asset_id,brand,model,serial_number,asset_type,status,owner_type,warranty_start,warranty_end,created_at
		FROM assets
		WHERE archived_at IS NULL
		ORDER BY created_at DESC
	`
	var assets []models.Asset

	if err := database.DB.Select(&assets, query); err != nil {
		return nil, err
	}

	return assets, nil
}
