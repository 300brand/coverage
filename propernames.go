package main

import (
	"fmt"
	"sort"
)

type ProperName struct {
	Name   string
	Offset int
}

func (p ProperName) End() int       { return p.Offset + len(p.Name) }
func (p ProperName) String() string { return p.Name }

type ProperNames []*ProperName

func (p ProperNames) Len() int           { return len(p) }
func (p ProperNames) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ProperNames) Less(i, j int) bool { return p[i].Name < p[j].Name }

func (p ProperNames) Counts() (c ProperNameCounts) {
	for _, n := range p {
		i := 0
		found := false
		for i = range c {
			if found = c[i].Name == n.Name; found {
				c[i].Count++
				break
			}
		}
		if !found {
			c = append(c, &ProperNameCount{
				Name:  n.Name,
				Count: 1,
			})
		}
	}
	sort.Sort(c)
	return
}

type ProperNameCount struct {
	Name  string
	Count int
}

func (p ProperNameCount) String() string { return p.Name }

type ProperNameCounts []*ProperNameCount

func (p ProperNameCounts) Len() int           { return len(p) }
func (p ProperNameCounts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ProperNameCounts) Less(i, j int) bool { return p[i].Count > p[j].Count }
