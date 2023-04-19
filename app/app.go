package app

import (
	"blog_backend/app/router"
	"blog_backend/app/validation"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ReadRouters 读取router下的路由组
// 自动执行，不依赖函数名
func ReadRouters(g *gin.RouterGroup) {
	routes := router.Router{}
	val := reflect.ValueOf(routes)
	// 获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	for i := 0; i < numOfMethod; i++ {
		// 断言特定的方法
		fn, ok := val.Method(i).Interface().(func(g *gin.RouterGroup))
		if !ok {
			continue
		}
		fn(g)
	}
}

func RegisterValidation() {
	vf := validation.ValidateFunc{}
	typ := reflect.TypeOf(vf)
	val := reflect.ValueOf(vf)
	if val.Kind() != reflect.Struct {
		return
	}
	// 获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	if numOfMethod == 0 {
		return
	}
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	count := 0
	for i := 0; i < numOfMethod; i++ {
		fn, ok := val.Method(i).Interface().(func(fl validator.FieldLevel) bool)
		if !ok {
			continue
		}
		// 注册自定义校验函数
		err := validate.RegisterValidation(typ.Method(i).Name, fn)
		if err != nil {
			fmt.Println(typ.Method(i).Name, err)
			continue
		}
		count++
	}
	if count == numOfMethod {
		fmt.Println(">>>自定义校验函数注册完成")
	}
}
