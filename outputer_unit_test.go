package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon_String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15").String(),
			want:   "0000-01-01 13:14:15",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15").String(),
			want:   "0001-01-01 13:14:15",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15").String(),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "String()")
		})
	}
}

func TestCarbon_GoString(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "case1",
			actual: Parse("").GoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15").GoString(),
			want:   "time.Date(0, time.January, 1, 13, 14, 15, 0, time.Local)",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15").GoString(),
			want:   "time.Date(1, time.January, 1, 13, 14, 15, 0, time.Local)",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15").GoString(),
			want:   "time.Date(2020, time.August, 5, 13, 14, 15, 0, time.Local)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "GoString()")
		})
	}
}

func TestCarbon_ToString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15").ToString(),
			want:   "0000-01-01 13:14:15 +0805 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15").ToString(),
			want:   "0001-01-01 13:14:15 +0805 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15").ToString(PRC),
			want:   "2020-08-05 13:14:15 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToString()")
		})
	}
}

func TestCarbon_ToMonthString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToMonthString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-01-05").ToMonthString(),
			want:   "January",
		},
		{
			name:   "case3",
			actual: Parse("2020-02-05").ToMonthString(),
			want:   "February",
		},
		{
			name:   "case4",
			actual: Parse("2020-03-05").ToMonthString(),
			want:   "March",
		},
		{
			name:   "case5",
			actual: Parse("2020-04-05").ToMonthString(),
			want:   "April",
		},
		{
			name:   "case6",
			actual: Parse("2020-05-05").ToMonthString(),
			want:   "May",
		},
		{
			name:   "case7",
			actual: Parse("2020-06-05").ToMonthString(),
			want:   "June",
		},
		{
			name:   "case8",
			actual: Parse("2020-07-05").ToMonthString(),
			want:   "July",
		},
		{
			name:   "case9",
			actual: Parse("2020-08-05").ToMonthString(),
			want:   "August",
		},
		{
			name:   "case10",
			actual: Parse("2020-09-05").ToMonthString(),
			want:   "September",
		},
		{
			name:   "case11",
			actual: Parse("2020-10-05").ToMonthString(),
			want:   "October",
		},
		{
			name:   "case12",
			actual: Parse("2020-11-05").ToMonthString(),
			want:   "November",
		},
		{
			name:   "case13",
			actual: Parse("2020-12-05").ToMonthString(PRC),
			want:   "December",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToMonthString()")
		})
	}
}

func TestCarbon_ToShortMonthString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortMonthString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-01-05").ToShortMonthString(),
			want:   "Jan",
		},
		{
			name:   "case3",
			actual: Parse("2020-02-05").ToShortMonthString(),
			want:   "Feb",
		},
		{
			name:   "case4",
			actual: Parse("2020-03-05").ToShortMonthString(),
			want:   "Mar",
		},
		{
			name:   "case5",
			actual: Parse("2020-04-05").ToShortMonthString(),
			want:   "Apr",
		},
		{
			name:   "case6",
			actual: Parse("2020-05-05").ToShortMonthString(),
			want:   "May",
		},
		{
			name:   "case7",
			actual: Parse("2020-06-05").ToShortMonthString(),
			want:   "Jun",
		},
		{
			name:   "case8",
			actual: Parse("2020-07-05").ToShortMonthString(),
			want:   "Jul",
		},
		{
			name:   "case9",
			actual: Parse("2020-08-05").ToShortMonthString(),
			want:   "Aug",
		},
		{
			name:   "case10",
			actual: Parse("2020-09-05").ToShortMonthString(),
			want:   "Sep",
		},
		{
			name:   "case11",
			actual: Parse("2020-10-05").ToShortMonthString(),
			want:   "Oct",
		},
		{
			name:   "case12",
			actual: Parse("2020-11-05").ToShortMonthString(),
			want:   "Nov",
		},
		{
			name:   "case13",
			actual: Parse("2020-12-05").ToShortMonthString(PRC),
			want:   "Dec",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortMonthString()")
		})
	}
}

