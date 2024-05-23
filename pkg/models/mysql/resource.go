package mysql

import (
	"database/sql"
	"github.com/buhtigexa/renter/pkg/models"
)

type ResourceModel struct {
	DB *sql.DB
}

func (m *ResourceModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}
func (m *ResourceModel) Get(id int) (*models.Resource, error) {
	return nil, nil
}
func (m *ResourceModel) Latest() ([]*models.Resource, error) {
	return nil, nil
}
