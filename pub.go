package main

import (
	"bytes"
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"time"
)

func main() {
	sc, _ := stan.Connect("mess", "pub")
	response := Message{
		OrderUid:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: Deliveries{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: Payments{
			Transaction:  "b563feb7b2b84b6test",
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []Item{
			{
				ChrtId:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmId:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       time.Date(2021, time.December, 26, 6, 22, 19, 1, time.UTC),
		OofShard:          "1",
	}
	responseBytes := new(bytes.Buffer)

	err := json.NewEncoder(responseBytes).Encode(response)
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	for i := 1; ; i++ {
		sc.Publish("message", responseBytes.Bytes())
		time.Sleep(2 * time.Second)
	}

}