func TestCarbon_ToWeekString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToWeekString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-01").ToWeekString(),
			want:   "Saturday",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-02").ToWeekString(),
			want:   "Sunday",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-03").ToWeekString(),
			want:   "Monday",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-04").ToWeekString(),
			want:   "Tuesday",
		},
		{
			name:   "case6",
			actual: Parse("2020-08-05").ToWeekString(),
			want:   "Wednesday",
		},
		{
			name:   "case7",
			actual: Parse("2020-08-06").ToWeekString(),
			want:   "Thursday",
		},
		{
			name:   "case8",
			actual: Parse("2020-08-07").ToWeekString(PRC),
			want:   "Friday",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToWeekString()")
		})
	}
}

func TestCarbon_ToShortWeekString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortWeekString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-01").ToShortWeekString(),
			want:   "Sat",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-02").ToShortWeekString(),
			want:   "Sun",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-03").ToShortWeekString(),
			want:   "Mon",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-04").ToShortWeekString(),
			want:   "Tue",
		},
		{
			name:   "case6",
			actual: Parse("2020-08-05").ToShortWeekString(),
			want:   "Wed",
		},
		{
			name:   "case7",
			actual: Parse("2020-08-06").ToShortWeekString(),
			want:   "Thu",
		},
		{
			name:   "case8",
			actual: Parse("2020-08-07").ToShortWeekString(PRC),
			want:   "Fri",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortWeekString()")
		})
	}
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDayDateTimeString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05 13:14:15").ToDayDateTimeString(),
			want:   "Wed, Aug 5, 2020 1:14 PM",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDayDateTimeString(PRC),
			want:   "Wed, Aug 5, 2020 12:00 AM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDayDateTimeString()")
		})
	}
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateTimeString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05 13:14:15").ToDateTimeString(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateTimeString(PRC),
			want:   "2020-08-05 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateTimeString()")
		})
	}
}

func TestCarbon_ToDateTimeMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateTimeMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateTimeMilliString(PRC),
			want:   "2020-08-05 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateTimeMilliString()")
		})
	}
}

func TestCarbon_ToDateTimeMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateTimeMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateTimeMicroString(PRC),
			want:   "2020-08-05 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateTimeMicroString()")
		})
	}
}

func TestCarbon_ToDateTimeNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateTimeNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString(),
			want:   "2020-08-05 13:14:15.999999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateTimeNanoString(PRC),
			want:   "2020-08-05 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateTimeNanoString()")
		})
	}
}

func TestCarbon_ToShortDateTimeString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateTimeString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("xxx").ToShortDateTimeString(),
			want:   "",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString(),
			want:   "20200805131415",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").ToShortDateTimeString(PRC),
			want:   "20200805000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateTimeString()")
		})
	}
}

func TestCarbon_ToShortDateTimeMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateTimeMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString(),
			want:   "20200805131415.999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateTimeMilliString(PRC),
			want:   "20200805000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateTimeMilliString()")
		})
	}
}

func TestCarbon_ToShortDateTimeMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateTimeMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString(),
			want:   "20200805131415.999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateTimeMicroString(PRC),
			want:   "20200805000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateTimeMicroString()")
		})
	}
}

func TestCarbon_ToShortDateTimeNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateTimeNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString(),
			want:   "20200805131415.999999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateTimeNanoString(PRC),
			want:   "20200805000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateTimeNanoString()")
		})
	}
}

func TestCarbon_ToDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateString(),
			want:   "2020-08-05",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateString(PRC),
			want:   "2020-08-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateString()")
		})
	}
}

func TestCarbon_ToDateMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateMilliString(),
			want:   "2020-08-05.999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateMilliString(PRC),
			want:   "2020-08-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateMilliString()")
		})
	}
}

func TestCarbon_ToDateMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateMicroString(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateMicroString(PRC),
			want:   "2020-08-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateMicroString()")
		})
	}
}

func TestCarbon_ToDateNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToDateNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToDateNanoString(),
			want:   "2020-08-05.999999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToDateNanoString(PRC),
			want:   "2020-08-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToDateNanoString()")
		})
	}
}

func TestCarbon_ToShortDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateString(),
			want:   "20200805",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateString(PRC),
			want:   "20200805",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateString()")
		})
	}
}

func TestCarbon_ToShortDateMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateMilliString(),
			want:   "20200805.999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateMilliString(PRC),
			want:   "20200805",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateMilliString()")
		})
	}
}

func TestCarbon_ToShortDateNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateNanoString(),
			want:   "20200805.999999999",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").ToShortDateNanoString(PRC),
			want:   "20200805",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateNanoString()")
		})
	}
}

