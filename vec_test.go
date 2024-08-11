package goveccompute_test

import (
	"fmt"
	"testing"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	"github.com/apache/arrow/go/v18/arrow/memory"
	goveccompute "github.com/ringsaturn/go-vec-compute"
)

type TestCaseTable struct {
	Table  arrow.Table
	Record arrow.Record
}

func NewTestCaseTable(pool memory.Allocator, l int) *TestCaseTable {
	// 定义表结构
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "temperature", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
		{Name: "humidity", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
		{Name: "pressure", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
		{Name: "windspeed", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
		{Name: "apparent_temperature", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
	}, nil)

	// 创建记录生成器
	builder := array.NewRecordBuilder(pool, schema)
	defer builder.Release()

	temps := goveccompute.RandsInRange(l, 20, 40)
	humidities := goveccompute.RandsInRange(l, 50, 90)
	pressures := goveccompute.RandsInRange(l, 1000, 1020)
	windspeeds := goveccompute.RandsInRange(l, 1, 10)
	apparentTemps := make([]float64, l) // require to be filled

	builder.Field(0).(*array.Float64Builder).AppendValues(temps, nil)
	builder.Field(1).(*array.Float64Builder).AppendValues(humidities, nil)
	builder.Field(2).(*array.Float64Builder).AppendValues(pressures, nil)
	builder.Field(3).(*array.Float64Builder).AppendValues(windspeeds, nil)
	builder.Field(4).(*array.Float64Builder).AppendValues(apparentTemps, nil)

	record := builder.NewRecord()
	table := array.NewTableFromRecords(schema, []arrow.Record{record})
	return &TestCaseTable{
		Table:  table,
		Record: record,
	}
}

func (t *TestCaseTable) Release() {
	t.Table.Release()
}

var pool = memory.NewGoAllocator()

func benchmarkApparentTemperature(b *testing.B, l int) {
	// 创建内存分配器
	tbl := NewTestCaseTable(pool, l)
	defer tbl.Release()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		func() {
			apparentTempBuilder := array.NewFloat64Builder(pool)
			defer apparentTempBuilder.Release()
			for i := 0; i < int(tbl.Record.NumRows()); i++ {
				temp := tbl.Record.Column(0).(*array.Float64).Value(i)
				humidity := tbl.Record.Column(1).(*array.Float64).Value(i)
				pressure := tbl.Record.Column(2).(*array.Float64).Value(i)
				windspeed := tbl.Record.Column(3).(*array.Float64).Value(i)

				// 简化的体感温度公式（实际公式可能更加复杂）
				apparentTemp := goveccompute.ComputeApparentTemperature(temp, pressure, windspeed, humidity)
				apparentTempBuilder.Append(apparentTemp)
			}

			apparentTempArray := apparentTempBuilder.NewArray()
			defer apparentTempArray.Release()
		}()
	}
}

func BenchmarkApparentTemperature_Arrow(b *testing.B) {
	for _, hours := range []int{10, 120, 384, 480} {
		b.Run(fmt.Sprintf("%v hours", hours), func(b *testing.B) {
			benchmarkApparentTemperature(b, hours)
		})
	}
}
