package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	a := readLines("ff.txt")
	b := readLines("dd.txt")
	d := diff(a, b)
	fmt.Println(len(d))
	fmt.Println(strings.Join(d, ","))
}

func printformula() {
	tmpl := "=LOOKUP(A%d,工作表4!A2:A1819,工作表4!I2:I1819)"
	for i := 2; i <= 1718; i++ {
		fmt.Printf(tmpl, i)
		fmt.Println()
	}
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	seen := map[string]bool{}
	var ss []string
	for scanner.Scan() {
		t := scanner.Text()
		if t == "company_id" {
			continue
		}
		t = strings.TrimPrefix(t, "c")
		if seen[t] {
			continue
		}
		ss = append(ss, t)
		seen[t] = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ss
}

func diff(a, b []string) []string {
	dict := map[string]bool{}
	for _, c := range a {
		dict[c] = false
	}
	for _, c := range b {
		if _, found := dict[c]; found {
			dict[c] = true
		}
	}
	var ss []string
	for k, v := range dict {
		if v == false {
			ss = append(ss, k)
		}
	}
	return ss
}

func read() {
	f, err := os.Open("ff.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	seen := map[string]bool{}
	var ss []string
	for scanner.Scan() {
		t := scanner.Text()
		if t == "company_id" {
			continue
		}
		t = strings.TrimPrefix(t, "c")
		if seen[t] {
			continue
		}
		ss = append(ss, t)
		seen[t] = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(len(ss))
	fmt.Println(strings.Join(ss, ","))
}

func printUniq() {
	f, err := os.Open("ff.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	seen := map[string]bool{}
	var ss []string
	for scanner.Scan() {
		t := scanner.Text()
		if t == "company_id" {
			continue
		}
		t = strings.TrimPrefix(t, "c")
		if seen[t] {
			continue
		}
		ss = append(ss, t)
		seen[t] = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(len(ss))
	fmt.Println(strings.Join(ss, ","))
}
