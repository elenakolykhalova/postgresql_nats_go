package valid

import (
	"encoding/json"
)

func ValidateMyJSON(Data []byte)int{
	var myOrder Order
	err := json.Unmarshal(Data, &myOrder)
	if err != nil{
		return 0
	}
	if len(myOrder.Order_uid) < 10 || len(myOrder.Track_number) < 10 || 
	len(myOrder.Delivery.Name) < 1 || len(myOrder.Delivery.Phone) < 10 ||
	len(myOrder.Delivery.Zip) < 6 || len(myOrder.Delivery.City) < 3 || 
	len(myOrder.Delivery.Address) < 6 || len(myOrder.Delivery.Region) < 3 ||
	len(myOrder.Delivery.Email) < 3 || len(myOrder.Custumer_id) < 2 ||
	len(myOrder.Delivery_service) < 2 {
		return 0
	}
	if (len(myOrder.Payment.Transaction) < 10 || len(myOrder.Payment.Currency) < 1 ||
	len(myOrder.Payment.Provider) < 1 || myOrder.Payment.Amount < 0 || 
	myOrder.Payment.Payment_dt < 0 || len(myOrder.Payment.Bank) < 1 ||
	myOrder.Payment.Delivery_cost < 0 || myOrder.Payment.Goods_total < 0 || 
	myOrder.Payment.Custom_fee < 0) {
		return 0
	}
	for _, item := range(myOrder.Items){
		if item.Chrt_id < 9999 || len(item.Track_number) < 10 ||
		item.Price < 0 || len(item.Rid) < 10 ||
		len(item.Name) < 3 || (item.Sale < 0 || item.Sale > 100) ||
		item.Total_price < 0 || item.Nm_id < 9999 ||
		len(item.Brand) < 3 || item.Status < 0 || item.Track_number != myOrder.Track_number {
			return 0
		}
	}
	if myOrder.Order_uid != myOrder.Payment.Transaction{
		return 0
	}
	return 1
}