package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// HeatMap represents a heatmap chart.
type HeatMap struct {
	RectChart
}

func (HeatMap) Type() string { return types.ChartHeatMap }

// NewHeatMap creates a new heatmap chart.
func NewHeatMap() *HeatMap {
	c := &HeatMap{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.HasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *HeatMap) SetXAxis(x interface{}) *HeatMap {
	c.xAxisData = x
	return c
}

// AddSeries adds the new series.
func (c *HeatMap) AddSeries(name string, data []opts.HeatMapData, options ...SeriesOpts) *HeatMap {
	series := SingleSeries{Name: name, Type: types.ChartHeatMap, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *HeatMap) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
