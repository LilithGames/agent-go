package agent

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/olekukonko/tablewriter"
)

func newQuantity(name string) *transfer.Quantity {
	var q transfer.Quantity
	q.ErrorMap = make(map[string]int64)
	q.Name = name
	return &q
}

type quantities struct {
	Name      string
	Value     map[string]*transfer.Quantity `json:"value"`
	ErrResult map[string]map[string]int
}

func newQuantities(name string) *quantities {
	return &quantities{
		Name:      name,
		Value:     map[string]*transfer.Quantity{},
		ErrResult: map[string]map[string]int{},
	}
}

func computeAverage(val1, num1, val2, num2 int64) int64 {
	total := num1 + num2
	part1 := val1 / total * num1
	part2 := val2 / total * num2
	return part1 + part2
}

func (q *quantities) print(opts ...*ViewOpt) {
	q.printQuantitySlice(opts...)
	q.printErrorMessage(opts...)
}

func (q *quantities) printQuantitySlice(opts ...*ViewOpt) {
	quantities := make([]*transfer.Quantity, 0)
	for _, quantity := range q.Value {
		quantities = append(quantities, quantity)
	}
	if len(quantities) <= 1 {
		return
	}
	opt := mergeViewOpt(opts...)
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	opt.apply(table)
	header := []string{"Name", "Total", "Fail", "Min", "Average", "Max"}
	table.SetHeader(header)
	names := make([]string, 0)
	var qm = make(map[string]*transfer.Quantity)
	for _, q := range quantities {
		if strings.Contains(q.Name, "polling_") {
			row := createQuantityRow(q)
			table.Append(row)
			continue
		}
		names = append(names, q.Name)
		qm[q.Name] = q
	}
	if table.NumLines() > 0 {
		table.Render()
		title := "Polling Statistic"
		fmt.Printf("\r\n%s\r\n%s", title, buf.String())
		table.ClearRows()
		buf.Reset()
	}
	sort.Strings(names)
	for _, name := range names {
		stat := qm[name]
		row := createQuantityRow(stat)
		table.Append(row)
	}
	table.SetRowLine(true)
	table.Render()
	title := fmt.Sprintf("Plan Statistic(%s): ", q.Name)
	fmt.Printf("\r\n%s\r\n%s", title, buf.String())
}

func createQuantityRow(stat *transfer.Quantity) []string {
	var row []string
	row = append(row, stat.Name)
	row = append(row, strconv.Itoa(int(stat.TotalNum)))
	row = append(row, strconv.Itoa(int(stat.ErrorNum)))
	row = append(row, toTimeStr(stat.MinTime))
	row = append(row, toTimeStr(stat.AvgTime))
	row = append(row, toTimeStr(stat.MaxTime))
	return row
}

func toTimeStr(val int64) string {
	dur := time.Duration(val)
	return dur.Round(time.Millisecond).String()
}

func (q *quantities) printErrorMessage(opts ...*ViewOpt) {
	if len(q.ErrResult) == 0 {
		return
	}
	opt := mergeViewOpt(opts...)
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	opt.apply(table)
	header := []string{"Name", "Reason", "Count"}
	table.SetHeader(header)
	for name, es := range q.ErrResult {
		for err, count := range es {
			row := make([]string, len(header))
			row[0] = name
			row[1] = err
			row[2] = strconv.Itoa(count)
			table.Append(row)
		}
	}
	table.SetRowLine(true)
	table.Render()
	title := fmt.Sprintf("Task Error(%s): ", q.Name)
	fmt.Printf("\r\n%s\r\n%s", title, buf.String())
}

func putReportData(quantities *quantities, outcomes []*transfer.Outcome) {
	for _, outcome := range outcomes {
		q, ok := quantities.Value[outcome.Name]
		if !ok {
			q = newQuantity(outcome.Name)
			quantities.Value[outcome.Name] = q
		}
		if q.MinTime == 0 || outcome.Consume < q.MinTime {
			q.MinTime = outcome.Consume
		}
		if outcome.Consume > q.MaxTime {
			q.MaxTime = outcome.Consume
		}
		q.AvgTime = computeAverage(q.AvgTime, q.TotalNum, outcome.Consume, 1)
		q.TotalNum++
		if outcome.Status == transfer.STATUS_FAILURE {
			q.ErrorNum++
		}
		if outcome.Err != "" {
			es, ok := quantities.ErrResult[outcome.Name]
			if !ok {
				es = make(map[string]int)
				quantities.ErrResult[outcome.Name] = es
			}
			es[outcome.Err]++
		}
	}
}
