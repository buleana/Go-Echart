package charts

type Scatter struct {
	RectChart
}

func NewScatter(routers ...HTTPRouter) *Scatter {
	chart := new(Scatter)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

// 提供 X 轴数据
func (c *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	c.xAxisData = xAxis
	return c
}

// 提供 Y 轴数据及 Series 配置项
func (c *Scatter) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Scatter {
	series := singleSeries{Name: name, Type: "scatter", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}
