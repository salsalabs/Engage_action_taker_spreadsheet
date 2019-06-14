package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

//Names for the sheets.
const (
	Actions      = "Actions"
	ActionTakers = "ActionTakers"
	Spreadsheet  = "action_analysis.xlsx"
)

//ActionKeys accepts the contents of an actions file and return a list of
//just the action keys
func ActionKeys(a [][]string) []string {
	var b []string
	for _, r := range a {
		b = append(b, r[0])
	}
	return b
}

//MapOffsets accepts the contents of an actions file.  It uses the action_KEY
//to create a lookup map of action_KEY to column offsets.
func MapOffsets(a [][]string) map[string]int {
	m := make(map[string]int)
	for i, r := range a {
		key := r[0]
		m[key] = i
	}
	return m
}

//Retrieve is a convenience mthod to accept  a tab-delimited file and return
//the contents.  Each line in the file is a single record. Fields in the record
//are delimited by tabs.
//
//Remember that the first line is the list of column heads...
func Retrieve(p string) (a [][]string, err error) {
	fmt.Printf("Retrieve: %v\n", p)
	f, err := os.Open(p)
	if err != nil {
		return a, err
	}
	r := csv.NewReader(f)
	a, err = r.ReadAll()
	return a, err
}

//StoreActions populates the actions sheet with the contents of a file.
func StoreActions(f *excelize.File, a [][]string) {
	for i, r := range a {
		for j, v := range r {
			loc, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(Actions, loc, v)

		}
	}
}

//StoreActionTakers populates the actions sheet with the contents of a file.
//Supporter information is inserted at the first of the line.  The last
//column in the supproter information is not inserted.  That is the action key.
//The action key is used to compute the column offset where the cell should go.
func StoreActionTakers(f *excelize.File, a [][]string, keys []string, offsets map[string]int) {
	for i, r := range a {
		for j, v := range r {
			loc, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(ActionTakers, loc, v)
		}
		width := len(r)
		if i == 0 {
			for j, v := range keys {
				if j == 0 {
					continue
				}
				//First action key overwrites the count at the end
				//of the action taker data.
				k := width - 1 + j
				loc, _ := excelize.CoordinatesToCellName(k+1, i+1)
				f.SetCellValue(ActionTakers, loc, v)
			}
		} else {
			key := r[width-2]
			count := r[width-1]
			d, ok := offsets[key]
			if ok {
				//First action key overwrites the count at the end
				//of the action taker data.
				j := width - 1 + d
				loc, _ := excelize.CoordinatesToCellName(j+1, i+1)
				f.SetCellValue(ActionTakers, loc, count)
			} else {
				fmt.Printf("Unable to find offset for action_KEY '%v'\n", key)
			}
		}
	}
}

//main is the application itself.
func main() {
	var (
		app             = kingpin.New("activity-analysis", "Create a spreadsheet of actions and action takers.")
		actionFile      = app.Flag("actions", "CSV file of action information").Required().String()
		actionTakerFile = app.Flag("action-takers", "CSV file of action taker information").Required().String()
	)
	app.Parse(os.Args[1:])
	actions, err := Retrieve(*actionFile)
	if err != nil {
		panic(err)
	}

	actionTakers, err := Retrieve(*actionTakerFile)
	if err != nil {
		panic(err)
	}

	f := excelize.NewFile()
	f.NewSheet(Actions)
	StoreActions(f, actions)
	a := ActionKeys(actions)
	m := MapOffsets(actions)
	f.NewSheet(ActionTakers)
	StoreActionTakers(f, actionTakers, a, m)

	// Save xlsx file by the given path.
	err = f.SaveAs(Spreadsheet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output is in %v\n", Spreadsheet)
}
