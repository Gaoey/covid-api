package models

type CovidStatResponse struct {
	Data []CovidStat `json:"Data"`
}
type CovidStat struct {
	Age      *int   `json:"Age"`
	Province string `json:"Province"`
}

type CovidSummaryResponse struct {
	Province map[string]int
	AgeGroup map[string]int
}
