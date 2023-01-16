package api

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
	"fmt"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	username := r.Context().Value("username").(string) // TODO: replace this
	r.ParseForm()

	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	// TODO: answer here

	if (len(r.Form) == 0){
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := model.ErrorResponse{
			Error: "Request Product Not Found",
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

	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{}
	cart.Name = username
	cart.Cart = list
	
	hasil := 0.0
	
	for i, data := range list{
		hasil = hasil + data.Total
		fmt.Sprintln(i)
	}
	
	cart.TotalPrice = hasil// TODO: replace this

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)

}
