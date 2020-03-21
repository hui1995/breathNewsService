package Request

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/21 9:05 PM
 */
type AlipayRequest struct {
	Alipay string `json:"alipay"`
	// 密码
	RealName string `json:"realname"`
}
