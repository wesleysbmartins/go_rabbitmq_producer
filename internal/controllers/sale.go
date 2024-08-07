package controllers

import (
	"encoding/json"
	"fmt"
	sale_usecase "go_rabbitmq_producer/internal/usecases/sale"
	"net/http"
)

func SaleController(w http.ResponseWriter, r *http.Request) {
	dto := sale_usecase.SaleMessageDTO{}
	json.NewDecoder(r.Body).Decode(&dto)

	w.Header().Set("Content-Type", "application/json")

	err := validateFields(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		usecase := sale_usecase.SaleUsecase{}
		err := usecase.Produce(dto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Failed To Proccess Message!\nERROR: %s", err)
			json.NewEncoder(w).Encode(msg)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Sale Received With Success!")
		}
	}
}

func validateFields(dto sale_usecase.SaleMessageDTO) error {
	if dto.AppId == "" {
		return fmt.Errorf("Field appId is required!")
	} else if dto.UserId == "" {
		return fmt.Errorf("Field userId is required!")
	} else if dto.Sale.SellingCompany == "" {
		return fmt.Errorf("Field sellingCompany is required!")
	} else if dto.Sale.Product == "" {
		return fmt.Errorf("Field product is required!")
	} else if dto.Sale.Price == "" {
		return fmt.Errorf("Field price is required!")
	} else if dto.Sale.DeliveryCompany == "" {
		return fmt.Errorf("Field deliveryCompany is required!")
	} else if dto.Sale.OriginAddress == "" {
		return fmt.Errorf("Field originAddress is required!")
	} else if dto.Sale.DestinationAddress == "" {
		return fmt.Errorf("Field destinationAddress is required!")
	} else if dto.Sale.ClientName == "" {
		return fmt.Errorf("Field clientName is required!")
	} else if dto.Sale.Order == 0 {
		return fmt.Errorf("Field order is required!")
	}

	return nil
}
