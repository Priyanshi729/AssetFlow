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

func GetAssetByID(assetID string) (*models.AssetDetails, error) {

	query := `
	SELECT
		a.asset_id,a.brand,a.model,a.serial_number,a.asset_type,a.status,a.owner_type,a.warranty_start,a.warranty_end,a.created_at,

		l.processor,
		COALESCE(l.ram, mb.ram) AS ram,
		COALESCE(l.storage, mb.storage) AS storage,
		COALESCE(l.operating_system, mb.operating_system) AS operating_system,
		COALESCE(l.charger, mb.charger) AS charger,
		COALESCE(l.device_password, mb.device_password) AS device_password,

		k.layout,
		COALESCE(k.connectivity, ms.connectivity) AS connectivity,

		ms.dpi

	FROM assets a

	LEFT JOIN laptops l
		ON a.asset_id = l.asset_id

	LEFT JOIN mobiles mb
		ON a.asset_id = mb.asset_id

	LEFT JOIN keyboards k
		ON a.asset_id = k.asset_id

	LEFT JOIN mouses ms
		ON a.asset_id = ms.asset_id

	WHERE a.asset_id = $1
	  AND a.archived_at IS NULL;
	`

	var asset models.AssetDetails

	if err := database.DB.Get(&asset, query, assetID); err != nil {
		return nil, err
	}

	return &asset, nil
}
