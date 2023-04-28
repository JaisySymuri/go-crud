package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"go.starlark.net/lib/time"

	// "time"

	appd "appdynamics"

	_ "github.com/go-sql-driver/mysql"
)


type Employee struct {
	ID_Pegawai    int
	NIK           string
	Nama          string
	Username      string
	Password      string
	Alamat        string
	Tempat_Lahir  string
	Tanggal_Lahir time.Time
	No_HP         int
	Pekerjaan     string
	Gender        string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "handalin.com"
	dbName := "goblog2"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY ID_Pegawai DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var ID_Pegawai int
		var NIK string
		var Nama string
		var Username string
		var Password string
		var Alamat string
		var Tempat_Lahir string
		var Tanggal_Lahir time.Time
		var No_HP int
		var Pekerjaan string
		var Gender string
		var updated_at string
		var created_at string

		err = selDB.Scan(&ID_Pegawai, &NIK, &Nama, &Username, &Password, &Alamat, &Tempat_Lahir, &Tanggal_Lahir, &No_HP, &Pekerjaan, &Gender, &updated_at, &created_at)
		if err != nil {
			panic(err.Error())
		}
		emp.ID_Pegawai = ID_Pegawai
		emp.NIK = NIK
		emp.Nama = Nama
		emp.Username = Username
		emp.Password = Password
		emp.Alamat = Alamat
		emp.Tempat_Lahir = Tempat_Lahir
		emp.Tanggal_Lahir = Tanggal_Lahir
		emp.No_HP = No_HP
		emp.Pekerjaan = Pekerjaan
		emp.Gender = Gender

		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID_Pegawai := r.URL.Query().Get("ID_Pegawai")
	selDB, err := db.Query("SELECT * FROM Employee WHERE ID_Pegawai=?", nID_Pegawai)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var ID_Pegawai int
		var NIK string
		var Nama string
		var Username string
		var Password string
		var Alamat string
		var Tempat_Lahir string
		var Tanggal_Lahir time.Time
		var No_HP int
		var Pekerjaan string
		var Gender string
		var updated_at string
		var created_at string

		err = selDB.Scan(&ID_Pegawai, &NIK, &Nama, &Username, &Password, &Alamat, &Tempat_Lahir, &Tanggal_Lahir, &No_HP, &Pekerjaan, &Gender, &updated_at, &created_at)

		if err != nil {
			panic(err.Error())
		}
		emp.ID_Pegawai = ID_Pegawai
		emp.NIK = NIK
		emp.Nama = Nama
		emp.Username = Username
		emp.Password = Password
		emp.Alamat = Alamat
		emp.Tempat_Lahir = Tempat_Lahir
		emp.Tanggal_Lahir = Tanggal_Lahir
		emp.No_HP = No_HP
		emp.Pekerjaan = Pekerjaan
		emp.Gender = Gender
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID_Pegawai := r.URL.Query().Get("ID_Pegawai")
	selDB, err := db.Query("SELECT * FROM Employee WHERE ID_Pegawai=?", nID_Pegawai)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var ID_Pegawai int
		var NIK string
		var Nama string
		var Username string
		var Password string
		var Alamat string
		var Tempat_Lahir string
		var Tanggal_Lahir time.Time
		var No_HP int
		var Pekerjaan string
		var Gender string
		var updated_at string
		var created_at string

		err = selDB.Scan(&ID_Pegawai, &NIK, &Nama, &Username, &Password, &Alamat, &Tempat_Lahir, &Tanggal_Lahir, &No_HP, &Pekerjaan, &Gender, &updated_at, &created_at)

		if err != nil {
			panic(err.Error())
		}
		emp.ID_Pegawai = ID_Pegawai
		emp.NIK = NIK
		emp.Nama = Nama
		emp.Username = Username
		emp.Password = Password
		emp.Alamat = Alamat
		emp.Tempat_Lahir = Tempat_Lahir
		emp.Tanggal_Lahir = Tanggal_Lahir
		emp.No_HP = No_HP
		emp.Pekerjaan = Pekerjaan
		emp.Gender = Gender
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		NIK := r.FormValue("NIK")
		Nama := r.FormValue("Nama")
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")
		Alamat := r.FormValue("Alamat")
		Tempat_Lahir := r.FormValue("Tempat_Lahir")
		Tanggal_Lahir := r.FormValue("Tanggal_Lahir")
		No_HP := r.FormValue("No_HP")
		Pekerjaan := r.FormValue("Pekerjaan")
		Gender := r.FormValue("Gender")
		insForm, err := db.Prepare("INSERT INTO Employee(NIK, Nama, Username, Password, Alamat, Tempat_Lahir, Tanggal_Lahir, No_HP, Pekerjaan, Gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(NIK, Nama, Username, Password, Alamat, Tempat_Lahir, Tanggal_Lahir, No_HP, Pekerjaan, Gender)
		log.Println("INSERT: NIK: " + NIK + " | Nama: " + Nama + " | Username: " + Username + " | Password: " + Password + " | Alamat: " + Alamat + " | Tempat_Lahir: " + Tempat_Lahir + " | Tanggal_Lahir: " + Tanggal_Lahir + " | No_HP: " + No_HP + " | Pekerjaan: " + Pekerjaan + " | Gender: " + Gender)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		NIK := r.FormValue("NIK")
		Nama := r.FormValue("Nama")
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")
		Alamat := r.FormValue("Alamat")
		Tempat_Lahir := r.FormValue("Tempat_Lahir")
		Tanggal_Lahir := r.FormValue("Tanggal_Lahir")
		No_HP := r.FormValue("No_HP")
		Pekerjaan := r.FormValue("Pekerjaan")
		Gender := r.FormValue("Gender")
		ID_Pegawai := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET NIK=?, Nama=?, Username=?, Password=?, Alamat=?, Tempat_Lahir=?, Tanggal_Lahir=?, No_HP=?, Pekerjaan=?, Gender=? WHERE ID_Pegawai=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(NIK, Nama, Username, Password, Alamat, Tempat_Lahir, Tanggal_Lahir, No_HP, Pekerjaan, Gender, ID_Pegawai)
		log.Println("UPDATE: NIK: " + NIK + " | Nama: " + Nama + " | Username: " + Username + " | Password: " + Password + " | Alamat: " + Alamat + " | Tempat_Lahir: " + Tempat_Lahir + " | Tanggal_Lahir: " + Tanggal_Lahir + " | No_HP: " + No_HP + " | Pekerjaan: " + Pekerjaan + " | Gender: " + Gender)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("ID_Pegawai")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE ID_Pegawai=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:9999")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":9999", nil)

	cfg := appd.Config{}
 
cfg.AppName = "golang-crud"
cfg.TierName = "golang"
cfg.NodeName = "vm-jaisy"
cfg.Controller.Host = "pthandalinformasiteknologi-nfr.saas.appdynamics.com"
cfg.Controller.Port = 443
cfg.Controller.UseSSL = true
cfg.Controller.Account = "pthandalinformasiteknologi-nfr"
cfg.Controller.AccessKey = "cw1cyt3chkw8"
cfg.InitTimeoutMs = 10000  // Wait up to 10s for initialization to finish

if err := appd.InitSDK(&cfg); err != nil {
	fmt.Printf("Error initializing the AppDynamics SDK\n")
} else {
	fmt.Printf("Initialized AppDynamics SDK successfully\n")
}

// start the "Checkout" transaction
btHandle := appd.StartBT("Checkout", "")

// Optionally store the handle in the global registry
appd.StoreBT(btHandle, my_bt_guid)
... 

// Retrieve a stored handle from the global registry
myBtHandle = appd.GetBT(my_bt_guid)

// end the transaction
appd.EndBT(btHandle)
}
