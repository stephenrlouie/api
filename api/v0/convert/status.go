package convert

import (
	"fmt"

	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

func StatusToJSON(status *release.Status) (*models.ReleaseStatus, error) {
	fmt.Printf("StatusToJSON\n")
	fmt.Printf("Received:  %+v\n", status)
	suite, err := TestSuiteToJSON(status.LastTestSuiteRun)
	if err != nil {
		return &models.ReleaseStatus{}, err
	}
	return &models.ReleaseStatus{
		Code:             release.Status_Code_name[int32(status.Code)],
		LastTestSuiteRun: suite,
		Notes:            status.Notes,
		Resources:        status.Resources,
	}, nil
}
