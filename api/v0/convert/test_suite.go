package convert

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

func TestSuiteToJSON(testSuite *release.TestSuite) (*models.ReleaseTestSuite, error) {
	fmt.Printf("TestSuiteToJSON\n")
	fmt.Printf("Received:  %+v\n", testSuite)
	ret := models.ReleaseTestSuite{}
	completeTime := strfmt.DateTime{}
	startedTime := strfmt.DateTime{}
	var err error

	if IsTimeValid(testSuite.GetCompletedAt().String()) {
		// completeTime, err = strfmt.ParseDateTime(testSuite.GetCompletedAt().String())
		// if err != nil {
		// 	return &ret, err
		// }
	}

	if IsTimeValid(testSuite.GetStartedAt().String()) {
		// startedTime, err = strfmt.ParseDateTime(testSuite.GetStartedAt().String())
		// if err != nil {
		// 	return &ret, err
		// }
	}

	results, err := TestRunsToJSON(testSuite.Results)
	if err != nil {
		return &ret, err
	}
	return &models.ReleaseTestSuite{
		CompletedAt: completeTime,
		StartedAt:   startedTime,
		Results:     results,
	}, nil
}
