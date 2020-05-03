package demo

import (
	"time"
	"fmt"
)

func main()  {
	t := time.Now()
	fmt.Println("t ", t)
	fmt.Println("t unix", t.Unix())
	fmt.Println("t year", t.Year())
	fmt.Println("t mouth", t.Month())
	fmt.Println("t day", t.Day())
	fmt.Println("t hour", t.Hour())
	fmt.Println("t minute", t.Minute())
	fmt.Println("t second", t.Second())

	datetime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("t datetime ", datetime)

	t1 := int64(1588515795)
	t2 := "2020-05-03 22:23:15"
	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05"
	timeTemplate2 := "2006/01/02 15:04:05"
	timeTemplate3 := "2006-01-02"
	timeTemplate4 := "15:04:05"

	// ======= 将时间戳格式化为日期字符串 =======
	fmt.Println("时-分-秒 年：月：日 ", time.Unix(t1, 0).Format(timeTemplate1))
	fmt.Println("时/分/秒 年：月：日 ", time.Unix(t1, 0).Format(timeTemplate2))
	fmt.Println("时-分-秒 ", time.Unix(t1, 0).Format(timeTemplate3))
	fmt.Println("年：月：日 ", time.Unix(t1, 0).Format(timeTemplate4))

	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t2, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp.Unix())  //
}