package models

type Status struct {
	Water uint `json:"water"`
	Wind  uint `json:"wind"`
}

type Data struct {
	Status Status `json:"status"`
}
