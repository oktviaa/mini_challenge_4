package main

import (
	"fmt"
	"html/template"
	"net/http"
)
	
	type Biodata struct {Nama, Email,Alamat,Pekerjaan,Alasan string}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", routeIndexGet)
    http.HandleFunc("/process", routeSubmitPost)

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
        var err = tmpl.Execute(w, nil)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

         email := r.FormValue("email")
		//  fmt.Printf("Email: %s\n", email)
		var keyEmail int
		emails := []string{"via@mail.com","okta@mail.com","ayu@mail.com","fitri@mail.com"}
		result := getDataPeserta(emails)
		checkExist:=false
		for key,value:=range emails{
			if email == value{
				keyEmail = key
				checkExist = true
			}
		}

		if checkExist==false {
			http.Error(w, "Email doesn't exist", http.StatusBadRequest)
			return
		}
		for key, value := range result{
			if keyEmail==key{
				var data = map[string]string{
					"email":email,
					"nama":value.Nama,
					"alamat":value.Alamat,
					"pekerjaan":value.Pekerjaan,
					"alasan":value.Alasan,
				}
				if err := tmpl.Execute(w, data); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func getDataPeserta(data []string)[]Biodata{
	
	datas :=[]Biodata{}
	for i:=0;i<len(data);i++{
		newData:=Biodata{Nama:"Via", Email:data[i],Alamat:"Jln lorem",Pekerjaan:"Backend",Alasan:"seru" }
		datas = append(datas, newData)
	}
	return datas
}

