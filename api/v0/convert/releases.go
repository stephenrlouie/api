package convert

import (
	"k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/timeconv"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// Takes GRPC release structure and converts it to opitkon API structure
func ReleaseToJSON(rel *release.Release) (*models.ReleaseRelease, error) {
	hks, err := HooksToJSON(rel.Hooks)
	if err != nil {
		return &models.ReleaseRelease{}, err
	}

	ifo, err := InfoToJSON(rel.Info)
	if err != nil {
		return &models.ReleaseRelease{}, err
	}

	return &models.ReleaseRelease{
		Chart:     ChartToJSON(rel.Chart),
		Config:    ConfigToJSON(rel.Config),
		Hooks:     hks,
		Info:      ifo,
		Manifest:  rel.Manifest,
		Name:      rel.Name,
		Namespace: rel.Namespace,
		Version:   rel.Version,
	}, nil
}

func InfoToJSON(info *release.Info) (*models.ReleaseInfo, error) {
	delTime := ""
	if info.GetDeleted() != nil {
		delTime = timeconv.String(info.GetDeleted())
	}

	firstDeployed := ""
	if info.GetFirstDeployed() != nil {
		firstDeployed = timeconv.String(info.GetFirstDeployed())
	}

	lastDeployed := ""
	if info.GetLastDeployed() != nil {
		lastDeployed = timeconv.String(info.GetLastDeployed())
	}

	return &models.ReleaseInfo{
		Deleted:       delTime,
		Description:   info.Description,
		FirstDeployed: firstDeployed,
		LastDeployed:  lastDeployed,
	}, nil
}

func HooksToJSON(hks []*release.Hook) (models.ReleaseReleaseHooks, error) {
	rets := models.ReleaseReleaseHooks{}

	for _, h := range hks {
		last := ""
		if h.GetLastRun() != nil {
			last = timeconv.String(h.GetLastRun())
		}
		rets = append(rets, &models.ReleaseHook{
			DeletePolicies: deletePoliciesToJSON(h.DeletePolicies),
			Events:         eventsToJSON(h.Events),
			Kind:           h.Kind,
			LastRun:        last,
			Manifest:       h.Manifest,
			Name:           h.Name,
			Path:           h.Path,
			Weight:         h.Weight,
		})
	}
	return rets, nil
}

func deletePoliciesToJSON(policies []release.Hook_DeletePolicy) []string {
	ret := []string{}
	for _, v := range policies {
		ret = append(ret, release.Hook_DeletePolicy_name[int32(v)])
	}
	return ret
}

func eventsToJSON(events []release.Hook_Event) []string {
	ret := []string{}
	for _, v := range events {
		ret = append(ret, release.Hook_Event_name[int32(v)])
	}
	return ret
}

func StatusToJSON(status *release.Status) (*models.ReleaseStatus, error) {
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

// TestRunsToJSON - this is a comment
func TestRunsToJSON(testRun []*release.TestRun) ([]*models.ReleaseTestRun, error) {
	ret := []*models.ReleaseTestRun{}

	for _, run := range testRun {
		startedAt := ""
		completedAt := ""
		//var err error
		if run.GetStartedAt() != nil {
			startedAt = timeconv.String(run.GetStartedAt())
		}

		if run.GetCompletedAt() != nil {
			completedAt = timeconv.String(run.GetCompletedAt())
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

func TestSuiteToJSON(testSuite *release.TestSuite) (*models.ReleaseTestSuite, error) {
	ret := models.ReleaseTestSuite{}
	completeTime := ""
	startedTime := ""
	var err error

	if testSuite.GetCompletedAt() != nil {
		completeTime = timeconv.String(testSuite.GetCompletedAt())
	}

	if testSuite.GetStartedAt() != nil {
		startedTime = timeconv.String(testSuite.GetStartedAt())
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

func IsTimeValid(time string) bool {
	if time == "<nil>" {
		return false
	}
	return true
}
