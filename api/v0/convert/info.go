package convert

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

func InfoToJSON(info *release.Info) (*models.ReleaseInfo, error) {
	fmt.Printf("InfoToJSON\n")
	fmt.Printf("Received:  %+v\n", info)
	//var err error
	//ret := models.ReleaseInfo{}

	delTime := strfmt.DateTime{}
	if IsTimeValid(info.GetDeleted().String()) {
		// delTime, err = strfmt.ParseDateTime(info.GetDeleted().String())
		// if err != nil {
		// 	return &ret, err
		// }
	}

	firstDepTime := strfmt.DateTime{}
	if IsTimeValid(info.GetFirstDeployed().String()) {
		// firstDepTime, err = strfmt.ParseDateTime(info.GetFirstDeployed().String())
		// if err != nil {
		// 	return &ret, err
		// }
	}

	lastDepTime := strfmt.DateTime{}
	if IsTimeValid(info.GetLastDeployed().String()) {
		// lastDepTime, err = strfmt.ParseDateTime(info.GetLastDeployed().String())
		// if err != nil {
		// 	return &ret, err
		// }
	}

	return &models.ReleaseInfo{
		Deleted:       delTime,
		Description:   info.Description,
		FirstDeployed: firstDepTime,
		LastDeployed:  lastDepTime,
	}, nil
}