func TestCarbon_ToShortDateMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortDateMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15").ToShortDateMicroString(),
			want:   "00000101",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15").ToShortDateMicroString(),
			want:   "00010101",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateMicroString(),
			want:   "20200805.999999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToShortDateMicroString(PRC),
			want:   "20200805",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortDateMicroString()")
		})
	}
}

func TestCarbon_ToTimeString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToTimeString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15").ToTimeString(),
			want:   "13:14:15",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15").ToTimeString(),
			want:   "13:14:15",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15").ToTimeString(),
			want:   "13:14:15",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToTimeString(PRC),
			want:   "00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToTimeString()")
		})
	}
}

func TestCarbon_ToTimeMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToTimeMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToTimeMilliString(),
			want:   "13:14:15.999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToTimeMilliString(),
			want:   "13:14:15.999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString(),
			want:   "13:14:15.999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToTimeMilliString(PRC),
			want:   "00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToTimeMilliString()")
		})
	}
}

func TestCarbon_ToTimeMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToTimeMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToTimeMicroString(),
			want:   "13:14:15.999999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToTimeMicroString(),
			want:   "13:14:15.999999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString(),
			want:   "13:14:15.999999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToTimeMicroString(PRC),
			want:   "00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToTimeMicroString()")
		})
	}
}

func TestCarbon_ToTimeNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToTimeNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToTimeNanoString(),
			want:   "13:14:15.999999999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToTimeNanoString(),
			want:   "13:14:15.999999999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString(),
			want:   "13:14:15.999999999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToTimeNanoString(PRC),
			want:   "00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToTimeNanoString()")
		})
	}
}

func TestCarbon_ToShortTimeString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortTimeString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToShortTimeString(),
			want:   "131415",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToShortTimeString(),
			want:   "131415",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToShortTimeString(),
			want:   "131415",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToShortTimeString(PRC),
			want:   "000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortTimeString()")
		})
	}
}

func TestCarbon_ToShortTimeMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortTimeMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToShortTimeMilliString(),
			want:   "131415.999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToShortTimeMilliString(),
			want:   "131415.999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString(),
			want:   "131415.999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToShortTimeMilliString(PRC),
			want:   "000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortTimeMilliString()")
		})
	}
}

func TestCarbon_ToShortTimeMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortTimeMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToShortTimeMicroString(),
			want:   "131415.999999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToShortTimeMicroString(),
			want:   "131415.999999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString(),
			want:   "131415.999999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToShortTimeMicroString(PRC),
			want:   "000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortTimeMicroString()")
		})
	}
}

func TestCarbon_ToShortTimeNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToShortTimeNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToShortTimeNanoString(),
			want:   "131415.999999999",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToShortTimeNanoString(),
			want:   "131415.999999999",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString(),
			want:   "131415.999999999",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToShortTimeNanoString(PRC),
			want:   "000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortTimeNanoString()")
		})
	}
}

func TestCarbon_ToAtomString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToAtomString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToAtomString(),
			want:   "0000-01-01T13:14:15+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToAtomString(),
			want:   "0001-01-01T13:14:15+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToAtomString(),
			want:   "2020-08-05T13:14:15+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToAtomString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToAtomString()")
		})
	}
}

func TestCarbon_ToAnsicString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToAnsicString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToAnsicString(),
			want:   "Sat Jan  1 13:14:15 0000",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToAnsicString(),
			want:   "Mon Jan  1 13:14:15 0001",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToAnsicString(),
			want:   "Wed Aug  5 13:14:15 2020",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToAnsicString(PRC),
			want:   "Wed Aug  5 00:00:00 2020",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToAnsicString()")
		})
	}
}

func TestCarbon_ToCookieString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToCookieString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToCookieString(),
			want:   "Saturday, 01-Jan-0000 13:14:15 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToCookieString(),
			want:   "Monday, 01-Jan-0001 13:14:15 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToCookieString(),
			want:   "Wednesday, 05-Aug-2020 13:14:15 CST",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToCookieString(PRC),
			want:   "Wednesday, 05-Aug-2020 00:00:00 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToCookieString()")
		})
	}
}

