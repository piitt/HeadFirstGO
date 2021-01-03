package hfdates

const DaysInWeek int = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysInWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
