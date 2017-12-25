package ledger

import (
	"errors"
	"strconv"
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
	cents := e.Change
	negative := cents < 0
	if negative { cents = -cents }
	
	var a string
	if locale == "nl-NL" {
		
		if currency == "EUR" {
			a += "€"
		} else if currency == "USD" {
			a += "$"
		}
		
		a += " "
		
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}

		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + "."
		}
		a = a[:len(a)-1]
		a += ","
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += "-"
		} else {
			a += " "
		}
	} else if locale == "en-US" {
		if negative {
			a += "("
		}
		if currency == "EUR" {
			a += "€"
		} else if currency == "USD" {
			a += "$"
		}
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + ","
		}
		a = a[:len(a)-1]
		a += "."
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += ")"
		} else {
			a += " "
		}
	}

	return a
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	if len(entries) == 0 && currency != "USD"  {
		return "", errors.New("")
	} else if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("")
	}

	entriesSorted := copyAndSortEntries(entries)
	fmt.Print("")
	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, entry := range entriesSorted {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			d2, d4 := entry.Date[4], entry.Date[7]
			if d2 != '-' || d4 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			de := entry.formatDescription()
			d := entry.formatDate(locale)
			a := entry.formatAmount(locale, currency)
			
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, entry)
	}


	ss := make([]string, len(entriesSorted))
	for range entriesSorted {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	header := generateHeader(locale)
	for i := 0; i < len(entriesSorted); i++ {
		header += ss[i]
	}
	return header, nil
}
