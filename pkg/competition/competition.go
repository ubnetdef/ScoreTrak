package competition

import (
	"github.com/ubnetdef/ScoreTrak/pkg/check"
	"github.com/ubnetdef/ScoreTrak/pkg/config"
	"github.com/ubnetdef/ScoreTrak/pkg/host"
	"github.com/ubnetdef/ScoreTrak/pkg/hostgroup"
	"github.com/ubnetdef/ScoreTrak/pkg/policy"
	"github.com/ubnetdef/ScoreTrak/pkg/property"
	"github.com/ubnetdef/ScoreTrak/pkg/report"
	"github.com/ubnetdef/ScoreTrak/pkg/round"
	"github.com/ubnetdef/ScoreTrak/pkg/service"
	"github.com/ubnetdef/ScoreTrak/pkg/servicegroup"
	"github.com/ubnetdef/ScoreTrak/pkg/team"
	"github.com/ubnetdef/ScoreTrak/pkg/user"
)

// Competition is a struct that holds an aggregate of all models. This is used to upload/export competition as a file.
type Competition struct {
	Config        *config.DynamicConfig
	Report        *report.Report
	HostGroups    []*hostgroup.HostGroup
	Hosts         []*host.Host
	Teams         []*team.Team
	Services      []*service.Service
	ServiceGroups []*servicegroup.ServiceGroup
	Rounds        []*round.Round
	Properties    []*property.Property
	Checks        []*check.Check
	Users         []*user.User
	Policy        *policy.Policy
}
