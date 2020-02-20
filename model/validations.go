package model

// IsValidMonth do validations input Month
// if it is in required range then will be return TRUE
func IsValidMonth(m int) bool {
	return m > 0 && m <= 12
}
// IsValidDay do validations input Day
// if it is in required range then will be return TRUE
func IsValidDay(d int) bool {
	return d > 0 && d <= 31
}
// IsValidWeekDay do validations input Week Day
// if it is in required range then will be return TRUE
func IsValidWeekDay(w int) bool {
	return w >= 0 && w <= 6
}
// IsValidHours do validations input Hours
// if they are in required range then will be return TRUE
func IsValidHours(h int) bool {
	return h >= 0 && h <= 23
}
// IsValidMinutes do validations input Minutes
// if they are in required range then will be return TRUE
func IsValidMinutes(m int) bool {
	return m >= 0 && m <= 60
}
// IsValidSeconds do validations input Seconds
// if they are in required range then will be return TRUE
func IsValidSeconds(s int) bool {
	return s >= 0 && s <= 60
}