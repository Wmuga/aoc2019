package days

import "github.com/wmuga/aoc2019/pkg/models"

var days = []models.Day{}

func GetDay(num int) (day models.Day, ok bool) {
	if num >= len(days) {
		return nil, false
	}
	return days[num-1], true
}
