package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"` //后面这个json是一个标签 格式为`key:value`
	Message string `json:"msg"`  //在json序列化的时候 会把结构体中对应变量的名字
} //变为标签的名字
//比如这里 在转化为json输出的时候 Message的key会变成msg {"code":200,"msg":"success"}

func main() {
	var res Result = Result{Code: 200, Message: "success"}
	jsons, errors := json.Marshal(res)
	if errors != nil {
		fmt.Println("error")
		fmt.Println(errors)
		return
	}
	fmt.Println(string(jsons))
	//注意 返回的jsons是以序列化的形式存储的
	//也就是一个int类型的数组
	//要用string将其变成可读形式

	//反序列化
	var res2 Result
	errs := json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2 :", res2)

	//用map生成json也和这个差不多
	//interface是一种万能类型
	//可以理解成java的object
	res3 := make(map[string]interface{})
	res3["code"] = 200
	res3["msg"] = "success"
	res3["data"] = map[string]interface{}{
		"name": "whz",
		"age":  18,
	}

	j, _ := json.Marshal(res3)
	fmt.Println(string(j))
}