func TestCarbon_ToRssString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRssString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRssString(),
			want:   "Sat, 01 Jan 0000 13:14:15 +0805",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRssString(),
			want:   "Mon, 01 Jan 0001 13:14:15 +0805",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRssString(),
			want:   "Wed, 05 Aug 2020 13:14:15 +0800",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRssString(PRC),
			want:   "Wed, 05 Aug 2020 00:00:00 +0800",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRssString()")
		})
	}
}

func TestCarbon_ToW3cString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToW3cString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToW3cString(),
			want:   "0000-01-01T13:14:15+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToW3cString(),
			want:   "0001-01-01T13:14:15+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToW3cString(),
			want:   "2020-08-05T13:14:15+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToW3cString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToW3cString()")
		})
	}
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToUnixDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToUnixDateString(),
			want:   "Sat Jan  1 13:14:15 LMT 0000",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToUnixDateString(),
			want:   "Mon Jan  1 13:14:15 LMT 0001",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToUnixDateString(),
			want:   "Wed Aug  5 13:14:15 CST 2020",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToUnixDateString(PRC),
			want:   "Wed Aug  5 00:00:00 CST 2020",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToUnixDateString()")
		})
	}
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRubyDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRubyDateString(),
			want:   "Sat Jan 01 13:14:15 +0805 0000",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRubyDateString(),
			want:   "Mon Jan 01 13:14:15 +0805 0001",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRubyDateString(),
			want:   "Wed Aug 05 13:14:15 +0800 2020",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRubyDateString(PRC),
			want:   "Wed Aug 05 00:00:00 +0800 2020",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRubyDateString()")
		})
	}
}

func TestCarbon_ToKitchenString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToKitchenString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToKitchenString(),
			want:   "1:14PM",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToKitchenString(),
			want:   "1:14PM",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToKitchenString(),
			want:   "1:14PM",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToKitchenString(PRC),
			want:   "12:00AM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToKitchenString()")
		})
	}
}

func TestCarbon_ToIso8601String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601String(),
			want:   "0000-01-01T13:14:15+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601String(),
			want:   "0001-01-01T13:14:15+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601String(),
			want:   "2020-08-05T13:14:15+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601String(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601String()")
		})
	}
}

func TestCarbon_ToIso8601MilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601MilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601MilliString(),
			want:   "0000-01-01T13:14:15.999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601MilliString(),
			want:   "0001-01-01T13:14:15.999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString(),
			want:   "2020-08-05T13:14:15.999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601MilliString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601MilliString()")
		})
	}
}

func TestCarbon_ToIso8601MicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601MicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601MicroString(),
			want:   "0000-01-01T13:14:15.999999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601MicroString(),
			want:   "0001-01-01T13:14:15.999999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString(),
			want:   "2020-08-05T13:14:15.999999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601MicroString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601MicroString()")
		})
	}
}

func TestCarbon_ToIso8601NanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601NanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601NanoString(),
			want:   "0000-01-01T13:14:15.999999999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601NanoString(),
			want:   "0001-01-01T13:14:15.999999999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString(),
			want:   "2020-08-05T13:14:15.999999999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601NanoString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601NanoString()")
		})
	}
}

func TestCarbon_ToIso8601ZuluString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601ZuluString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601ZuluString(),
			want:   "0000-01-01T13:14:15Z",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601ZuluString(),
			want:   "0001-01-01T13:14:15Z",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString(),
			want:   "2020-08-05T13:14:15Z",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601ZuluString(PRC),
			want:   "2020-08-05T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601ZuluString()")
		})
	}
}

func TestCarbon_ToIso8601ZuluMilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601ZuluMilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601ZuluMilliString(),
			want:   "0000-01-01T13:14:15.999Z",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601ZuluMilliString(),
			want:   "0001-01-01T13:14:15.999Z",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString(),
			want:   "2020-08-05T13:14:15.999Z",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601ZuluMilliString(PRC),
			want:   "2020-08-05T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601ZuluMilliString()")
		})
	}
}

func TestCarbon_ToIso8601ZuluMicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601ZuluMicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601ZuluMicroString(),
			want:   "0000-01-01T13:14:15.999999Z",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601ZuluMicroString(),
			want:   "0001-01-01T13:14:15.999999Z",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString(),
			want:   "2020-08-05T13:14:15.999999Z",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601ZuluMicroString(PRC),
			want:   "2020-08-05T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601ZuluMicroString()")
		})
	}
}

