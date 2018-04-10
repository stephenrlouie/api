package convert

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// TestRunsToJSON - this is a comment
func TestRunsToJSON(testRun []*release.TestRun) ([]*models.ReleaseTestRun, error) {
	fmt.Printf("TestRunToJSON\n")
	fmt.Printf("Received:  %+v\n", testRun)
	ret := []*models.ReleaseTestRun{}

	for _, run := range testRun {
		startedAt := strfmt.DateTime{}
		completedAt := strfmt.DateTime{}
		//var err error
		if IsTimeValid(run.GetStartedAt().String()) {
			// startedAt, err = strfmt.ParseDateTime(run.GetStartedAt().String())
			// if err != nil {
			// 	return ret, err
			// }
		}

		if IsTimeValid(run.GetCompletedAt().String()) {
			// completedAt, err = strfmt.ParseDateTime(run.GetCompletedAt().String())
			// if err != nil {
			// 	return ret, err
			// }
		}

		ret = append(ret, &models.ReleaseTestRun{
			Info:        run.Info,
			Name:        run.Name,
			StartedAt:   startedAt,
			CompletedAt: completedAt,
			Status:      testResultStatus(run.Status),
		})
	}
	return ret, nil
}

func testResultStatus(statusCode release.TestRun_Status) string {
	switch statusCode {
	case 0:
		return "UKNOWN"
	case 1:
		return "SUCCESS"
	case 2:
		return "FAILURE"
	case 3:
		return "RUNNING"
	}
	return "UNKNOWN"
}
