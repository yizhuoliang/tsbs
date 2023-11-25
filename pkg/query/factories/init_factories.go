package factories

import (
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/akumuli"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/cassandra"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/clickhouse"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/cratedb"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/influx"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/mongo"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/questdb"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/siridb"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/timescaledb"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/timestream"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/databases/victoriametrics"
	"github.com/yizhuoliang/tsbs/pkg/query/config"
	"github.com/yizhuoliang/tsbs/pkg/targets/constants"
)

func InitQueryFactories(config *config.QueryGeneratorConfig) map[string]interface{} {
	factories := make(map[string]interface{})
	factories[constants.FormatCassandra] = &cassandra.BaseGenerator{}
	factories[constants.FormatClickhouse] = &clickhouse.BaseGenerator{
		UseTags: config.ClickhouseUseTags,
	}
	factories[constants.FormatCrateDB] = &cratedb.BaseGenerator{}
	factories[constants.FormatInflux] = &influx.BaseGenerator{}
	factories[constants.FormatTimescaleDB] = &timescaledb.BaseGenerator{
		UseJSON:       config.TimescaleUseJSON,
		UseTags:       config.TimescaleUseTags,
		UseTimeBucket: config.TimescaleUseTimeBucket,
	}
	factories[constants.FormatSiriDB] = &siridb.BaseGenerator{}
	factories[constants.FormatMongo] = &mongo.BaseGenerator{
		UseNaive: config.MongoUseNaive,
	}
	factories[constants.FormatAkumuli] = &akumuli.BaseGenerator{}
	factories[constants.FormatVictoriaMetrics] = &victoriametrics.BaseGenerator{}
	factories[constants.FormatTimestream] = &timestream.BaseGenerator{
		DBName: config.DbName,
	}
	factories[constants.FormatQuestDB] = &questdb.BaseGenerator{}
	return factories
}
