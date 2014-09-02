package store

type Address struct {
	Street string
	City   string
	Code   int
}

var sqlInsertCity = `
	INSERT INTO city (name)
	VALUES (?)
	ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
`

var sqlInsertAddress = `
	INSERT INTO address (street, code, city_id)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
`

var sqlUpdateAddress = `
	UPDATE address
	SET street = ?, code = ?, city_id = ?
	WHERE id=(SELECT address_id FROM customer WHERE id = ?)
`
