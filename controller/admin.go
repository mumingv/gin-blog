package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

// Login 登录
func (a *AdminController) Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		// 用户输入用户名、密码后的提交操作
	} else {
		// 登录页面的操作内容
		//c.String(http.StatusOK, "hello")
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

// Logout 退出
func (a *AdminController) Logout(c *gin.Context) {

}

// Main 主页
func (a *AdminController) Main(c *gin.Context) {

}

// Config 系统配置信息展示
func (a *AdminController) Config(c *gin.Context) {

}

// AddConfig 系统配置信息更新
func (a *AdminController) AddConfig(c *gin.Context) {

}

// Index 后台首页
func (a *AdminController) Index(c *gin.Context) {

}

// Article 博文添加
func (a *AdminController) Article(c *gin.Context) {

}

// Save 保存
func (a *AdminController) Save(c *gin.Context) {

}

// PostDel 文章删除
func (a *AdminController) PostDel(c *gin.Context) {

}

// Category 类目
func (a *AdminController) Category(c *gin.Context) {

}

// Categoryadd 添加修改类目
func (a *AdminController) CategoryAdd(c *gin.Context) {

}

// CategorySave 保存类目
func (a *AdminController) CategorySave(c *gin.Context) {

}

// CategoryDel 类目删除
func (a *AdminController) CategoryDel(c *gin.Context) {

}
