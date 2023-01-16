package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	data, Alert := ioutil.ReadFile("./assets/images/img-avatar.png")
	
	if (Alert != nil){
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Alert := model.ErrorResponse{
			Error: "Server Internal Error",
		}
		DataAlert, _ := json.Marshal(Alert)
		w.Write(DataAlert)
		return
	}else{

	}

	// body, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(body, &cred)
	// fmt.Printf("%#v\n", cred)
	// db.Users[cred.Username] = cred.Password
	// fmt.Println(db.Users)
	w.Write(data)
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	// TODO: answer here
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	// TODO: answer here

	r.ParseMultipartForm(2 << 20)
	file, _, Alert := r.FormFile("file-avatar.png")
	hasil, Alert := os.OpenFile("./assets/images/img-avatar.png", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	hasilAkhir, Alert := io.Copy(hasil, file)
	fmt.Sprintln(hasilAkhir)

	if (Alert != nil){
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Alert := model.ErrorResponse{
			Error: "Server Internal Error",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}else{

	}

	// body, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(body, &cred)
	// fmt.Printf("%#v\n", cred)
	// db.Users[cred.Username] = cred.Password
	// fmt.Println(db.Users)
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	api.dashboardView(w, r)
}
