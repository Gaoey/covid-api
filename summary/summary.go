package summary

import "github.com/Gaoey/covid-api/models"

const (
	TEEN  = "0-30"
	ADULT = "31-60"
	OLD   = "61+"
	OTHER = "N/A"
)

func GetSummary(stats []models.CovidStat) (*models.CovidSummaryResponse, error) {
	ageGroupCount := map[string]int{
		TEEN:  0,
		ADULT: 0,
		OLD:   0,
		OTHER: 0,
	}
	provinceCount := map[string]int{}

	for _, v := range stats {
		provinceCount = ProvinceCountFunc(provinceCount, v.Province)
		ageGroupCount = AgeGroupCountFunc(ageGroupCount, v.Age)
	}

	return &models.CovidSummaryResponse{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}, nil
}

func ProvinceCountFunc(provinceCount map[string]int, province string) map[string]int {
	// avoiding concurrency by pure function
	tempProvinceCount := provinceCount
	if province == "" {
		tempProvinceCount[OTHER] = tempProvinceCount[OTHER] + 1
	} else {
		tempProvinceCount[province] = tempProvinceCount[province] + 1
	}

	return tempProvinceCount
}

func AgeGroupCountFunc(ageGroup map[string]int, ptAge *int) map[string]int {
	// avoiding concurrency by pure function
	tempAgeGroup := ageGroup
	if ptAge == nil {
		tempAgeGroup[OTHER] = tempAgeGroup[OTHER] + 1
		return ageGroup
	}

	age := *ptAge
	if age >= 0 && age <= 30 {
		tempAgeGroup[TEEN] = tempAgeGroup[TEEN] + 1
	} else if age >= 31 && age <= 60 {
		tempAgeGroup[ADULT] = tempAgeGroup[ADULT] + 1
	} else {
		tempAgeGroup[OLD] = tempAgeGroup[OLD] + 1
	}

	return tempAgeGroup
}
