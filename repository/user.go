package repository

import "AssetFlow/database"

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
