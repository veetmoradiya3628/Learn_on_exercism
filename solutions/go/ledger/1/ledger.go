package ledger

import (
	"errors"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type result struct {
	i int
	s string
	e error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("")
	}

	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("")
	}
    
	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}

	// Sort by:
	// 1. Date
	// 2. Description
	// 3. Change
	//
	// Keep your custom sorting approach, but fix it.
	for i := 0; i < len(entriesCopy); i++ {
		for j := i + 1; j < len(entriesCopy); j++ {
			a := entriesCopy[i]
			b := entriesCopy[j]

			shouldSwap := false

			if b.Date < a.Date {
				shouldSwap = true
			} else if b.Date == a.Date {
				if b.Description < a.Description {
					shouldSwap = true
				} else if b.Description == a.Description && b.Change < a.Change {
					shouldSwap = true
				}
			}

			if shouldSwap {
				entriesCopy[i], entriesCopy[j] = entriesCopy[j], entriesCopy[i]
			}
		}
	}

	var s string

	if locale == "nl-NL" {
		s = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " +
			"Verandering" +
			strings.Repeat(" ", 13-len("Verandering")) +
			"\n"
	} else if locale == "en-US" {
		s = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " +
			"Change" +
			strings.Repeat(" ", 13-len("Change")) +
			"\n"
	} else {
		return "", errors.New("")
	}

	// Buffered channel.
	// This prevents a goroutine from blocking if it sends an error
	// and the main goroutine returns immediately.
	co := make(chan result, len(entriesCopy))

	for i, et := range entriesCopy {
		go func(i int, entry Entry) {

			// Validate date length.
			if len(entry.Date) != 10 {
				co <- result{e: errors.New("")}
				return
			}

			// Parse date.
			d1 := entry.Date[0:4]
			d2 := entry.Date[4]
			d3 := entry.Date[5:7]
			d4 := entry.Date[7]
			d5 := entry.Date[8:10]

			// Validate first '-'.
			if d2 != '-' {
				co <- result{e: errors.New("")}
				return
			}

			// Validate second '-'.
			if d4 != '-' {
				co <- result{e: errors.New("")}
				return
			}

			// Format description.
			de := entry.Description

			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}

			// Format date according to locale.
			var d string

			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			} else {
				co <- result{e: errors.New("")}
				return
			}

			// Handle negative values.
			negative := false
			cents := entry.Change

			if cents < 0 {
				cents = cents * -1
				negative = true
			}

			var a string

			// -------------------------
			// Dutch locale
			// -------------------------
			if locale == "nl-NL" {

				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- result{e: errors.New("")}
					return
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

				if negative {
					a += "-"
				}

				for j := len(parts) - 1; j >= 0; j-- {
					a += parts[j] + "."
				}

				// Remove trailing '.'
				a = a[:len(a)-1]

				a += ","
				a += centsStr[len(centsStr)-2:]
				a += " "

			// -------------------------
			// English locale
			// -------------------------
			} else if locale == "en-US" {

				if negative {
					a += "("
				}

				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- result{e: errors.New("")}
					return
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

				for j := len(parts) - 1; j >= 0; j-- {
					a += parts[j] + ","
				}

				// Remove trailing ','
				a = a[:len(a)-1]

				a += "."
				a += centsStr[len(centsStr)-2:]

				if negative {
					a += ")"
				} else {
					a += " "
				}

			} else {
				co <- result{e: errors.New("")}
				return
			}

			// Calculate amount length.
			var al int

			for range a {
				al++
			}

			co <- result{
				i: i,
				s: d +
					strings.Repeat(" ", 10-len(d)) +
					" | " +
					de +
					" | " +
					strings.Repeat(" ", 13-al) +
					a +
					"\n",
			}

		}(i, et)
	}

	// Collect results in original sorted order.
	ss := make([]string, len(entriesCopy))

	for range entriesCopy {
		v := <-co

		if v.e != nil {
			return "", v.e
		}

		ss[v.i] = v.s
	}

	for i := range len(entriesCopy) {
		s += ss[i]
	}

	return s, nil
}