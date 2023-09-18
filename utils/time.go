package utils

import (
	"math"
	"time"
)

func NowUnix() int64 {
	return time.Now().Unix()
}

// 今天的日期
func TodayDate() string {
	return time.Now().Format("20060102")
}

// 今日0点时间戳
func TodayTime() time.Time {
	today := time.Now().Format("2006-01-02")
	to, _ := time.ParseInLocation("2006-01-02", today, time.Local)
	return to
}

// 零点时间戳
func DayTime(day int) time.Time {
	today := time.Now().Format("2006-01-02")
	to, _ := time.ParseInLocation("2006-01-02", today, time.Local)
	return to.AddDate(0, 0, day)
}

// 今日0点时间戳
func TodayUnix() int64 {
	today := time.Now().Format("2006-01-02")
	to, _ := time.ParseInLocation("2006-01-02", today, time.Local)
	return to.Unix()
}

// 零点时间戳
func DayUnix(day int) int64 {
	today := time.Now().Format("2006-01-02")
	to, _ := time.ParseInLocation("2006-01-02", today, time.Local)
	return to.AddDate(0, 0, day).Unix()
}

func Constellation(birthday string) (string, error) {
	dayArr := []int{20, 19, 21, 20, 21, 22, 23,
		23, 23, 24, 23, 22}

	constellations := [12]string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座",
		"巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座"}

	birthdayTime, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return "", err
	}

	month := birthdayTime.Month()
	day := birthdayTime.Day()

	if day < dayArr[month-1] {
		return constellations[month-1], nil
	}
	if month == 12 {
		return constellations[0], nil
	}
	return constellations[month], nil
}

func BirthdayToAge(date string) (int64, error) {
	birthday, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, err
	}

	time.Now().Sub(birthday)
	year := birthday.Year()
	month := birthday.Month()
	day := birthday.Day()

	y := time.Now().Year()
	m := time.Now().Month()
	d := time.Now().Day()
	age := y - year - 1

	if m > month || (month == m && d >= day) {
		age++
	}
	if age < 0 {
		age = 0
	}
	return int64(age), nil
}

func BirthdayToYear(date string) int {
	birthday, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	return birthday.Year()
}

// 通过生日获取出生时代，80后、90后
func BirthDayToRange(date string) int64 {
	year := BirthdayToYear(date)
	var yearRange int64
	if year < 2000 {
		yearRange = int64(math.Floor(float64((year - 1900) / 10)))
	} else {
		yearRange = int64(math.Floor(float64((year - 2000) / 10)))
	}

	return yearRange * 10
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(t time.Time) time.Time {
	t = t.AddDate(0, 0, -t.Day()+1)
	return GetZeroTime(t)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(t time.Time) time.Time {
	return GetFirstDateOfMonth(t).AddDate(0, 1, 0)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetFirstDateOfThisMonth() int64 {
	t := time.Now()
	ti := GetFirstDateOfMonth(t).AddDate(0, 0, 0)
	fti := ti.Format("2006-01-02 15:04:05")
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", fti, time.Local)
	return p.Unix()
}

// 获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 获取某一天的0点时间
func GetZeroDateTimeByTime(t int64) time.Time {
	time := time.Unix(t, 0)
	return GetZeroTime(time)
}

// 根据传入的天数，获取当天0点时间戳
func GetDateByDay(dayNum int) int64 {
	t := time.Now()
	ti := GetZeroTime(t).AddDate(0, 0, dayNum)
	fti := ti.Format("2006-01-02 15:04:05")
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", fti, time.Local)
	return p.Unix()
}

// 根据传入的天数，获取当天0点时间戳
func GetDateStringByDay(dayNum int) string {
	t := time.Now()
	ti := GetZeroTime(t).AddDate(0, 0, dayNum)
	fti := ti.Format("2006-01-02 15:04:05")
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", fti, time.Local)
	return p.Format("20060102")
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetTimeDate() string {
	t := time.Now()
	fti := t.Format("2006-01-02 15:04:05")
	return fti
}

/*
*
获取本周周一的日期
*/
func GetFirstDateOfWeek() int64 {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("2006-01-02 15:04:05")
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", weekMonday, time.Local)

	return p.Unix()
}

/*
*
获取本周周一的日期
*/
func GetMondayDate() string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("20060102")
	return weekMonday
}

// 获取过去几天的日期
func GetNumFmtDate(f string, days int) []string {
	if f == "" {
		f = "20060102"
	}

	if days == 0 {
		days = 7
	}
	list := make([]string, 0)
	for i := 0; i < days; i++ {
		d := time.Now().AddDate(0, 0, -i).Format(f)
		list = append(list, d)
	}

	return list
}
