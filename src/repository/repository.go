package repository

import "fmt"

type MainRepository struct {
	data string
}

func (m MainRepository) WriteData(new_data string) error {
	m.data += new_data
	return nil
}

func NewMainRepository() *MainRepository {
	return &MainRepository{}
}

func (m MainRepository) ReadData() (string, error) {
	if m.data == "" {
		return "", fmt.Errorf("Database has no data!")
	}
	return m.data, nil
}

// NOT IMPLEMENTED
func (m MainRepository) UpdateData() error {
	return nil
}

func (m MainRepository) CleanData() error {
	m.data = ""
	return nil
}
