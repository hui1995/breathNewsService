package Response

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/5 4:01 PM
 */
type AccountResponse struct {
	// 手机号

	UserName    string
	RealNum     int
	Money       float64
	FreezeMoney float64
	ActiveMoney float64
	Token       string `json:"token"`
}
