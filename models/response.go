package models

type CovidStatResponse struct {
	Data []CovidStat `json:"Data"`
}

type CovidStat struct {
	Age        int    `json:"Age"`
	Province   string `json:"Province"`
	ProvinceId int    `json:"ProvinceId"`
}

type Province map[string]int
type AgeGroup map[string]int
type CovidSummaryResponse struct {
	Province Province
	AgeGroup AgeGroup
}
