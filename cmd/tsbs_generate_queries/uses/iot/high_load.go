package iot

import (
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/yizhuoliang/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/yizhuoliang/tsbs/pkg/query"
)

// TrucksWithHighLoad contains info for filling in trucks with high load queries.
type TrucksWithHighLoad struct {
	core utils.QueryGenerator
}

// NewTruckWithHighLoad creates a new trucks with high load query filler.
func NewTruckWithHighLoad(core utils.QueryGenerator) utils.QueryFiller {
	return &TrucksWithHighLoad{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *TrucksWithHighLoad) Fill(q query.Query) query.Query {
	fc, ok := i.core.(TruckHighLoadFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.TrucksWithHighLoad(q)
	return q
}
