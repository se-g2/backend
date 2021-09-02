package types

type DriverProfileResponse struct {
	IdentityCode  string        `json:"identityCode"`
	DriverLicense string        `json:"driverLicense"`
	Vehicles      []VehicleInfo `json:"vehicles"`
	AvailableTime string        `json:"availableTime"`
}
