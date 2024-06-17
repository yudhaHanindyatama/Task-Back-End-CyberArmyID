package homecontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"task-back-end-cyberarmyid/entities"
	"task-back-end-cyberarmyid/models/indexmodel"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	mahasiswa, err := indexmodel.GetAllWithHighestScore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Input(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/home/add.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		kelas := r.FormValue("kelas")
		matematika, _ := strconv.Atoi(r.FormValue("matematika"))
		kimia, _ := strconv.Atoi(r.FormValue("kimia"))
		fisika, _ := strconv.Atoi(r.FormValue("fisika"))
		biologi, _ := strconv.Atoi(r.FormValue("biologi"))

		mahasiswa := entities.Mahasiswa{
			Nama:  nama,
			Kelas: kelas,
		}

		err := indexmodel.AddMahasiswa(mahasiswa, matematika, kimia, fisika, biologi)
		if err != nil {
			http.Error(w, "Unable to save data", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func ListKelas(w http.ResponseWriter, r *http.Request) {
	mahasiswa, err := indexmodel.GetByClass()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/home/listKelas.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mahasiswa, err := indexmodel.GetMahasiswaDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mahasiswa.TotalNilai = (mahasiswa.Matematika + mahasiswa.Kimia + mahasiswa.Fisika + mahasiswa.Biologi) / 4

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/home/detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
