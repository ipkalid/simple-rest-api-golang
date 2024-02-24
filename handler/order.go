package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ipkalid/order-api/model"
	order_repo "github.com/ipkalid/order-api/repository/order"
)

type Order struct {
	Repo *order_repo.RedisRepo
}

func NewOrder(orderRepo order_repo.RedisRepo) *Order {
	order := &Order{Repo: &orderRepo}

	return order
}

func (o Order) Create(w http.ResponseWriter, r *http.Request) {

	var body struct {
		ListItem []model.LineItem `json:"list_item"`
	}

	var error_message model.ErrorJson = model.ErrorJson{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		error_message.StatusCode = http.StatusBadRequest
		error_message.ErrorMessage = fmt.Sprintf("error decoding body %v \n", err)
		data, err := error_message.Json()

		if err != nil {
			return
		}

		w.Write(data)
		return
	}
	var order = model.Order{
		OrderId:    3,
		CustomerID: uuid.Must(uuid.NewUUID()),
		LineItems:  body.ListItem,
		CreatedAt:  time.Now().UTC(),
	}
	data, err := json.Marshal(order)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		error_message.StatusCode = http.StatusBadRequest
		error_message.ErrorMessage = fmt.Sprintf("error decoding body %v \n", err)
		data, err := error_message.Json()

		if err != nil {
			return
		}

		w.Write(data)
		return
	}

	err = o.Repo.Insert(r.Context(), order)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		error_message.StatusCode = http.StatusBadRequest
		error_message.ErrorMessage = fmt.Sprintf("error decoding body %v \n", err)
		data, err := error_message.Json()

		if err != nil {
			return
		}

		w.Write(data)
		return
	}

	w.Write(data)

}
func (o Order) List(w http.ResponseWriter, r *http.Request) {

}
func (o Order) GetByID(w http.ResponseWriter, r *http.Request) {

}
func (o Order) UpdateByID(w http.ResponseWriter, r *http.Request) {

}
func (o Order) DeleteByID(w http.ResponseWriter, r *http.Request) {

}