func TestCarbon_ToIso8601ZuluNanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToIso8601ZuluNanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToIso8601ZuluNanoString(),
			want:   "0000-01-01T13:14:15.999999999Z",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToIso8601ZuluNanoString(),
			want:   "0001-01-01T13:14:15.999999999Z",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString(),
			want:   "2020-08-05T13:14:15.999999999Z",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToIso8601ZuluNanoString(PRC),
			want:   "2020-08-05T00:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToIso8601ZuluNanoString()")
		})
	}
}

func TestCarbon_ToRfc822String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc822String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc822String(),
			want:   "01 Jan 00 13:14 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc822String(),
			want:   "01 Jan 01 13:14 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc822String(),
			want:   "05 Aug 20 13:14 CST",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc822String(PRC),
			want:   "05 Aug 20 00:00 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc822String()")
		})
	}
}

func TestCarbon_ToRfc822zString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc822zString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc822zString(),
			want:   "01 Jan 00 13:14 +0805",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc822zString(),
			want:   "01 Jan 01 13:14 +0805",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc822zString(),
			want:   "05 Aug 20 13:14 +0800",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc822zString(PRC),
			want:   "05 Aug 20 00:00 +0800",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc822zString()")
		})
	}
}

func TestCarbon_ToRfc850String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc850String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc850String(),
			want:   "Saturday, 01-Jan-00 13:14:15 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc850String(),
			want:   "Monday, 01-Jan-01 13:14:15 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc850String(),
			want:   "Wednesday, 05-Aug-20 13:14:15 CST",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc850String(PRC),
			want:   "Wednesday, 05-Aug-20 00:00:00 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc850String()")
		})
	}
}

func TestCarbon_ToRfc1036String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc1036String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc1036String(),
			want:   "Sat, 01 Jan 00 13:14:15 +0805",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc1036String(),
			want:   "Mon, 01 Jan 01 13:14:15 +0805",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc1036String(),
			want:   "Wed, 05 Aug 20 13:14:15 +0800",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc1036String(PRC),
			want:   "Wed, 05 Aug 20 00:00:00 +0800",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc1036String()")
		})
	}
}

func TestCarbon_ToRfc1123String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc1123String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc1123String(),
			want:   "Sat, 01 Jan 0000 13:14:15 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc1123String(),
			want:   "Mon, 01 Jan 0001 13:14:15 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc1123String(),
			want:   "Wed, 05 Aug 2020 13:14:15 CST",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc1123String(PRC),
			want:   "Wed, 05 Aug 2020 00:00:00 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc1123String()")
		})
	}
}

func TestCarbon_ToRfc1123zString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc1123zString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc1123zString(),
			want:   "Sat, 01 Jan 0000 13:14:15 +0805",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc1123zString(),
			want:   "Mon, 01 Jan 0001 13:14:15 +0805",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc1123zString(),
			want:   "Wed, 05 Aug 2020 13:14:15 +0800",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc1123zString(PRC),
			want:   "Wed, 05 Aug 2020 00:00:00 +0800",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc1123String()")
		})
	}
}

func TestCarbon_ToRfc2822String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc2822String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc2822String(),
			want:   "Sat, 01 Jan 0000 13:14:15 +0805",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc2822String(),
			want:   "Mon, 01 Jan 0001 13:14:15 +0805",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc2822String(),
			want:   "Wed, 05 Aug 2020 13:14:15 +0800",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc2822String(PRC),
			want:   "Wed, 05 Aug 2020 00:00:00 +0800",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc2822String()")
		})
	}
}

func TestCarbon_ToRfc3339String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc3339String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc3339String(),
			want:   "0000-01-01T13:14:15+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc3339String(),
			want:   "0001-01-01T13:14:15+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc3339String(),
			want:   "2020-08-05T13:14:15+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc3339String(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc3339String()")
		})
	}
}

func TestCarbon_ToRfc3339MilliString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc3339MilliString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc3339MilliString(),
			want:   "0000-01-01T13:14:15.999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc3339MilliString(),
			want:   "0001-01-01T13:14:15.999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc3339MilliString(),
			want:   "2020-08-05T13:14:15.999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc3339MilliString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc3339MilliString()")
		})
	}
}

