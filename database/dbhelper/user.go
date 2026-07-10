package dbhelper

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

func CreateUser(name, email, phoneNo, role, userType, password string) (string, error) {
	var userID string

	query := `
		INSERT INTO users (
			name,email,phone_no,role,user_type,password
		)
		VALUES (
			TRIM($1),
			TRIM($2),
			TRIM($3),
			$4,
			$5,
			$6
		)
		RETURNING user_id;
	`

	err := database.DB.QueryRow(
		query,
		name,
		email,
		phoneNo,
		role,
		userType,
		password,
	).Scan(&userID)

	if err != nil {
		return "", err
	}

	return userID, nil
}
