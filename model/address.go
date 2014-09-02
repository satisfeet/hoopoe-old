package model

import "github.com/satisfeet/hoopoe/model/store"

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

func insertCity(e store.Execer, city string) (int64, error) {
	res, err := e.Exec(sqlInsertCity, city)

	return res.Id, err
}

func insertAddress(e store.Execer, m Address) (int64, error) {
	id, err := insertCity(e, m.City)
	if err != nil {
		return 0, err
	}

	res, err := e.Exec(sqlInsertAddress, m.Street, m.Code, id)

	return res.Id, err
}

func updateAddress(e store.Execer, m Address, cid interface{}) error {
	id, err := insertCity(e, m.City)
	if err != nil {
		return err
	}

	_, err = e.Exec(sqlUpdateAddress, m.Street, m.Code, id, cid)

	return err
}
