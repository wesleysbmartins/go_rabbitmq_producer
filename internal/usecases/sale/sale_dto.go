package sale_usecase

type SaleMessageDTO struct {
	Sale struct {
		SellingCompany     string `json:"sellingCompany"`
		Product            string `json:"product"`
		Price              string `json:"price"`
		DeliveryCompany    string `json:"deliveryCompany"`
		OriginAddress      string `json:"originAddress"`
		DestinationAddress string `json:"destinationAddress"`
		ClientName         string `json:"clientName"`
		Order              int64  `json:"order"`
	} `json:"sale"`
	UserId string `json:"userId"`
	AppId  string `json:"appId"`
}
