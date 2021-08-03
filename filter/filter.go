//go:generate mockgen -destination mock_filter/filter.go github.com/andrewrobinson/news RunFilters

package filter

import (
	"fmt"
	"strings"
)

// type ClientInteractionClient interface {
// 	CreateServiceRequest(serviceRequest ServiceRequest) (*http.Response, *ServiceRequest, error)
// 	CreateInteraction(interaction AdhocInteraction) (*http.Response, *AdhocInteraction, error)
// 	CreateSurveyInteraction(survey SurveyInteraction) (*SurveyInteraction, error)
// 	CreateInteractionReason(reason InteractionReason) (*http.Response, error)
// 	UpdateSurvey(survey ClientInteractionSurvey) error
// }

func RunFilters(in int) int {
	return 5
}

func MultiTable(number int) string {
	//good luck

	var lines []string

	for i := 1; i <= 10; i++ {
		line := fmt.Sprintf("%v * %v = %v\n", i, number, i*number)
		lines = append(lines, line)
		//output = output + fmt.Sprintf("%v * %v = %v\n", i, number, i*number)
	}

	return strings.Join(lines, "\n")

}
