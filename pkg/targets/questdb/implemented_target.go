package questdb

import (
	"github.com/blagojts/viper"
	"github.com/spf13/pflag"
	"github.com/yizhuoliang/tsbs/pkg/data/serialize"
	"github.com/yizhuoliang/tsbs/pkg/data/source"
	"github.com/yizhuoliang/tsbs/pkg/targets"
	"github.com/yizhuoliang/tsbs/pkg/targets/constants"
)

func NewTarget() targets.ImplementedTarget {
	return &influxTarget{}
}

type influxTarget struct {
}

func (t *influxTarget) TargetSpecificFlags(flagPrefix string, flagSet *pflag.FlagSet) {
	flagSet.String(flagPrefix+"url", "http://localhost:9000/", "QuestDB REST end point")
	flagSet.String(flagPrefix+"ilp-bind-to", "127.0.0.1:9009", "QuestDB influx line protocol TCP ip:port")
}

func (t *influxTarget) TargetName() string {
	return constants.FormatQuestDB
}

func (t *influxTarget) Serializer() serialize.PointSerializer {
	return &Serializer{}
}

func (t *influxTarget) Benchmark(string, *source.DataSourceConfig, *viper.Viper) (targets.Benchmark, error) {
	panic("not implemented")
}
