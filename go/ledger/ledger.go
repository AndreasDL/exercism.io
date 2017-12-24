package ledger

import (
	"errors"
	"strconv"
	"strings"
	"sort"
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

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	if len(entries) == 0 && currency != "USD"  {
		return "", errors.New("")
	} else if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("")
	}

	entriesSorted := copyAndSortEntries(entries)

	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesSorted {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			if d4 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
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
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
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
			} else {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
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
		}(i, et)
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
