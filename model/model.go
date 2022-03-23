package model

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Delivery struct {
	Name    **string `json:"name" validate:"required"`
	Phone   *string  `json:"phone" validate:"required"`
	Zip     *string  `json:"zip" validate:"required"`
	City    *string  `json:"city" validate:"required"`
	Address *string  `json:"address" validate:"required"`
	Region  *string  `json:"region" validate:"required"`
	Email   *string  `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  *string `json:"transaction" validate:"required"`
	RequestId    *string `json:"request_id" validate:"required"`
	Currency     *string `json:"currency" validate:"required"`
	Provider     *string `json:"provider" validate:"required"`
	Amount       *int    `json:"amount" validate:"required,gt=0"`
	PaymentDate  *int    `json:"payment_dt" validate:"required,gt=0"`
	Bank         *string `json:"bank" validate:"required"`
	DeliveryCost *int    `json:"delivery_cost" validate:"required,gte=0"`
	GoodsTotal   *int    `json:"goods_total" validate:"required,gt=0"`
	CustomFee    *int    `json:"custom_fee" validate:"required,gte=0"`
}

type Item struct {
	ChrtId      *int    `json:"chrt_id" validate:"required"`
	TrackNumber *string `json:"track_number" validate:"required"`
	Price       *int    `json:"price" validate:"required"`
	Rid         *string `json:"rid" validate:"required"`
	Name        *string `json:"name" validate:"required"`
	Sale        *int    `json:"sale" validate:"required"`
	Size        *string `json:"size" validate:"required"`
	TotalPrice  *int    `json:"total_price" validate:"required"`
	NmId        *int    `json:"nm_id" validate:"required"`
	Brand       *string `json:"brand" validate:"required"`
	Status      *int    `json:"status" validate:"required"`
}

type Order struct {
	Uid         *string `json:"order_uid" validate:"required"`
	TrackNumber *string `json:"track_number" validate:"required"`
	Entry       *string `json:"entry" validate:"required"`

	Delivery *Delivery `json:"delivery" validate:"required"`
	Payment  *Payment  `json:"payment" validate:"required"`
	Items    *[]Item   `json:"items" validate:"required"`

	Locale            *string `json:"locale" validate:"required"`
	InternalSignature *string `json:"internal_signature" validate:"required"`
	CustomerId        *string `json:"customer_id" validate:"required"`
	DeliveryService   *string `json:"delivery_service" validate:"required"`

	Shardkey    *string `json:"shardkey" validate:"required"`
	SmId        *int    `json:"sm_id" validate:"required"`
	DateCreated *string `json:"date_created" validate:"required"`
	OofShard    *string `json:"oof_shard" validate:"required"`
}

func (o *Order) FromJSONBuffer(data []byte) error {
	return json.Unmarshal(data, o)
}

func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
