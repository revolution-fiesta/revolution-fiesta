package store

func CreateUser(name, hashedPasswd, salt, email, phone string) error {
	sql := `INSERT INTO users (name, passwd_hash, salt, email, phone)
VALUES ($1, $2, $3, $4, $5)`
	_, err := Db.Exec(sql, name, hashedPasswd, salt, email, phone)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(name string) (salt string, passwd_hash string, id int, err error) {
	query := "SELECT salt, passwd_hash,id FROM users WHERE name = $1"
	err = Db.QueryRow(query, name).Scan(&salt, &passwd_hash, &id)
	return
}
