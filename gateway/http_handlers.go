package main

import (
	"errors"
	"net/http"

	"github.com/balajiss36/common"
	"github.com/balajiss36/common/api"
	pb "github.com/balajiss36/common/api"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(api.OrderServiceClient) *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Path
	var items []*pb.ItemsWithQuantity

	err := common.ReadJSON(r, &items)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err = validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	errStatus := status.Convert(err)
	if errStatus != nil {
		common.WriteError(w, http.StatusBadRequest, errStatus.Err())
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return errors.New("items must not be empty")
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item ID must not be empty")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}
	return nil
}
