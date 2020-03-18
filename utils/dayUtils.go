package utils

import "strconv"

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/18 12:04 AM
 */
func TransDayInt(year, month, day int) string {

	t1 := strconv.Itoa(year)
	t2 := strconv.Itoa(month)
	t3 := strconv.Itoa(day)
	if month < 10 {
		t2 = "0" + t2

	}
	if day < 10 {
		t3 = "0" + t3
	}
	return t1 + t2 + t3

}