func TestCarbon_ToRfc3339MicroString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc3339MicroString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc3339MicroString(),
			want:   "0000-01-01T13:14:15.999999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc3339MicroString(),
			want:   "0001-01-01T13:14:15.999999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc3339MicroString(),
			want:   "2020-08-05T13:14:15.999999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc3339MicroString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc3339MicroString()")
		})
	}
}

func TestCarbon_ToRfc3339NanoString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc3339NanoString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc3339NanoString(),
			want:   "0000-01-01T13:14:15.999999999+08:05",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc3339NanoString(),
			want:   "0001-01-01T13:14:15.999999999+08:05",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc3339NanoString(),
			want:   "2020-08-05T13:14:15.999999999+08:00",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc3339NanoString(PRC),
			want:   "2020-08-05T00:00:00+08:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc3339NanoString()")
		})
	}
}

func TestCarbon_ToRfc7231String(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToRfc7231String(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToRfc7231String(),
			want:   "Sat, 01 Jan 0000 13:14:15 LMT",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToRfc7231String(),
			want:   "Mon, 01 Jan 0001 13:14:15 LMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToRfc7231String(),
			want:   "Wed, 05 Aug 2020 13:14:15 CST",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToRfc7231String(PRC),
			want:   "Wed, 05 Aug 2020 00:00:00 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToRfc7231String()")
		})
	}
}

func TestCarbon_ToFormattedDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToFormattedDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToFormattedDateString(),
			want:   "Jan 1, 0000",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToFormattedDateString(),
			want:   "Jan 1, 0001",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToFormattedDateString(),
			want:   "Aug 5, 2020",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToFormattedDateString(PRC),
			want:   "Aug 5, 2020",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToFormattedDateString()")
		})
	}
}

func TestCarbon_ToFormattedDayDateString(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").ToFormattedDayDateString(),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("0000-01-01 13:14:15.999999999").ToFormattedDayDateString(),
			want:   "Sat, Jan 1, 0000",
		},
		{
			name:   "case3",
			actual: Parse("0001-01-01 13:14:15.999999999").ToFormattedDayDateString(),
			want:   "Mon, Jan 1, 0001",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15.999999999").ToFormattedDayDateString(),
			want:   "Wed, Aug 5, 2020",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").ToFormattedDayDateString(PRC),
			want:   "Wed, Aug 5, 2020",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToFormattedDayDateString()")
		})
	}
}

func TestCarbon_Layout(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").Layout(""),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05 13:14:15").Layout("2006年01月02日"),
			want:   "2020年08月05日",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05 13:14:15").Layout("Mon, 02 Jan 2006 15:04:05 GMT"),
			want:   "Wed, 05 Aug 2020 13:14:15 GMT",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 13:14:15").Layout(DateTimeLayout, PRC),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Layout()")
		})
	}
}

