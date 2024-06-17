package indexmodel

import (
	"task-back-end-cyberarmyid/config"
	"task-back-end-cyberarmyid/entities"
)

func GetAllWithHighestScore() ([]entities.Mahasiswa, error) {
	query := `
		SELECT m.id_mahasiswa, m.nama, m.kelas, 
		GREATEST(n.matematika, n.kimia, n.fisika, n.biologi) as nilai
		FROM mahasiswa m
		JOIN nilai_mahasiswa n ON m.id_mahasiswa = n.id_mahasiswa
		ORDER BY nilai DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var indexList []entities.Mahasiswa
	for rows.Next() {
		var index entities.Mahasiswa
		err := rows.Scan(&index.Id_mahasiswa, &index.Nama, &index.Kelas, &index.Nilai)
		if err != nil {
			return nil, err
		}
		indexList = append(indexList, index)
	}
	return indexList, nil
}

func AddMahasiswa(mahasiswa entities.Mahasiswa, matematika, kimia, fisika, biologi int) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	queryMahasiswa := "INSERT INTO mahasiswa (nama, kelas) VALUES (?, ?)"
	result, err := tx.Exec(queryMahasiswa, mahasiswa.Nama, mahasiswa.Kelas)
	if err != nil {
		tx.Rollback()
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	queryNilai := "INSERT INTO nilai_mahasiswa (id_mahasiswa, matematika, kimia, fisika, biologi) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.Exec(queryNilai, lastInsertId, matematika, kimia, fisika, biologi)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetByClass() ([]entities.MahasiswaGrouped, error) {
	query := `
		SELECT m.id_mahasiswa, m.nama, m.kelas, 
		GREATEST(n.matematika, n.kimia, n.fisika, n.biologi) as nilai
		FROM mahasiswa m
		JOIN nilai_mahasiswa n ON m.id_mahasiswa = n.id_mahasiswa
		ORDER BY m.kelas, nilai DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupedData []entities.MahasiswaGrouped
	var currentClass string
	var currentGroup entities.MahasiswaGrouped

	for rows.Next() {
		var index entities.Mahasiswa
		err := rows.Scan(&index.Id_mahasiswa, &index.Nama, &index.Kelas, &index.Nilai)
		if err != nil {
			return nil, err
		}

		if index.Kelas != currentClass {
			if currentClass != "" {
				groupedData = append(groupedData, currentGroup)
			}
			currentClass = index.Kelas
			currentGroup = entities.MahasiswaGrouped{
				Kelas:      index.Kelas,
				Mahasiswas: []entities.Mahasiswa{},
				TotalNilai: 0,
			}
		}
		currentGroup.Mahasiswas = append(currentGroup.Mahasiswas, index)
		currentGroup.TotalNilai += index.Nilai
	}

	if currentClass != "" {
		groupedData = append(groupedData, currentGroup)
	}

	return groupedData, nil
}

func GetMahasiswaDetail(id int) (entities.MahasiswaDetail, error) {
	query := `
		SELECT m.id_mahasiswa, m.nama, m.kelas, 
		n.matematika, n.kimia, n.fisika, n.biologi
		FROM mahasiswa m
		JOIN nilai_mahasiswa n ON m.id_mahasiswa = n.id_mahasiswa
		WHERE m.id_mahasiswa = ?
	`
	row := config.DB.QueryRow(query, id)

	var detail entities.MahasiswaDetail
	err := row.Scan(&detail.Id_mahasiswa, &detail.Nama, &detail.Kelas, &detail.Matematika, &detail.Kimia, &detail.Fisika, &detail.Biologi)
	if err != nil {
		return entities.MahasiswaDetail{}, err
	}

	return detail, nil
}
