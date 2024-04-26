package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	temp := `
{\"origin_id\":\"DKI00102\",\"destination_id\":\"DKI30\",\"commodity_id\":499,\"booked_for_type\":\"pos\",\"booked_for_id\":10120,\"goods_price\":50000,\"insurance_type\":\"free\",\"is_woodpacking\":false,\"is_have_tax_id\":false,\"pieces\":[{\"piece_length\":10,\"piece_width\":10,\"piece_height\":10,\"piece_gross_weight\":3}],\"city_rates\":22000,\"forward_rates\":3000,\"chargeable_weight\":3,\"shipping_cost\":25000,\"commodity_surcharge\":0,\"heavy_weight_surcharge\":0,\"document_surcharge\":0,\"insurance_rates\":0,\"woodpacking_rates\":0,\"total_tariff\":25000,\"tax_rates\":0,\"bm_tax_rate\":0,\"ppn_tax_rate\":0,\"pph_tax_rate\":0,\"origin_district_rate\":0,\"destination_district_rate\":3000,\"publish_rate\":19750,\"shipping_surcharge_rate\":2250,\"product_type\":\"REGPACK\",\"booked_by\":10120,\"booked_by_type\":\"pos\",\"stt_no\":\"10LP3514720229319\",\"shipment_id\":\"\",\"commodity_code\":\"GEN\",\"city_origin_id\":\"CGK\",\"city_destination_id\":\"CGK\",\"is_cod\":true,\"cod_fee\":1000,\"is_pad\":false,\"pad_fee\":0,\"is_stt_reverse_journey\":false,\"is_zero_credit_debit\":false,\"is_zero_cod_fee\":false,\"is_stt_shipment_favorite_reverse_journey\":false,\"partner_pos_parent_id\":0,\"partner_pos_branch_commission\":0,\"is_credit_debit_to_booked_by\":false,\"is_zero_credit_commission\":false,\"is_only_debit_transaction\":false,\"is_debit_to_booked_for\":false,\"credit_debit_actor_id\":0,\"credit_debit_actor_type\":\"\",\"is_hold_commission\":false,\"event_created_at\":\"0001-01-01T00:00:00Z\",\"is_discount_exceed_max_promo\":false,\"is_promo\":true,\"discount_type\":\"\",\"discount\":0,\"parameter_calculation\":\"\",\"total_discount\":7500,\"publish_rate_after_discount\":13825,\"shipping_surcharge_rate_after_discount\":1575,\"origin_district_rate_after_discount\":0,\"destination_district_rate_after_discount\":2100,\"document_surcharge_after_discount\":0,\"commodity_surcharge_after_discount\":0,\"heavy_weight_surcharge_after_discount\":0,\"woodpacking_rates_after_discount\":0,\"insurance_rates_after_discount\":0,\"cod_fee_after_discount\":1000,\"total_tariff_after_discount\":18500,\"total_surcharge_after_discount\":0,\"client_payment_method\":\"\",\"client_cod_config_amount\":\"\",\"client_cod_shipment_discount\":0,\"booking_return\":12500,\"cod_handling\":\"\",\"currency\":\"IDR\"}
`
	b, err := json.Marshal(temp)
	if err != nil {
		return
		fmt.Println("err", err)
	}

	lll := PgBookingRequest{}
	json.Unmarshal(b, &lll)

	fmt.Println("lorem", DumpToStringg(lll))
}

type Piece struct {
	PieceLength      float64 `json:"piece_length"`
	PieceWidth       float64 `json:"piece_width"`
	PieceHeight      float64 `json:"piece_height"`
	PieceGrossWeight float64 `json:"piece_gross_weight"`
}

