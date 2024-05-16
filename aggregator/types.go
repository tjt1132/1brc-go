package aggregator

import (
	"fmt"
	"math"
)

type StationStats struct {
	Name  string
	Min   float64
	Max   float64
	Total float64
	Count float64
}

func (s *StationStats) AddNewValue(value float64) {
	if value < s.Min {
		s.Min = value
	}
	if value > s.Max {
		s.Max = value
	}
	s.Total += value
	s.Count += 1
}

func (s *StationStats) Report() string {
	return fmt.Sprintf("%s=%.1f/%.1f/%.1f", s.Name, s.Min, math.Ceil((s.Total*10/s.Count))/10, s.Max)
}

type StationStatsMap = map[string]*StationStats
