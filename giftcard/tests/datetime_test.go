package tests

import (
	"fmt"
	"giftcard/application/service"
	"giftcard/library/datetime"
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"
)

func TestCycleTime(t *testing.T) {
	timeNow := time.Now()

	timeFormat := timeNow.Format("2006-01-02 00:00:00")
	monthFirstDay := service.GetFirstDateOfMonth(timeNow)
	monthLastDay := service.GetLastDateOfMonth(timeNow) + " 23:59:59"
	fmt.Println(timeFormat, monthFirstDay, monthLastDay)
}

func TestTime(t *testing.T) {

	a := carbon.Parse("2022-05-01 10:00:00").ToDateTimeString()
	fmt.Println(a)

	a = datetime.FormatTime("10:00:00")
	fmt.Println(a)
}

func TestTimeZone(t *testing.T) {

	str := carbon.Now(carbon.Bangkok).ToDateTimeString()
	fmt.Println(str)

	ts := datetime.Timestamp(10002)

	ts2 := carbon.Now().SubHour().Timestamp()

	ts3 := carbon.Now(carbon.Shanghai).Timestamp()

	fmt.Println(ts, ts2, ts3)
}