func TestCarbon_Format(t *testing.T) {
	tests := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Parse("").Format(""),
			want:   "",
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05 01:14:15").Format("D"),
			want:   "Wed",
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05 01:14:15").Format("l"),
			want:   "Wednesday",
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05 01:14:15").Format("F"),
			want:   "August",
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05 01:14:15").Format("M"),
			want:   "Aug",
		},
		{
			name:   "case6",
			actual: Parse("2020-08-05 13:14:15").Format("Y年m月d日"),
			want:   "2020年08月05日",
		},
		{
			name:   "case7",
			actual: Parse("2020-08-05 13:14:15").Format("j"),
			want:   "5",
		},
		{
			name:   "case8",
			actual: Parse("2020-08-05 13:14:15").Format("W"),
			want:   "32",
		},
		{
			name:   "case9",
			actual: Parse("2020-08-05 13:14:15").Format("F"),
			want:   "August",
		},
		{
			name:   "case10",
			actual: Parse("2020-08-05 13:14:15").Format("F"),
			want:   "August",
		},
		{
			name:   "case11",
			actual: Parse("2020-08-05 13:14:15").Format("N"),
			want:   "03",
		},
		{
			name:   "case12",
			actual: Parse("2020-08-05 13:14:15").Format("L"),
			want:   "1",
		},
		{
			name:   "case13",
			actual: Parse("2021-08-05 01:14:15").Format("L"),
			want:   "0",
		},
		{
			name:   "case14",
			actual: Parse("2020-08-05 01:14:15").Format("G"),
			want:   "1",
		},
		{
			name:   "case15",
			actual: Parse("2020-08-05 13:14:15").Format("U"),
			want:   "1596604455",
		},
		{
			name:   "case16",
			actual: Parse("2020-08-05 13:14:15").Format("V"),
			want:   "1596604455000",
		},
		{
			name:   "case17",
			actual: Parse("2020-08-05 13:14:15").Format("X"),
			want:   "1596604455000000",
		},
		{
			name:   "case18",
			actual: Parse("2020-08-05 13:14:15").Format("Z"),
			want:   "1596604455000000000",
		},
		{
			name:   "case19",
			actual: Parse("2020-08-05 13:14:15.999").Format("v"),
			want:   "999",
		},
		{
			name:   "case20",
			actual: Parse("2020-08-05 13:14:15.999999").Format("u"),
			want:   "999999",
		},
		{
			name:   "case21",
			actual: Parse("2020-08-05 13:14:15.999999999").Format("x"),
			want:   "999999999",
		},
		{
			name:   "case22",
			actual: Parse("2020-08-05 13:14:15").Format("w"),
			want:   "2",
		},
		{
			name:   "case23",
			actual: Parse("2020-08-05 13:14:15").Format("t"),
			want:   "31",
		},
		{
			name:   "case24",
			actual: Parse("2020-08-05 13:14:15").Format("z"),
			want:   "217",
		},
		{
			name:   "case25",
			actual: Parse("2020-08-05 13:14:15", PRC).Format("e"),
			want:   "PRC",
		},
		{
			name:   "case26",
			actual: Parse("2020-08-05 13:14:15").Format("Q"),
			want:   "3",
		},
		{
			name:   "case27",
			actual: Parse("2020-08-05 13:14:15").Format("C"),
			want:   "21",
		},
		{
			name:   "case28",
			actual: Parse("2020-08-05 13:14:15").Format("jS"),
			want:   "5th",
		},
		{
			name:   "case29",
			actual: Parse("2020-08-22 13:14:15").Format("jS"),
			want:   "22nd",
		},
		{
			name:   "case30",
			actual: Parse("2020-08-23 13:14:15").Format("jS"),
			want:   "23rd",
		},
		{
			name:   "case31",
			actual: Parse("2020-08-31 13:14:15").Format("jS"),
			want:   "31st",
		},
		{
			name:   "case32",
			actual: Parse("2020-08-31 13:14:15").Format("I\\t \\i\\s Y-m-d H:i:s"),
			want:   "It is 2020-08-31 13:14:15",
		},
		{
			name:   "case33",
			actual: Parse("2020-08-05 13:14:15").Format("上次上报时间:Y-m-d H:i:s，请每日按时打卡"),
			want:   "上次上报时间:2020-08-05 13:14:15，请每日按时打卡",
		},
		{
			name:   "case34",
			actual: Parse("2020-08-05 01:14:15").Format(DateTimeFormat, PRC),
			want:   "2020-08-05 01:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Format()")
		})
	}
}

func TestCarbon_ToStdTime(t *testing.T) {
	expected := time.Now().Format(DateTimeFormat)
	actual := Now().ToStdTime().Format(DateTimeFormat)
	assert.Equal(t, expected, actual)
}

// https://github.com/golang-module/carbon/issues/200
func TestCarbon_Issue200(t *testing.T) {
	tests1 := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case1",
			actual: Now().StartOfWeek().ToWeekString(),
			want:   "Sunday",
		},
		{
			name:   "case2",
			actual: Now().SetWeekStartsAt(Monday).StartOfWeek().ToWeekString(),
			want:   "Monday",
		},
		{
			name:   "case3",
			actual: Now().SetWeekStartsAt(Wednesday).StartOfWeek().ToWeekString(PRC),
			want:   "Wednesday",
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToWeekString()")
		})
	}

	tests2 := []struct {
		name   string
		actual string
		want   string
	}{
		{
			name:   "case4",
			actual: Now().StartOfWeek().ToShortWeekString(),
			want:   "Sun",
		},
		{
			name:   "case5",
			actual: Now().SetWeekStartsAt(Monday).StartOfWeek().ToShortWeekString(),
			want:   "Mon",
		},
		{
			name:   "case6",
			actual: Now().SetWeekStartsAt(Wednesday).StartOfWeek().ToShortWeekString(PRC),
			want:   "Wed",
		},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "ToShortWeekString()")
		})
	}
}
