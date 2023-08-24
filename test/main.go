package main

import (
	"log"
	"reflect"
	"time"

	hessian "github.com/LeakIX/go-hessian"
)

type User struct {
	hessian.Package `hessian:"lab.ggw.shs.service.User"`
	Name            string      `hessian:"name"`
	Email           interface{} `hessian:"email"`
	Father          *User       `hessian:"father"`
}

func main() {
	var addr = "http://localhost:8080/simple"

	proxy, err := hessian.NewProxy(&hessian.ProxyConfig{
		Version: hessian.V1,
		URL:     addr,
	})
	if err != nil {
		panic(err)
	}
	// proxy.SetTypeMapping("lab.ggw.shs.service.User", reflect.TypeOf(&User{}))
	proxy.RegisterType(reflect.TypeOf(User{}))

	log.Println(proxy.Invoke("date"))
	log.Println(proxy.Invoke("dateI", time.Now()))
	// log.Println(proxy.Invoke("bytes"))
	// log.Println(proxy.Invoke("bytesI", []byte("Str")))
	// log.Println(proxy.Invoke("str"))
	// log.Println(proxy.Invoke("strI", "abc"))
	// log.Println(proxy.Invoke("strI2", "ggwhite", "this is message"))
	// log.Println(proxy.Invoke("bool"))
	// log.Println(proxy.Invoke("boolI", true))
	// log.Println(proxy.Invoke("boolI", false))
	// log.Println(proxy.Invoke("integer"))
	// log.Println(proxy.Invoke("integerI", int(math.MaxInt32)))
	// log.Println(math.MaxInt32)
	// log.Println(proxy.Invoke("integerI", int(math.MinInt32)))
	// log.Println(math.MinInt32)
	// log.Println(proxy.Invoke("integerI", int8(math.MaxInt8)))
	// log.Println(math.MaxInt8)
	// log.Println(proxy.Invoke("integerI", int8(math.MinInt8)))
	// log.Println(math.MinInt8)
	// log.Println(proxy.Invoke("integerI", int16(math.MaxInt16)))
	// log.Println(math.MaxInt16)
	// log.Println(proxy.Invoke("integerI", int16(math.MinInt16)))
	// log.Println(math.MinInt16)
	// log.Println(proxy.Invoke("integerI", int32(math.MaxInt32)))
	// log.Println(math.MaxInt32)
	// log.Println(proxy.Invoke("integerI", int32(math.MinInt32)))
	// log.Println(math.MinInt32)
	// log.Println(proxy.Invoke("integerI", int64(math.MaxInt64)))
	// log.Println(proxy.Invoke("longX"))
	// log.Println(proxy.Invoke("longI", int64(math.MaxInt64)))
	// log.Println(math.MaxInt64)
	// log.Println(proxy.Invoke("longI", int64(math.MinInt64)))
	// log.Println(math.MinInt64)
	// log.Println(proxy.Invoke("doubleX"))
	// log.Println(proxy.Invoke("doubleI", float32(math.MaxFloat32)))
	// log.Println(math.MaxFloat32)
	// log.Println(proxy.Invoke("doubleI", float64(math.MaxFloat64)))
	// log.Println(math.MaxFloat64)
	// log.Println(proxy.Invoke("objI", User{
	// 	Name:  "ggwhite",
	// 	Email: "ggw.chang@gmail.com",
	// }))
	// log.Println(proxy.Invoke("objI", &User{
	// 	Name:  "ggwhite",
	// 	Email: "ggw.chang@gmail.com",
	// }))
	// log.Println(proxy.Invoke("objI", User{
	// 	Name:  "ggwhite",
	// 	Email: "ggw.chang@gmail.com",
	// 	Father: &User{
	// 		Name:  "father",
	// 		Email: "ggw.chang@gmail.com",
	// 	},
	// }))
	// ans, err := proxy.Invoke("obj")
	// log.Println(ans[0], err)
	// log.Println(proxy.Invoke("map"))
	// log.Println(proxy.Invoke("mapI", map[string]string{
	// 	"KeyA": "ValueA",
	// 	"KeyB": "ValueB",
	// }))
	// log.Println(proxy.Invoke("mapI", map[int]string{
	// 	1: "ValueA",
	// 	2: "ValueB",
	// }))
	// log.Println(proxy.Invoke("mapI", map[bool]string{
	// 	true:  "ValueA",
	// 	false: "ValueB",
	// }))
	// log.Println(proxy.Invoke("mapI", map[bool]float64{
	// 	true:  1233.129,
	// 	false: 321.12,
	// }))
	// log.Println(proxy.Invoke("array"))
	// log.Println(proxy.Invoke("arrayI", []interface{}{123, "abc", "qq"}))
	// log.Println(proxy.Invoke("strarrayI", []string{"123", "abc", "qq"}))
	// log.Println(proxy.Invoke("intarrayI", []int{123, 456, -123}))
}
