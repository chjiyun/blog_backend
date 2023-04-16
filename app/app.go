package app

import (
	"blog_backend/app/router"
	"blog_backend/app/util"
	"blog_backend/app/validation"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ReadRouters 读取router下的路由组
func ReadRouters(g *gin.RouterGroup) {
	var funcNames = util.GetFileBasename("app/router", []string{"go"})
	if len(funcNames) == 0 {
		return
	}
	// 获取反射值
	value := reflect.ValueOf(&router.Router{})
	in := []reflect.Value{reflect.ValueOf(g)}
	for _, fnName := range funcNames {
		fn := value.MethodByName(fnName) //通过反射获取它对应的函数
		if fn.Kind() != reflect.Func || fn.IsNil() {
			continue
		}
		fn.Call(in)
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
