package main

// Request/response structs for the API

type CreateShipmentReq struct {
	ItemType     string  `json:"item_type"`
	ItemDesc     string  `json:"item_desc"`
	ItemValue    float64 `json:"item_value"`
	Reason       string  `json:"reason"`
	SenderName   string  `json:"sender_name"`
	SenderEmail  string  `json:"sender_email"`
	SenderPhone  string  `json:"sender_phone"`
	RecvName     string  `json:"recv_name"`
	RecvAddress  string  `json:"recv_address"`
	RecvCity     string  `json:"recv_city"`
	RecvPostcode string  `json:"recv_postcode"`
	RecvCountry  string  `json:"recv_country"`
	Dest         string  `json:"dest"`
	Courier      string  `json:"courier"`
	PkgSize      string  `json:"pkg_size"`
	Total        float64 `json:"total"`
	Insurance    bool    `json:"insurance"`
}

type UpdateStatusReq struct {
	Status string `json:"status"`
	Note   string `json:"note"`
}

type VerifyPinReq struct {
	Pin string `json:"pin"`
}

type ChangePinReq struct {
	CurrentPin string `json:"current_pin"`
	NewPin     string `json:"new_pin"`
}

type UpdateSettingReq struct {
	Value string `json:"value"`
}
