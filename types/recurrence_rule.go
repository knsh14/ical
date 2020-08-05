package types

type FrequencyPattern string

const (
	FrequencyPatternInvalid  FrequencyPattern = ""
	FrequencyPatternSecondly FrequencyPattern = "SECONDLY"
	FrequencyPatternMinutely FrequencyPattern = "MINUTELY"
	FrequencyPatternHourly   FrequencyPattern = "HOURLY"
	FrequencyPatternDaily    FrequencyPattern = "DAILY"
	FrequencyPatternWeekly   FrequencyPattern = "WEEKLY"
	FrequencyPatternMonthly  FrequencyPattern = "MONTHLY"
	FrequencyPatternYearly   FrequencyPattern = "YEARLY"
)

type WeekDayPattern string

const (
	WeekDayPatternInvalid   WeekDayPattern = ""
	WeekDayPatternSunday    WeekDayPattern = "SU"
	WeekDayPatternMonday    WeekDayPattern = "MO"
	WeekDayPatternTuesday   WeekDayPattern = "TU"
	WeekDayPatternWednesday WeekDayPattern = "WE"
	WeekDayPatternThursday  WeekDayPattern = "TH"
	WeekDayPatternFriday    WeekDayPattern = "FR"
	WeekDayPatternSaturday  WeekDayPattern = "SA"
)
