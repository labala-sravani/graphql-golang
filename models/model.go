package models

type Vendor struct {
	PlatformVendorId string
	Description      string
	Error            ErrorResponse
}

type ErrorResponse struct {
	Message string
	Type    ErrorTypes
}

type TES struct {
	MinimumDeliverTime string
}

type DPS struct {
	MinimumOrderAmount string
}

type Characteristics struct{}
