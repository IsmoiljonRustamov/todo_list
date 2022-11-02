package main

import (
	"database/sql"
	"fmt"
	"time"
)

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: *db}

}

type DBManager struct {
	db sql.DB
}

type TODO struct {
	id           int
	title        string
	descriptions string
	assignee     string
	status       bool
	deadline     time.Time
	created_at   time.Time
	updated_at   time.Time
	deleted_at   time.Time
}

type GetAllParam struct {
	limit int
	page int
	title string
	assignee string
}
	

func (b *DBManager) CreateToDo(td TODO) (*TODO, error) {
	query := `INSERT INTO to_do_list(
		title,
		descriptions,
		assignee,
		status,
		deadline
		) VALUES($1,$2,$3,$4,$5)
		RETURNING id,title,descriptions,assignee,status,deadline,created_at`

	row := b.db.QueryRow(
		query,
		td.title,
		td.descriptions,
		td.assignee,
		td.status,
		td.deadline,
	)

	var res TODO
	err := row.Scan(
		&res.id,
		&res.title,
		&res.descriptions,
		&res.assignee,
		&res.status,
		&res.deadline,
		&res.created_at,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (b *DBManager) Get(id int) (*TODO, error) {
	var res TODO

	query := `
	SELECT 
	id,
	title,
	description,
	assignee,
	status,
	deadline,
	created_at,`

	row :=b.db.QueryRow(query,id)

	err := row.Scan(
		&res.id,
		&res.title,
		&res.assignee,
		&res.status,
		&res.deadline,
		&res.created_at,
	)

	if err != nil {
		return nil, err
	}
	return &res,nil 

}

func (d *DBManager) Update (dt *TODO) (*TODO,error) {
	query := `UPDATE to_do_list SET 
		updated_at = $1
		WHERE id = $2
		RETURNING id,updated_at`

	row := d.db.QueryRow(
		query,
		dt.updated_at,
		dt.id,
	)	

	var res TODO 
	err := row.Scan(
		&res.id,
		&res.updated_at,
	)

	if err != nil {
		return nil,err
	}
	return dt,nil
}

func (d *DBManager) Delete (dt *TODO) {
	query := `UPDATE to_do_list SET deleted_at=$1 WHERE id=$2`

	d.db.Exec(
		query,
		dt.deleted_at,
		dt.id,
	)


}

func (d *DBManager) GetAll(gg *GetAllParam) ([]*TODO,error) {
	offset := (gg.page - 1) * gg.limit
	
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", gg.limit,offset)

	filter := "WHERE true and deleted_at is null"
	if gg.assignee != ""{ 
		filter += "AND assignee ILIKE '%" + gg.assignee + "%' "
	}

	if gg.title != "" {
		filter += " AND title ILIKE '%" + gg.title + "%'"
	}
	query := `
		SELECT 
		id,
		title,
		descriptions,
		assignee,
		status,
		deadline,
		created_at
		FROM to_do_list deleted_at 
		` + filter + `
		ORDER BY id DESC
		` + limit

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*TODO
	for rows.Next() {
		var u TODO

		err := rows.Scan(
			&u.id,
			&u.title,
			&u.descriptions,
			&u.assignee,
			&u.status,
			&u.deadline,
			&u.created_at,
		)
		if err != nil {
			return nil,err 
		}

		todos = append(todos, &u)
	}
	return todos,nil
}

