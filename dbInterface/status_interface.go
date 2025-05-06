package dbInterface

import (
	"database/sql"
	"fmt"
)

func (r *DbRepo) CreateStatus(status *Status) (int, error) {
	query := `INSERT INTO statuses (name, color) VALUES (?, ?)`
	result, err := r.DB.Exec(query, status.StatusName, status.Color)
	if err != nil {
		return 0, fmt.Errorf("failed to create status: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get new status ID: %v", err)
	}
	return int(id), nil
}

func (r *DbRepo) GetStatus(id string) (*Status, error) {
	query := `SELECT id, name, color FROM statuses WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var status Status
	err := row.Scan(&status.StatusID, &status.StatusName, &status.Color)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("status with ID %s not found", id)
	} else if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *DbRepo) UpdateStatus(status *Status) error {
	query := `UPDATE statuses SET name = ?, color = ? WHERE id = ?`
	result, err := r.DB.Exec(query, status.StatusName, status.Color, status.StatusID)
	if err != nil {
		return fmt.Errorf("failed to update status: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not verify update: %v", err)
	}
	if affected == 0 {
		return fmt.Errorf("no status found with ID %d", status.StatusID)
	}
	return nil
}

func (r *DbRepo) DeleteStatus(id string) error {
	query := `DELETE FROM statuses WHERE id = ?`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete status: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not confirm deletion: %v", err)
	}
	if affected == 0 {
		return fmt.Errorf("no status found with ID %s", id)
	}
	return nil
}
