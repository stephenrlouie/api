package convert

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

func HooksToJSON(hks []*release.Hook) (models.ReleaseReleaseHooks, error) {

	fmt.Printf("HooksToJSON\n")
	fmt.Printf("Received:  %+v\n", hks)
	rets := models.ReleaseReleaseHooks{}

	for _, h := range hks {
		//var err error
		last := strfmt.DateTime{}
		if IsTimeValid(h.LastRun.String()) {
			// last, err = strfmt.ParseDateTime(h.LastRun.String())
			// if err != nil {
			// 	fmt.Printf("BUTTS\n")
			// 	return rets, err
			// }
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
