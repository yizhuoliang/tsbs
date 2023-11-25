package iot

import (
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/yizhuoliang/tsbs/pkg/query"
)

// AvgVsProjectedFuelConsumption contains info for filling in avg vs projected fuel consumption queries.
type AvgVsProjectedFuelConsumption struct {
	core utils.QueryGenerator
}

// NewAvgVsProjectedFuelConsumption creates a new avg vs projected fuel consumption query filler.
func NewAvgVsProjectedFuelConsumption(core utils.QueryGenerator) utils.QueryFiller {
	return &AvgVsProjectedFuelConsumption{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *AvgVsProjectedFuelConsumption) Fill(q query.Query) query.Query {
	fc, ok := i.core.(AvgVsProjectedFuelConsumptionFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.AvgVsProjectedFuelConsumption(q)
	return q
}
