package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"github.com/google/uuid"
	"time"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// Read username and password request with FormValue.
	creds := model.Credentials{} // TODO: replace this
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here

	// nilai := json.NewDecoder(r.Body)
	// Alert := nilai.Decode(&creds)

	// if(Alert != nil){
	// 	w.WriteHeader(400)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	Alert := model.ErrorResponse{
	// 		Error: "Username or Password empty",
	// 	}
	// 	dataAlert, _ := json.Marshal(Alert)
	// 	w.Write(dataAlert)
	// 	return
	// }

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := model.ErrorResponse{
			Error: "Username or Password empty",
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

	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	// Read usernmae and password request with FormValue.
	creds := model.Credentials{} // TODO: replace this
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here

	// nilai := json.NewDecoder(r.Body)
	// Alert := nilai.Decode(&creds)

	// if(Alert != nil){
	// 	w.WriteHeader(400)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	Alert := model.ErrorResponse{
	// 		Error: "Username or Password empty",
	// 	}
	// 	dataAlert, _ := json.Marshal(Alert)
	// 	w.Write(dataAlert)
	// 	return
	// }

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := model.ErrorResponse{
			Error: "Username or Password empty",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}else{
		
	}

	// w.WriteHeader(200)
	// w.WriteHeader(http.StatusOK)

	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
	// TODO: answer here

	cookie := uuid.NewString()
	expires := time.Now().Add(5 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Path: "/",
		Name: "session_token",
		Value: cookie,
		Expires: expires,
	})

	session := model.Session{
		Token: cookie,
		Username: creds.Username,
		Expiry: expires,
	} // TODO: replace this
	err = api.sessionsRepo.AddSessions(session)

	// body, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(body, &cred)
	// fmt.Printf("%#v\n", cred)
	// db.Users[cred.Username] = cred.Password
	// fmt.Println(db.Users)
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	//Read session_token and get Value:
	sessionToken := "" // TODO: replace this
	data := r.Context().Value("username").(string)
	dataterbaca, _ := api.sessionsRepo.ReadSessions()

	for i, _ := range dataterbaca{
		if (data == dataterbaca[i].Username){
			sessionToken = dataterbaca[i].Token
		}
	}
	api.sessionsRepo.DeleteSessions(sessionToken)

	//Set Cookie name session_token value to empty and set expires time to Now:
	// TODO: answer here

	http.SetCookie(w, &http.Cookie{
		Path: "/",
		Name: "session_token",
		Value: "",
		Expires: time.Now(),
	})

	// body, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(body, &cred)
	// fmt.Printf("%#v\n", cred)
	// db.Users[cred.Username] = cred.Password
	// fmt.Println(db.Users)
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
