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

var quantities *transfer.Quantities

func init() {
	quantities = newQuantities()
}

func newQuantity(name string, class transfer.CLASS) *transfer.Quantity {
	var q transfer.Quantity
	q.ErrorMap = make(map[string]int64)
	q.Name = name
	q.Class = class
	return &q
}

func newQuantities() *transfer.Quantities {
	return &transfer.Quantities{
		Handler: map[string]*transfer.Quantity{},
		Event:   map[string]*transfer.Quantity{},
	}
}

func printQuantitySlice(name string, qs map[string]*transfer.Quantity, opts ...*ViewOpt) {
	if len(qs) == 0 {
		return
	}
	opt := mergeViewOpt(opts...)
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	opt.apply(table)
	header := []string{"Name", "Total", "Fail", " 0~50MS", "50~100MS", "100~200MS", "200~500MS", "500~1S", "1~2S", "2~5S", "5~10S"}
	table.SetHeader(header)
	names := make([]string, 0)
	var qm = make(map[string]*transfer.Quantity)
	for _, q := range qs {
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
		title := "Polling Stat"
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
	title := fmt.Sprintf("Plan Stat (%s)", name)
	fmt.Printf("\r\n%s\r\n%s", title, buf.String())
}

func createQuantityRow(q *transfer.Quantity) []string {
	var row []string
	row = append(row, q.Name)
	row = append(row, strconv.Itoa(int(q.TotalNum)))
	row = append(row, strconv.Itoa(int(q.ErrorNum)))
	row = append(row, strconv.Itoa(int(q.Le50Ms)))
	row = append(row, strconv.Itoa(int(q.Le100Ms)))
	row = append(row, strconv.Itoa(int(q.Le200Ms)))
	row = append(row, strconv.Itoa(int(q.Le500Ms)))
	row = append(row, strconv.Itoa(int(q.Le1S)))
	row = append(row, strconv.Itoa(int(q.Le2S)))
	row = append(row, strconv.Itoa(int(q.Le5S)))
	row = append(row, strconv.Itoa(int(q.Le10S)))
	return row
}

func printErrorMessage(name string, qs map[string]*transfer.Quantity, opts ...*ViewOpt) {
	if len(qs) == 0 {
		return
	}
	opt := mergeViewOpt(opts...)
	buf := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	opt.apply(table)
	header := []string{"Name", "Reason", "Count"}
	table.SetHeader(header)
	for _, q := range qs {
		es := q.ErrorMap
		for err, count := range es {
			row := make([]string, len(header))
			row[0] = name
			row[1] = err
			row[2] = strconv.Itoa(int(count))
			table.Append(row)
		}
	}
	if table.NumLines() == 0 {
		return
	}
	table.SetRowLine(true)
	table.Render()
	title := fmt.Sprintf("Task Error (%s): ", name)
	fmt.Printf("\r\n%s\r\n%s", title, buf.String())
}

func pushData(outcomes []*transfer.Outcome) {
	for _, outcome := range outcomes {
		var value map[string]*transfer.Quantity
		switch outcome.Class {
		case transfer.CLASS_HANDLER:
			value = quantities.Handler
		default:
			value = quantities.Event
		}
		q, ok := value[outcome.Name]
		if !ok {
			q = newQuantity(outcome.Name, outcome.Class)
			value[outcome.Name] = q
		}
		ms := int64(time.Millisecond)
		s := int64(time.Second)
		switch {
		case outcome.Consume <= ms*50:
			q.Le50Ms++
		case outcome.Consume <= ms*100 && outcome.Consume > ms*50:
			q.Le100Ms++
		case outcome.Consume <= ms*200 && outcome.Consume > ms*100:
			q.Le200Ms++
		case outcome.Consume <= ms*500 && outcome.Consume > ms*200:
			q.Le500Ms++
		case outcome.Consume <= s*1 && outcome.Consume > ms*500:
			q.Le1S++
		case outcome.Consume <= s*2 && outcome.Consume > s*1:
			q.Le2S++
		case outcome.Consume <= s*5 && outcome.Consume > s*2:
			q.Le5S++
		case outcome.Consume <= s*10 && outcome.Consume > s*5:
			q.Le10S++
		}
		q.TotalNum++
		if outcome.Status == transfer.STATUS_FAILURE {
			q.ErrorNum++
		}
		if outcome.Err != "" {
			q.ErrorMap[outcome.Err]++
		}
	}
}

func pushLocalData(planName string, report *transfer.Report) error {
	pushData(report.Outcomes)
	return nil
}

func echoLocalData(planName string, view *ViewOpt) {
	printQuantitySlice(planName+":H", quantities.Handler, view)
	printQuantitySlice(planName+":E", quantities.Event, view)
	printErrorMessage(planName+":H", quantities.Handler, view)
	printErrorMessage(planName+":E", quantities.Handler, view)
	quantities = newQuantities()
}
