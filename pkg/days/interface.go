package days

import (
	"github.com/wmuga/aoc2019/pkg/days/day1"
	"github.com/wmuga/aoc2019/pkg/days/day2"
	"github.com/wmuga/aoc2019/pkg/days/day3"
	"github.com/wmuga/aoc2019/pkg/days/day4"
	"github.com/wmuga/aoc2019/pkg/days/day5"
	"github.com/wmuga/aoc2019/pkg/days/day6"
	"github.com/wmuga/aoc2019/pkg/days/day7"
	"github.com/wmuga/aoc2019/pkg/days/day8"
	"github.com/wmuga/aoc2019/pkg/models"
)

var days = []models.Day{
	day1.Day{},
	day2.Day{},
	&day3.Day{},
	day4.Day{},
	day5.Day{},
	day6.Day{},
	day7.Day{},
	day8.Day{},
}

func GetDay(num int) (day models.Day, ok bool) {
	if num > len(days) {
		return nil, false
	}

	return days[num-1], true
}
