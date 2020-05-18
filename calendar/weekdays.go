package calendar

import (
	"strings"

	// standard libs only above!

	"github.com/litesoft-go/utils/enums"
	"github.com/litesoft-go/utils/strs"
)

type Weekday struct {
	enums.Enum
	weekdayData
}

func (wdd *weekdayData) GetDayNumber() int {
	if wdd == nil {
		return 0
	}
	return wdd.dayNumber
}

func (wdd *weekdayData) IsWeekend() bool {
	dayNumber := wdd.GetDayNumber()
	return (dayNumber == 6) || (dayNumber == 7)
}

func (wdd *weekdayData) GetAbbreviation3() string {
	return wdd.getAbbr(func() string { return wdd.abbreviation3 })
}

func (wdd *weekdayData) GetAbbreviation2() string {
	return wdd.getAbbr(func() string { return wdd.abbreviation2 })
}

func (wdd *weekdayData) GetAbbreviation1() string {
	return wdd.getAbbr(func() string { return wdd.abbreviation1 })
}

var (
	Monday = Weekday{Enum: enums.New("Monday"), weekdayData: wddBuilder(1).abbrs("Mon", "Mo", "M").build()}

	Tuesday = Weekday{Enum: enums.New("Tuesday"), weekdayData: wddBuilder(2).abbrs("Tue", "Tu", "T").build()}

	Wednesday = Weekday{Enum: enums.New("Wednesday"), weekdayData: wddBuilder(3).abbrs("Wed", "We", "W").build()}

	Thursday = Weekday{Enum: enums.New("Thursday"), weekdayData: wddBuilder(4).abbrs("Thu", "Th", "R").build()}

	Friday = Weekday{Enum: enums.New("Friday"), weekdayData: wddBuilder(5).abbrs("Fri", "Fr", "F").build()}

	Saturday = Weekday{Enum: enums.New("Saturday"), weekdayData: wddBuilder(6).abbrs("Sat", "Sa", "S").build()}

	Sunday = Weekday{Enum: enums.New("Sunday"), weekdayData: wddBuilder(7).abbrs("Sun", "Su", "U").build()}

	defaultWeekday = Weekday{}
)

func init() {
	enums.AddDefaultWithTransformer(&defaultWeekday, strings.ToLower)
	enums.AddWithAliases(&Monday, Monday.getAbbrs()...)
	enums.AddWithAliases(&Tuesday, Tuesday.getAbbrs()...)
	enums.AddWithAliases(&Wednesday, Wednesday.getAbbrs()...)
	enums.AddWithAliases(&Thursday, Thursday.getAbbrs()...)
	enums.AddWithAliases(&Friday, Friday.getAbbrs()...)
	enums.AddWithAliases(&Saturday, Saturday.getAbbrs()...)
	enums.AddWithAliases(&Sunday, Sunday.getAbbrs()...)
}

func (wd *Weekday) UpdateFrom(found enums.IEnum) {
	src := found.(*Weekday)
	wd.weekdayData = src.weekdayData // wd is Dst
}

func (wd *Weekday) UnmarshalJSON(data []byte) error {
	return enums.UnmarshalJSON(wd, data) // wd is Dst
}

type weekdayData struct {
	dayNumber     int
	abbreviation3 string
	abbreviation2 string
	abbreviation1 string
}

func wddBuilder(dayNumber int) *weekdayData {
	return &weekdayData{dayNumber: dayNumber}
}

func (wdd *weekdayData) build() weekdayData {
	return *wdd
}

func (wdd *weekdayData) abbrs(abbreviation3, abbreviation2, abbreviation1 string) *weekdayData {
	wdd.abbreviation3, wdd.abbreviation2, wdd.abbreviation1 = abbreviation3, abbreviation2, abbreviation1
	return wdd
}

func (wdd *weekdayData) getAbbr(f func() string) string {
	if wdd == nil {
		return ""
	}
	return f()
}

func (wdd *weekdayData) getAbbrs() (abbrs []string) {
	return strs.AppendNonEmpty(strs.AppendNonEmpty(strs.AppendNonEmpty(abbrs,
		wdd.GetAbbreviation1()), wdd.GetAbbreviation2()), wdd.GetAbbreviation3())
}
