package goecharts

import (
	"io"
)

type Geo struct {
	BaseOpts
	Series
}

//工厂函数，生成 `Geo` 实例
func NewGeo(mapType string, routers ...HttpRouter) *Geo {
	geoChart := new(Geo)
	geoChart.HasXYAxis = false
	geoChart.initBaseOpts(routers...)
	geoChart.initAssetsOpts()
	geoChart.JSAssets = append(geoChart.JSAssets, "maps/"+mapType+".js")
	geoChart.GeoOpts.Map = mapType
	return geoChart
}

func (c *Geo) Add(name, geoType string, data map[string]float32, options ...interface{}) *Geo {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, c.extendValue(k, v)})
	}
	series := singleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: "geo"}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Geo) extendValue(region string, v float32) []float32 {
	res := make([]float32, 0)
	tv := Coordinates[region]
	res = append(tv[:], v)
	return res
}

func (c *Geo) SetGlobalConfig(options ...interface{}) *Geo {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Geo) validateOpts() {
	c.validateInitOpt()
	c.validateAssets(c.AssetsHost)
}

func (c *Geo) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	if err := renderToWriter(c, "chart", w...); err != nil {
		return err
	}
	return nil
}
