package summary

import "github.com/Gaoey/covid-api/models"

var (
	teen  = "0-30"
	adult = "31-60"
	old   = "61+"
	other = "N/A"
)

func GetSummary(stats []models.CovidStat) (*models.CovidSummaryResponse, error) {
	ageGroup := map[string]int{
		teen:  0,
		adult: 0,
		old:   0,
		other: 0,
	}
	provinceCount := map[string]int{}

	for _, v := range stats {
		provinceCount = ProvinceCountFunc(provinceCount, v.Province)
		ageGroup = AgeGroupCountFunc(ageGroup, v.Age)
	}

	return &models.CovidSummaryResponse{
		Province: provinceCount,
		AgeGroup: ageGroup,
	}, nil
}

func ProvinceCountFunc(provinceCount map[string]int, province string) map[string]int {
	if province == "" {
		provinceCount["N/A"] = provinceCount["N/A"] + 1
	} else {
		provinceCount[province] = provinceCount[province] + 1
	}

	return provinceCount
}

func AgeGroupCountFunc(ageGroup map[string]int, ptAge *int) map[string]int {
	if ptAge == nil {
		ageGroup[other] = ageGroup[other] + 1
		return ageGroup
	}

	age := *ptAge
	if age >= 0 && age <= 30 {
		ageGroup[teen] = ageGroup[teen] + 1
	} else if age >= 31 && age <= 60 {
		ageGroup[adult] = ageGroup[adult] + 1
	} else {
		ageGroup[old] = ageGroup[old] + 1
	}

	return ageGroup
}
