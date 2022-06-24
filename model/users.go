package model

import (
	"database/sql"
	"errors"
)

type Users struct {
	DB *DB
}

const insertUserQuery = `insert into users (uid, gid) values (?, ?)`

// Insert adds a user.
func (u *Users) Insert(gid, uid int64) error {
	_, err := u.DB.Exec(insertUserQuery, uid, gid)
	if err != nil {
		return err
	}
	return nil
}

const deleteUserQuery = `delete from users where gid = ? and uid = ?`

// Delete removes a user.
func (u *Users) Delete(gid, uid int64) error {
	_, err := u.DB.Exec(deleteUserQuery, gid, uid)
	if err != nil {
		return err
	}
	return nil
}

const listUsersQuery = `select uid from users where gid = ?`

// List returns all users.
func (u *Users) List(gid int64) ([]int64, error) {
	rows, err := u.DB.Query(listUsersQuery, gid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

const existsUserQuery = `select 1 from users where gid = ? and uid = ?`

// Exists returns true if the given user is in the table, false otherwise.
func (u *Users) Exists(gid, uid int64) (bool, error) {
	var i int
	if err := u.DB.QueryRow(existsUserQuery, gid, uid).Scan(&i); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return i == 1, nil
}

const randomUserQuery = `
select uid from users
where gid = ?
order by random()
limit 1
`

// Random returns a random user.
func (u *Users) Random(gid int64) (int64, error) {
	var uid int64
	if err := u.DB.QueryRow(randomUserQuery, gid).Scan(&uid); err != nil {
		return 0, err
	}
	return uid, nil
}

const nRandomUserQuery = `
select uid from users
where gid = ?
order by random()
limit ?`

// NRandom returns n random users. If there is not enough users in the table, returns less than n users.
func (u *Users) NRandom(gid int64, n int) ([]int64, error) {
	rows, err := u.DB.Query(nRandomUserQuery, gid, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}