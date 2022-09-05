package models

import "database/sql"

type Articles struct {
	Id          int            `json:"id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	Author      sql.NullString `json:"author"`
	CreatedAt   sql.NullString `json:"created_at"`
	UpdatedAt   sql.NullString `json:"updated_at"`
	Category    sql.NullString `json:"category"`
	NoOfViews   sql.NullInt32  `json:"no_of_views"`
}
