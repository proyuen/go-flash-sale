package token

import "time"

type Maker interface {
	// 生成一个token
	CreateToken(username string, duration time.Duration) (string, error)

	// 检查token 是否有效 有效的话返回payload
	VerifyToken(token string) (*Payload, error)
}
