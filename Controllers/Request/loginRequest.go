package Request

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/5 3:39 PM
 */
type LoginRequest struct {
	// 手机号
	Phone string `json:"phone"`
	// 密码
	Pwd string `json:"pwd"`
}
