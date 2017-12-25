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

	a := "€"
	if currency == "USD" { a = "$" }

	centsStr := fmt.Sprintf("%0.3d", abs(e.Change))

	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}

	if locale == "nl-NL" {
		a += " "
		//thousand sep
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + "."
		}

		a = a[:len(a)-1]
		a += ","
		a += centsStr[len(centsStr)-2:]
		
		if e.Change < 0 {
			a += "-"
		} else {
			a += " "
		}

	} else if locale == "en-US" {

		//thousand sep
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + ","
		}
		a = a[:len(a)-1]
		a += "."
		a += centsStr[len(centsStr)-2:]
		
		if e.Change < 0 {
			a = "(" + a + ")"
		} else {
			a += " "
		}
	}


	return a
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
