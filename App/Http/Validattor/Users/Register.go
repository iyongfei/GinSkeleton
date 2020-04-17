package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Base
	Phone string `form:"phone" json:"phone"  bind:"required"`
	Pass  string `form:"pass" json:"pass bind:"required"`
}

func (c *Register) CheckParams(context *gin.Context) {
	var v_form_params *Register = &Register{
		//&CodelistBase{},
	}
	if err := context.ShouldBind(v_form_params); err != nil {
		fmt.Printf("验证器出错")
		return
	}
	fmt.Printf("%#v\n", v_form_params)
	fmt.Println(v_form_params.Name)

	if len(v_form_params.Name) < 3 || len((*v_form_params).Pass) < 6 || len((*v_form_params).Phone) != 11 {
		fmt.Println("参数不符合规定，name、pass、Phone 长度有问题，不允许注册")
		return
	}
	// 验证完成，调用控制器,同时将验证器传递给下一步

	if v_bytes, eror := json.Marshal(v_form_params); eror == nil {
		context.Set("formRegister", v_bytes)
		(&Admin.Users{}).Register(context)
	}
}

//  请记得将表单验证器注册在容器工厂
