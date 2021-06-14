package summary

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gaoey/covid-api/models"
)

func Test_Unmarshal_CovidStatResponse(t *testing.T) {
	respString := "{ \"Data\": [ { \"ConfirmDate\": \"2021-05-04\", \"No\": null, \"Age\": null, \"Gender\": \"หญิง\", \"GenderEn\": \"Female\", \"Nation\": null, \"NationEn\": \"China\", \"Province\": \"Phrae\", \"ProvinceId\": 46, \"District\": null, \"ProvinceEn\": \"Phrae\", \"StatQuarantine\": 5 } ]}"
	var covidStatResponse models.CovidStatResponse
	if err := json.Unmarshal([]byte(respString), &covidStatResponse); err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("covidStatResponse: %v", covidStatResponse)
}

type TestCase struct {
	given interface{}
	want  interface{}
}

func Test_Covid_Summary(t *testing.T) {
	// preparation
	fiftyOne := new(int)
	twenty := new(int)
	*fiftyOne = 51
	*twenty = 20
	// testcase
	testcases := []TestCase{
		{
			given: []models.CovidStat{
				{
					Age:      fiftyOne,
					Province: "Phrae",
				},
				{
					Age:      fiftyOne,
					Province: "Udon Thani",
				},
				{
					Age:      twenty,
					Province: "Udon Thani",
				},
				{
					Age:      nil,
					Province: "Phrae",
				},
			},
			want: models.CovidSummaryResponse{
				Province: map[string]int{
					"Phrae":      2,
					"Udon Thani": 2,
				},
				AgeGroup: map[string]int{
					"0-30":  1,
					"31-60": 2,
					"61+":   0,
					"N/A":   1,
				},
			},
		},
	}

	for _, v := range testcases {
		result, err := GetSummary(v.given.([]models.CovidStat))
		if err != nil {
			t.Errorf("get summary error: %#v\n", err)
		}
		if v.want != result {
			t.Errorf("given %#v want %#v but get %#v\n", v.given, v.want, result)
		}
	}
}
