package storagefx

import (
	"github.com/ubnetdef/ScoreTrak/pkg/check/checkservice"
	"github.com/ubnetdef/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ubnetdef/ScoreTrak/pkg/config/configservice"
	"github.com/ubnetdef/ScoreTrak/pkg/host/hostservice"
	"github.com/ubnetdef/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/ubnetdef/ScoreTrak/pkg/policy/policyservice"
	"github.com/ubnetdef/ScoreTrak/pkg/property/propertyservice"
	"github.com/ubnetdef/ScoreTrak/pkg/report/reportservice"
	"github.com/ubnetdef/ScoreTrak/pkg/round/roundservice"
	"github.com/ubnetdef/ScoreTrak/pkg/service/serviceservice"
	"github.com/ubnetdef/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/ubnetdef/ScoreTrak/pkg/storage"
	"github.com/ubnetdef/ScoreTrak/pkg/storage/orm"
	"github.com/ubnetdef/ScoreTrak/pkg/storage/util"
	"github.com/ubnetdef/ScoreTrak/pkg/team/teamservice"
	"github.com/ubnetdef/ScoreTrak/pkg/user/userservice"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewDB,
		orm.NewCheckRepo,
		orm.NewConfigRepo,
		orm.NewHostGroupRepo,
		orm.NewHostRepo,
		orm.NewReportRepo,
		orm.NewRoundRepo,
		orm.NewUserRepo,
		orm.NewPolicyRepo,
		orm.NewTeamRepo,
		orm.NewServiceRepo,
		orm.NewServiceGroupRepo,
		orm.NewPropertyRepo,
		util.NewRepoStore,
	),
	// Add Service Pattern Implementations
	fx.Provide(
		checkservice.NewCheckServ,
		hostgroupservice.NewHostGroupServ,
		hostservice.NewHostServ,
		propertyservice.NewPropertyServ,
		configservice.NewConfigServ,
		configservice.NewStaticConfigServ,
		competitionservice.NewCompetitionServ,
		policyservice.NewPolicyServ,
		reportservice.NewReportServ,
		userservice.NewUserServ,
		teamservice.NewTeamServ,
		servicegroupservice.NewServiceGroupServ,
		roundservice.NewRoundServ,
		serviceservice.NewServiceServ,
	),
	fx.Invoke(
		util.CheckDBTimeSync,
		util.AutoMigrate,
		util.LoadConfig,
		util.LoadReport,
		util.LoadPolicy,
	),
)