type PgBookingRequest struct {
	// request from tariff estimation
	OriginID      string `json:"origin_id" form:"origin_id" query:"origin_id"`
	DestinationID string `json:"destination_id" form:"destination_id" query:"destination_id"`
	ProductType   string `json:"product_type" form:"product_type" query:"product_type"`
	// POS Or Internal Or Client
	BookedForID       int     `json:"booked_for_id" form:"booked_for_id" query:"booked_for_id"`
	BookedForIDBranch int     `json:"booked_for_id_branch"`
	BookedForType     string  `json:"booked_for_type" form:"booked_for_type" query:"booked_for_type"`
	CommodityID       int     `json:"commodity_id" form:"commodity_id" query:"commodity_id"`
	GoodsPrice        float64 `json:"goods_price" form:"goods_price" query:"goods_price"`
	InsuranceType     string  `json:"insurance_type" form:"insurance_type" query:"insurance_type"`
	IsWoodpacking     bool    `json:"is_woodpacking" form:"is_woodpacking" query:"is_woodpacking"`
	IsHaveTaxID       bool    `json:"is_have_tax_id" form:"is_have_tax_id" query:"is_have_tax_id"`
	Pieces            []Piece `json:"pieces"`

	// Result from tariff estimation
	CityRates               float64 `json:"city_rates"`
	ForwardRates            float64 `json:"forward_rates"`
	ChargeAbleWeight        float64 `json:"chargeable_weight"`
	ShippingCosts           float64 `json:"shipping_cost"`
	CommoditySurcharge      float64 `json:"commodity_surcharge"`
	HeavyWeightSurcharge    float64 `json:"heavy_weight_surcharge"`
	DocumentSurcharge       float64 `json:"document_surcharge"`
	InsuranceRates          float64 `json:"insurance_rates"`
	InsuranceName           string  `json:"insurance_name"`
	InsuranceLabel          string  `json:"insurance_label"`
	WoodpackingRates        float64 `json:"woodpacking_rates"`
	TotalTarifff            float64 `json:"total_tariff"`
	TaxRates                float64 `json:"tax_rates"`
	BMTaxRate               float64 `json:"bm_tax_rate"`
	PPNTaxRate              float64 `json:"ppn_tax_rate"`
	PPHTaxRate              float64 `json:"pph_tax_rate"`
	OriginDistrictRate      float64 `json:"origin_district_rate"`
	DestinationDistrictRate float64 `json:"destination_district_rate"`
	PublishRate             float64 `json:"publish_rate"`
	ShippingSurchargeRate   float64 `json:"shipping_surcharge_rate"`

	// POS Or Internal Or Client (Who Actor doing actual booking)
	BookingType  string `json:"booking_type"`
	BookedBy     int    `json:"booked_by"`
	BookedByType string `json:"booked_by_type"`

	SttNo      string `json:"stt_no"`
	ShipmentID string `json:"shipment_id"`
	Token      string `json:"token"`

	// New params
	CommodityCode     string `json:"commodity_code"`
	CityOriginID      string `json:"city_origin_id"`
	CityDestinationID string `json:"city_destination_id"`

	IsCod  bool    `json:"is_cod"`
	CodFee float64 `json:"cod_fee"`

	IsPad  bool    `json:"is_pad"`
	PadFee float64 `json:"pad_fee"`

	// Retry
	IsRetry     bool   `json:"is_retry"`
	AccountID   int64  `json:"account_id"`
	AccountName string `json:"account_name"`
	AccountType string `json:"account_type"`

	IsSttReverseJourney                 bool `json:"is_stt_reverse_journey"`
	IsZeroCreditDebit                   bool `json:"is_zero_credit_debit"`
	IsSttShipmentFavoriteReverseJourney bool `json:"is_stt_shipment_favorite_reverse_journey"`

	// POS Branch
	PartnerPosParentID         int     `json:"partner_pos_parent_id"`
	PartnerPosBranchCommission float64 `json:"partner_pos_branch_commission"`

	IsCreditDebitSttBookedBy bool `json:"is_credit_debit_to_booked_by"`
	IsZeroCreditCommission   bool `json:"is_zero_credit_commission"`

	IsDebitToBookedFor     bool   `json:"is_debit_to_booked_for"`
	CreditDebitActorID     int    `json:"credit_debit_actor_id"`
	CreditDebitActorType   string `json:"credit_debit_actor_type"`
	IsOnlyDebitTransaction bool   `json:"is_only_debit_transaction"`

	IsHoldCommission bool `json:"is_hold_commission"`

	IsZeroCreditDebitTransaction bool `json:"is_zero_credit_debit_transaction"`

	IsDiscountBookingValid bool
	TarifAfterDiscount

	BookingReturn float64 `json:"booking_return"`
	CODHandling   string  `json:"cod_handling"`

	Currency string `json:"currency"`
}

type TarifAfterDiscount struct {
	IsDiscountExceedMaxPromo             bool    `json:"is_discount_exceed_max_promo"`
	IsPromo                              bool    `json:"is_promo" bson:"is_promo"`
	DiscountType                         string  `json:"discount_type" bson:"discount_type"`
	Discount                             float64 `json:"discount" bson:"discount"`
	ParameterCalculations                string  `json:"parameter_calculation" bson:"parameter_calculation"`
	TotalDiscount                        float64 `json:"total_discount" bson:"total_discount"`
	PublishRateAfterDiscount             float64 `json:"publish_rate_after_discount" bson:"publish_rate_after_discount"`
	ShippingSurchargeRateAfterDiscount   float64 `json:"shipping_surcharge_rate_after_discount" bson:"shipping_surcharge_rate_after_discount"`
	OriginDistrictRateAfterDiscount      float64 `json:"origin_district_rate_after_discount" bson:"origin_district_rate_after_discount"`
	DestinationDistrictRateAfterDiscount float64 `json:"destination_district_rate_after_discount" bson:"destination_district_rate_after_discount"`
	DocumentSurchargeAfterDiscount       float64 `json:"document_surcharge_after_discount" bson:"document_surcharge_after_discount"`
	CommoditySurchargeAfterDiscount      float64 `json:"commodity_surcharge_after_discount" bson:"commodity_surcharge_after_discount"`
	HeavyWeightSurchargeAfterDiscount    float64 `json:"heavy_weight_surcharge_after_discount" bson:"heavy_weight_surcharge_after_discount"`
	WoodpackingRatesAfterDiscount        float64 `json:"woodpacking_rates_after_discount" bson:"woodpacking_rates_after_discount"`
	InsuranceRatesAfterDiscount          float64 `json:"insurance_rates_after_discount" bson:"insurance_rates_after_discount"`
	CodFeeAfterDiscount                  float64 `json:"cod_fee_after_discount" bson:"cod_fee_after_discount"`
	TotalTariffAfterDiscount             float64 `json:"total_tariff_after_discount" bson:"total_tariff_after_discount"`
}

func DumpToStringg(v interface{}) string {

	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		json.NewEncoder(buff).Encode(v)
		return buff.String()
	}

	return str
}
