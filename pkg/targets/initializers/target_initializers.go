package initializers

import (
	"fmt"
	"github.com/yizhuoliang/tsbs/pkg/targets"
	"github.com/yizhuoliang/tsbs/pkg/targets/akumuli"
	"github.com/yizhuoliang/tsbs/pkg/targets/cassandra"
	"github.com/yizhuoliang/tsbs/pkg/targets/clickhouse"
	"github.com/yizhuoliang/tsbs/pkg/targets/constants"
	"github.com/yizhuoliang/tsbs/pkg/targets/crate"
	"github.com/yizhuoliang/tsbs/pkg/targets/influx"
	"github.com/yizhuoliang/tsbs/pkg/targets/mongo"
	"github.com/yizhuoliang/tsbs/pkg/targets/prometheus"
	"github.com/yizhuoliang/tsbs/pkg/targets/questdb"
	"github.com/yizhuoliang/tsbs/pkg/targets/siridb"
	"github.com/yizhuoliang/tsbs/pkg/targets/timescaledb"
	"github.com/yizhuoliang/tsbs/pkg/targets/timestream"
	"github.com/yizhuoliang/tsbs/pkg/targets/victoriametrics"
	"strings"
)

func GetTarget(format string) targets.ImplementedTarget {
	switch format {
	case constants.FormatTimescaleDB:
		return timescaledb.NewTarget()
	case constants.FormatAkumuli:
		return akumuli.NewTarget()
	case constants.FormatCassandra:
		return cassandra.NewTarget()
	case constants.FormatClickhouse:
		return clickhouse.NewTarget()
	case constants.FormatCrateDB:
		return crate.NewTarget()
	case constants.FormatInflux:
		return influx.NewTarget()
	case constants.FormatMongo:
		return mongo.NewTarget()
	case constants.FormatPrometheus:
		return prometheus.NewTarget()
	case constants.FormatSiriDB:
		return siridb.NewTarget()
	case constants.FormatVictoriaMetrics:
		return victoriametrics.NewTarget()
	case constants.FormatTimestream:
		return timestream.NewTarget()
	case constants.FormatQuestDB:
		return questdb.NewTarget()
	}

	supportedFormatsStr := strings.Join(constants.SupportedFormats(), ",")
	panic(fmt.Sprintf("Unrecognized format %s, supported: %s", format, supportedFormatsStr))
}
