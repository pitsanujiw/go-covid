package entity

type CovidRecord struct {
	ConfirmDate    string
	No             int
	Age            int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

type CovidResponse struct {
	Data []CovidRecord
}

type CovidSummaryResponse struct {
	Province map[string]int `json:"province"`
	AgeGroup map[string]int `json:"ageGroup"`
}