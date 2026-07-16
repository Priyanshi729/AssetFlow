package repository

import (
	"AssetFlow/database"
	"AssetFlow/models"

	"github.com/jmoiron/sqlx"
)

func CreateAsset(db sqlx.Ext, req models.CreateAssetRequest) (string, error) {

	var assetID string

	query := `
		INSERT INTO assets (brand,model,serial_number,asset_type,status,owner_type,warranty_start,warranty_end)
		VALUES (TRIM($1),TRIM($2),TRIM($3),$4,$5,$6,$7,$8)
		RETURNING asset_id`

	err := sqlx.Get(db, &assetID, query, req.Brand, req.Model, req.SerialNumber, req.AssetType, req.Status, req.OwnerType, req.WarrantyStart, req.WarrantyEnd)
	if err != nil {
		return "", err
	}

	return assetID, nil
}

func CreateLaptop(db sqlx.Ext, assetID string, req models.LaptopRequestSpecific) error {

	query := `
		INSERT INTO laptops (asset_id,processor,ram,storage,operating_system,charger,device_password)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := db.Exec(query, assetID, req.Processor, req.RAM, req.Storage, req.OperatingSystem, req.Charger, req.DevicePassword)

	return err
}

func CreateMobile(db sqlx.Ext, assetID string, req models.MobileRequestSpecific) error {

	query := `
		INSERT INTO mobiles (asset_id,operating_system,ram,storage,charger,device_password)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := db.Exec(query, assetID, req.OperatingSystem, req.RAM, req.Storage, req.Charger, req.DevicePassword)

	return err
}

func CreateKeyboard(db sqlx.Ext, assetID string, req models.KeyboardRequestSpecific) error {

	query := `
		INSERT INTO keyboards (asset_id,layout,connectivity)
		VALUES ($1,$2,$3)
	`

	_, err := db.Exec(query, assetID, req.Layout, req.Connectivity)

	return err
}

func CreateMouse(db sqlx.Ext, assetID string, req models.MouseRequestSpecific) error {

	query := `
		INSERT INTO mouses (asset_id,dpi,connectivity)
		VALUES ($1,$2,$3)
	`

	_, err := db.Exec(query, assetID, req.DPI, req.Connectivity)

	return err
}

func GetAssets() ([]models.Asset, error) {

	query := `
		SELECT
			asset_id,brand,model,serial_number,asset_type,status,owner_type,warranty_start,warranty_end
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

func GetAssetByID(assetID string) (*models.Asset, error) {

	query := `SELECT
		asset_id,brand,model,serial_number,asset_type,status,owner_type,warranty_start,warranty_end
	FROM assets
	WHERE asset_id = $1
	  AND archived_at IS NULL
	`
	var asset models.Asset

	if err := database.DB.Get(&asset, query, assetID); err != nil {
		return nil, err
	}

	return &asset, nil
}

func GetLaptopByID(assetID string) (*models.LaptopRequestSpecific, error) {
	query := `SELECT
              processor,ram,storage,operating_system,charger,device_password
              FROM laptops
              WHERE asset_id = $1;
              `
	var laptop models.LaptopRequestSpecific

	if err := database.DB.Get(&laptop, query, assetID); err != nil {
		return nil, err
	}
	return &laptop, nil
}

func GetMobileByID(assetID string) (*models.MobileRequestSpecific, error) {
	query := `SELECT
              ram,storage,operating_system,charger,device_password
              FROM mobiles
              WHERE asset_id = $1;
              `
	var mobile models.MobileRequestSpecific

	if err := database.DB.Get(&mobile, query, assetID); err != nil {
		return nil, err
	}
	return &mobile, nil
}

func GetKeyboardByID(assetID string) (*models.KeyboardRequestSpecific, error) {
	query := `SELECT
              layout,connectivity
              FROM keyboards
              WHERE asset_id = $1;
              `
	var keyboard models.KeyboardRequestSpecific

	if err := database.DB.Get(&keyboard, query, assetID); err != nil {
		return nil, err
	}
	return &keyboard, nil
}

func GetMouseByID(assetID string) (*models.MouseRequestSpecific, error) {
	query := `SELECT
              dpi,connectivity
              FROM mouses
              WHERE asset_id = $1;`

	var mouse models.MouseRequestSpecific
	if err := database.DB.Get(&mouse, query, assetID); err != nil {
		return nil, err
	}

	return &mouse, nil
}

func UpdateAsset(db sqlx.Ext, assetID string, req models.UpdateAssetRequest) error {

	query := `
	UPDATE assets
	SET
		brand = COALESCE($2,brand),model = COALESCE($3,model),serial_number = COALESCE($4,serial_number),status = COALESCE($5,status),
		owner_type = COALESCE($6,owner_type),warranty_start = COALESCE($7,warranty_start),warranty_end = COALESCE($8,warranty_end),updated_at = CURRENT_TIMESTAMP

	WHERE asset_id = $1
	  AND archived_at IS NULL
	`
	_, err := db.Exec(query, assetID, req.Brand, req.Model, req.SerialNumber, req.Status, req.OwnerType, req.WarrantyStart, req.WarrantyEnd)

	return err
}

func UpdateLaptop(db sqlx.Ext, assetID string, req *models.UpdateLaptopRequest) error {

	query := `UPDATE laptops
              SET
	            processor = COALESCE($2, processor),ram = COALESCE($3, ram),storage = COALESCE($4, storage),
	            operating_system = COALESCE($5, operating_system),charger = COALESCE($6, charger),device_password = COALESCE($7, device_password)
              WHERE asset_id = $1
	`
	_, err := db.Exec(query, assetID, req.Processor, req.RAM, req.Storage, req.OperatingSystem, req.Charger, req.DevicePassword)

	return err
}

func UpdateMobile(db sqlx.Ext, assetID string, req *models.UpdateMobileRequest) error {

	query := `UPDATE mobiles
              SET
	             ram = COALESCE($2, ram),storage = COALESCE($3, storage),operating_system = COALESCE($4, operating_system),charger = COALESCE($5, charger),device_password = COALESCE($6, device_password)
              WHERE asset_id = $1
`
	_, err := db.Exec(query, assetID, req.RAM, req.Storage, req.OperatingSystem, req.Charger, req.DevicePassword)
	return err
}

func UpdateKeyboard(db sqlx.Ext, assetID string, req *models.UpdateKeyboardRequest) error {

	query := `UPDATE keyboards
              SET
	            layout = COALESCE($2, layout),connectivity = COALESCE($3, connectivity)
              WHERE asset_id = $1
`
	_, err := db.Exec(query, assetID, req.Layout, req.Connectivity)
	return err
}

func UpdateMouse(db sqlx.Ext, assetID string, req *models.UpdateMouseRequest) error {

	query := `UPDATE mouses
            SET
	            dpi = COALESCE($2, dpi),connectivity = COALESCE($3, connectivity)
            WHERE asset_id = $1
`

	_, err := db.Exec(query, assetID, req.DPI, req.Connectivity)
	return err
}

func DeleteAsset(assetID string) error {

	query := `
		UPDATE assets
		SET
			archived_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
		  AND archived_at IS NULL
	`

	_, err := database.DB.Exec(query, assetID)
	if err != nil {
		return err
	}

	return nil
}
