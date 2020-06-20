package charts

import (
	"bytes"
	tpls "github.com/go-echarts/go-echarts/templates"
	"html/template"
	"io"
	"regexp"
)

// 渲染图表
func renderChart(chart interface{}, w io.Writer, name string) error {
	contents := []string{tpls.HeaderTpl, tpls.RoutersTpl, tpls.BaseTpl}
	switch name {
	case "chart":
		contents = append(contents, tpls.ChartTpl)
	case "page":
		contents = append(contents, tpls.PageTpl)
	}
	tpl := template.Must(template.New("").Parse(contents[0]))
	mustTpl(tpl, contents[1:]...)
	return tpl.ExecuteTemplate(w, name, chart)
}

func mustTpl(tpl *template.Template, html ...string) {
	for i := 0; i < len(html); i++ {
		tpl = template.Must(tpl.Parse(html[i]))
	}
}

func renderToWriter(chart interface{}, renderName string, removeStr []string, w ...io.Writer) error {
	var b bytes.Buffer
	if err := renderChart(chart, &b, renderName); err != nil {
		return err
	}
	res := replaceRender(b, removeStr...)
	for i := 0; i < len(w); i++ {
		_, err := w[i].Write(res)
		if err != nil {
			return err
		}
	}
	return nil
}

// 过滤替换渲染结果
func replaceRender(b bytes.Buffer, notReplace ...string) []byte {
	// __x__ 与模板占位符相匹配
	idPat, _ := regexp.Compile(`(__x__")|("__x__)`)
	// 替换并转为 []byte 类型
	content := idPat.ReplaceAllString(b.String(), "")
	unusedObj := []string{
		`geo: {},?`,
		`"normal":{},?`,
		`"textStyle":{},?`,
		`"subtextStyle":{},?`,
		`"inRange":{},?`,
		`"label":{},?`,
		`"markLine":{},?`,
		`"markPoint":{},?`,
		`"itemStyle":{},?`,
		`"areaStyle":{},?`,
		`"lineStyle":{},?`,
		`"rippleEffect":{},?`,
		`"splitArea":{},?`,
		`"outline":{"show":false},?`,
		`"waveAnimation":false,?`,
		`"viewControl":{},?`,
		`"force":{},?`,
	}
	unusedObj = removeNotReplace(unusedObj, notReplace...)
	// 移除无用的 JSON object
	// 另一种解决方案是使用 *struct
	var unusedPat string
	for i := 0; i < len(unusedObj); i++ {
		unusedPat += unusedObj[i] + "|"
	}
	p, _ := regexp.Compile(unusedPat)
	res := p.ReplaceAllString(content, "")
	return []byte(res)
}

// 针对某些图表移除
func removeNotReplace(unusedObj []string, removeStr ...string) []string {
	res := make([]string, 0)
	for i := 0; i < len(unusedObj); i++ {
		isRemove := false
		for j := 0; j < len(removeStr); j++ {
			if unusedObj[i] == removeStr[j] {
				isRemove = true
				break
			}
		}
		if !isRemove {
			res = append(res, unusedObj[i])
		}
	}
	return res
}
