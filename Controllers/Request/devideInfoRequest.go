package Request

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/14 9:03 PM
 */
type DevideInfoRequest struct {
	// 手机号
	DevideId string `json:"deviceId"`
	// 密码
	Platform     string `json:"platform"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}
