package admin

import (
	"blog/fox/str"
	"encoding/json"
	"fmt"
	"blog/fox"
	"blog/service/admin"
	"blog/fox/config"
)
//后台 Session 处理
type AdminSession struct {
	fox.Controller              //继承 控制器
	Session *admin.AdminSession //当前登录用户信息
}
// session 填充
func (c *AdminSession) SessionSet(session *admin.AdminSession) {
	SESSION_NAME := config.String("session_name")
	//存入 Session
	str2, err := str.JsonEnCode(session)
	if err != nil {
		fmt.Println("session:" + err.Error())
	}
	//fmt.Println("str => ?", str2)
	c.SetSession(SESSION_NAME, str2)
}
//获取
func (c *AdminSession) SessionGet() (*admin.AdminSession, error) {
	SESSION_NAME := config.String("session_name")
	session, ok := c.GetSession(SESSION_NAME).(string)
	fmt.Println("session:", session)
	fmt.Println("ok bool:", ok)
	if !ok {
		return nil, fox.NewError("Session 不存在")
	}
	if ok && session == "" {
		return nil, fox.NewError("Session 为空")
	}
	Sess := admin.NewAdminSessionService()
	err := json.Unmarshal([]byte(session), &Sess)
	if err != nil {
		return nil, fox.NewError("Session 序列号转换错误. " + err.Error())
	}
	return Sess, nil
}
//删除
func (c *AdminSession) SessionDel() {
	SESSION_NAME := config.String("session_name")
	c.DelSession(SESSION_NAME)
}
//错误信息
func (c *AdminSession) Error(key string, def ...map[string]interface{}) {
	if c.IsAjax() {
		c.ErrorJson(key, def...)
	} else {
		c.Data["content"] = key
		c.TplName = "error/404.html"
	}
}
func (c *AdminSession) Success(key string, def ...map[string]interface{}) {
	c.Data["content"] = key
	if c.IsAjax() {
		c.SuccessJson(key, def...)
	} else {
		c.TplName = "error/success.html"
	}
}
//错误信息
func (c *AdminSession) ErrorJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 0
	c.Data["json"] = m
	c.ServeJSON()
}
func (c *AdminSession) SuccessJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 1
	c.Data["json"] = m
	c.ServeJSON()
}