package summary

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gaoey/covid-api/models"
)

func Test_UnmarshalResponse(t *testing.T) {
	respString := "{ \"Data\": [ { \"ConfirmDate\": \"2021-05-04\", \"No\": null, \"Age\": 51, \"Gender\": \"หญิง\", \"GenderEn\": \"Female\", \"Nation\": null, \"NationEn\": \"China\", \"Province\": \"Phrae\", \"ProvinceId\": 46, \"District\": null, \"ProvinceEn\": \"Phrae\", \"StatQuarantine\": 5 } ]}"
	var covidStatResponse models.CovidStatResponse
	if err := json.Unmarshal([]byte(respString), &covidStatResponse); err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("covidStatResponse: %v", covidStatResponse)
}
