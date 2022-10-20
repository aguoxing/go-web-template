package framework

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-web-template/app/model/system/response"
	"go-web-template/configs"
	"go-web-template/global"
	"go-web-template/util"
	"time"
)

type TokenService struct{}

var expireTime = 30 * time.Minute

const MillisMinuteTen = 20 * time.Minute

func (t *TokenService) GetLoginUser(ctx *gin.Context) (loginUser *response.LoginUser, err error) {
	token := ctx.GetHeader(configs.AppConfig.JWT.Header)
	if token != "" {
		claims, err := util.ParseToken(token)
		if err != nil {
			global.Logger.Error(err)
		}
		userKey := "login_tokens:" + claims.LoginUserKey
		jsonData := global.Redis.Get(context.Background(), userKey).String()
		err = json.Unmarshal([]byte(jsonData), &loginUser)
		if err != nil {
			global.Logger.Error(err)
			return nil, err
		}
		return loginUser, nil
	}
	return nil, nil
}

// SetLoginUser 设置用户身份信息
func (t *TokenService) SetLoginUser(user *response.LoginUser) {
	if user != nil && user.UserKey != "" {
		refreshToken(user)
	}
}

// DelLoginUser 删除用户身份信息
func (t *TokenService) DelLoginUser(userKey string) {
	if userKey != "" {
		uk := "login_tokens:" + userKey
		global.Redis.Del(context.Background(), uk)
	}
}

// CreateToken 创建令牌
func (t *TokenService) CreateToken(user *response.LoginUser) (string, error) {
	user.UserKey = uuid.New().String()
	refreshToken(user)

	token, err := util.GenerateToken(user.UserKey)
	if err != nil {
		global.Logger.Error(err, "token签发失败")
	}
	return token, err
}

// VerifyToken 验证令牌有效期，相差不足20分钟，自动刷新缓存
func (t *TokenService) VerifyToken(user *response.LoginUser) {
	et := user.ExpireTime
	s := time.Now().Sub(et)
	if s <= MillisMinuteTen {
		refreshToken(user)
	}
}

func (t *TokenService) SetUserAgent(user *response.LoginUser) {

}

// refreshToken 刷新token
func refreshToken(user *response.LoginUser) {
	user.LoginTime = time.Now()
	user.ExpireTime = user.LoginTime.Add(expireTime)
	// 将用户信息存入redis
	data, err := json.Marshal(&user)
	userKey := "login_tokens:" + user.UserKey
	err = global.Redis.Set(context.Background(), userKey, data, expireTime).Err()
	if err != nil {
		global.Logger.Error(err)
	}
}
