package agent

import (
	"bytes"
	"fmt"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/olekukonko/tablewriter"
	"sort"
	"strconv"
	"strings"
	"time"
)

func newQuantity(name string) *transfer.Quantity {
	var q transfer.Quantity
	q.ErrorMap = make(map[string]int64)
	q.Name = name
	return &q
}

type quantities struct {
	Value map[string]*transfer.Quantity `json:"value"`
	ErrResult map[string]int
}

func newQuantities() *quantities {
	return &quantities{
		Value: map[string]*transfer.Quantity{},
		ErrResult: map[string]int{},
	}
}

func computeAverage(val1, num1, val2, num2 int64) int64 {
	total := num1 + num2
	part1 := val1 / total * num1
	part2 := val2 / total * num2
	return part1 + part2
}

func printQuantities(quantities *quantities) {
	quantitySlice := make([]*transfer.Quantity, 0)
	for _, quantity := range quantities.Value {
		quantitySlice = append(quantitySlice, quantity)
	}
	if len(quantitySlice) > 1 {
		printQuantitySlice(quantitySlice)
	}
	if len(quantities.ErrResult) > 0 {
		printErrorMessage(quantities.ErrResult)
	}
}

func printQuantitySlice(quantities []*transfer.Quantity) {
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
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
	table.Render()
	title := "Plan Statistic"
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

func printErrorMessage(errs map[string]int)  {
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	header := []string{"Reason", "Count"}
	table.SetHeader(header)
	for err, count := range errs {
		row := make([]string, len(header))
		row[0] = err
		row[1] = strconv.Itoa(count)
		table.Append(row)
	}
	table.Render()
	title := "Task Error: "
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
			v, ok := quantities.ErrResult[outcome.Err]
			if !ok {
				quantities.ErrResult[outcome.Err] = 1
			}
			quantities.ErrResult[outcome.Err] = v + 1
		}
	}
}
