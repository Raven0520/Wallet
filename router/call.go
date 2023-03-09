package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/middleware"
	"github.com/raven0520/wallet/params"
	"github.com/raven0520/wallet/util"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	Params params.Dictionary
	Self   = []string{ // gateway method
		"/upload/square",
	}
)

// Call Call method
func Call(p reflect.Value, context *gin.Context) {
	m := context.Param("method")
	f := util.UpperFirstLetter(m, "_", false)
	method := p.MethodByName(f)
	if !method.IsValid() {
		middleware.ResponseError(context, "", util.ErrMethodNotFound)
		return
	}
	t := method.Type().In(1).Elem()
	s, err := ValidateParams(context, t)
	if err != nil {
		middleware.ResponseVerification(context, err)
		return
	}
	i := reflect.New(t).Interface()
	params, err := BindParams(context, i, s)
	if err != nil {
		middleware.ResponseError(context, "", err)
		return
	}
	data, msg, err := ReadResult(method.Call(params))
	fmt.Println("Data : ", data)
	fmt.Println("Msg : ", msg)
	fmt.Println("Err : ", err)
	if err != nil {
		middleware.ResponseError(context, "", ReadError(err))
		return
	}
	middleware.ResponseSuccess(context, data.Interface(), msg)
}

// ValidateParams 自动调用词典验证数据
func ValidateParams(context *gin.Context, t reflect.Type) (interface{}, error) {
	d := reflect.ValueOf(Params) // 获取所有词典
	s := d.FieldByName(t.Name()) // 实例化对应词典
	if s.IsValid() {             // 需要传递参数时，才验证
		j := reflect.New(s.Type()).Interface() // 获取词典实体
		err := params.ParamsValidator(context, j)
		if err != nil {
			return nil, err
		}
		return j, nil
	}
	return reflect.Value{}, nil
}

// BindParams 绑定参数
func BindParams(context *gin.Context, i interface{}, s interface{}) ([]reflect.Value, error) {
	j, err := json.Marshal(reflect.ValueOf(s).Interface())
	if err != nil {
		fmt.Println("BindParams Marshal Error : ", err.Error())
		return nil, err
	}
	// err = protojson.Unmarshal(j, protoreflect.ValueOf(i).Message().Interface())
	err = json.Unmarshal(j, reflect.ValueOf(i).Interface())
	if err != nil {
		fmt.Println("BindParams Unmarshal Error : ", err.Error())
		return nil, err
	}
	URL := context.Request.URL
	if util.InSliceString(URL.Path, Self) {
		return []reflect.Value{reflect.ValueOf(context), reflect.ValueOf(i)}, nil
	}
	// 传递 Header
	header := metadata.Pairs("Authorization", context.GetHeader("Authorization"), "UUID", context.GetHeader("UUID"), "ClientIP", context.ClientIP())
	md := metadata.NewOutgoingContext(context, header)
	return []reflect.Value{reflect.ValueOf(md), reflect.ValueOf(i)}, nil
}

// ReadResult 读取结果
func ReadResult(result []reflect.Value) (data reflect.Value, message string, err error) {
	// 处理错误
	if result[1].Interface() != nil {
		s, ok := status.FromError(result[1].Interface().(error))
		if ok {
			err = errors.New(s.Message())
		} else {
			err = result[1].Interface().(error)
		}
	}
	// 处理返回值
	st := reflect.TypeOf(result[0].Interface())  // 键
	sv := reflect.ValueOf(result[0].Interface()) // 值
	if st != nil && !sv.IsNil() {
		for i := 0; i < st.Elem().NumField(); i++ {
			field := st.Elem().Field(i)
			value := sv.Elem().Field(i)
			if field.Name == "Data" {
				data = value
			}
			if field.Name == "Message" {
				message = value.String()
			}
		}
	}
	return
}

// ReadError 读取错误返回
func ReadError(err error) error {
	s := err.Error()
	if strings.Contains(s, "not implemented") || strings.Contains(s, "unknown method") {
		return util.ErrMethodNotFound
	}
	if strings.Contains(s, "no such file or directory") {
		return util.ErrFileOrDirectoryNotExists
	}
	return err
}
