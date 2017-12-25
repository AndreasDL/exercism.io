package ledger

import (
	"errors"
	"strings"
	"sort"
	"fmt"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}
func (e *Entry) formatDescription() string{
	if len(e.Description) > 25 {
		return e.Description[:22] + "..."
	} else {
		return e.Description + strings.Repeat(" ", 25-len(e.Description))
	}
}
func (e *Entry) formatDate(locale string) string{
	year, month, day := e.Date[0:4], e.Date[5:7], e.Date[8:10]
	if locale == "nl-NL" {
		return day + "-" + month + "-" + year
	} else if locale == "en-US" {
		return month + "/" + day + "/" + year
	}
	return ""
}
func (e *Entry) formatAmount(locale, currency string) string{
	centsStr := fmt.Sprintf("%0.3d", abs(e.Change))
	
	
	//split into groups of 3 to fix thousands separator
	parts := []string{}

	if leading := (len(centsStr) - 2) % 3 ; leading > 0 { 
		parts = append(parts, centsStr[:leading]) 
		centsStr = centsStr[leading:]
	}

	for b, e := 0, 3 ; e < len(centsStr) ; b,e = b+3, e+3 {
		parts = append(parts, centsStr[b:e])
	}

	result := "€"
	if currency == "USD" { result = "$" }

	if locale == "nl-NL" {
		result += " "
		result += strings.Join(parts, ".")
		result += ","
		result += centsStr[len(centsStr)-2:]
		
		if e.Change < 0 {
			result += "-"
		} else {
			result += " "
		}

	} else if locale == "en-US" {
		result += strings.Join(parts, ",")
		result += "."
		result += centsStr[len(centsStr)-2:]
		
		if e.Change < 0 {
			result = "(" + result + ")"
		} else {
			result += " "
		}
	}

	return result
}
func (e *Entry) format(locale, currency string) string{

	a := e.formatAmount(locale, currency)			
	al := 0 ; for range a { al++ }

	output := e.formatDate(locale)
	output += " | "
	output += e.formatDescription()
	output += " | " 
	output += strings.Repeat(" ", 13-al) 
	output += a 
	output += "\n"

	return output
}

func copyAndSortEntries( entries []Entry) []Entry {
	entriesSorted := make([]Entry, len(entries))
	for i, e := range entries { entriesSorted[i] = e }
	sort.Slice(entriesSorted, func(i,j int)bool{
		return entriesSorted[j].Change > entriesSorted[i].Change
	})
	return entriesSorted
}
func generateHeader(locale string) string {
	if locale == "nl-NL" {
		return "Datum      | Omschrijving              | Verandering\n"
	} else if locale == "en-US" {
		return "Date       | Description               | Change\n"
	}
	return ""
}
func abs(i int) int {
	if i < 0 { return -i }
	return i
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	//check flags
	if len(entries) == 0 && currency != "USD"  {
		return "", errors.New("")
	} else if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("")
	}

	//check entries
	for _, entry := range entries {
		if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
			return "", errors.New("")
		}
	}

	//format entries & save in order !
	output := make([]string, len(entries))
	for i, entry := range copyAndSortEntries(entries) {
		output[i]  = entry.format(locale, currency)
	}

	//generate string
	header := generateHeader(locale) + strings.Join(output, "")

	return header, nil
}