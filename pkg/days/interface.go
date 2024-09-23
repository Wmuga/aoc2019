package days

import (
	"github.com/wmuga/aoc2019/pkg/days/day1"
	"github.com/wmuga/aoc2019/pkg/days/day2"
	"github.com/wmuga/aoc2019/pkg/days/day3"
	"github.com/wmuga/aoc2019/pkg/models"
)

var days = []models.Day{
	day1.Day{},
	day2.Day{},
	&day3.Day{},
}

func GetDay(num int) (day models.Day, ok bool) {
	if num > len(days) {
		return nil, false
	}

	return days[num-1], true
}
