package store

func CreateUser(name, hashedPasswd, salt, email, phone string) error {
	sql := `INSERT INTO users (name, passwd_hash, salt, email, phone)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sql, name, hashedPasswd, salt, email, phone)
	if err != nil {
		return err
	}
	return nil
}
