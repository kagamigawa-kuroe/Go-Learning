package main

//在看go的反射之前
//需要先了解一下go中的接口类
//go中空接口可以作为类似java中object对象使用
//这得益于go底层对于空接口和普通接口做了区分

//普通接口的数据结构如下
// type iface struct {
// 	tab  *itab
// 	data unsafe.Pointer
// }

// type itab struct {
// 	inter *interfacetype
// 	_type *_type
// 	hash  uint32 // copy of _type.hash. Used for type switches.
// 	_     [4]byte
// 	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
// }

//而空接口的定义如下
// type eface struct {
// 	_type *_type
// 	data  unsafe.Pointer
// }

//type数据结构中包函数变量类型的信息
//而data中则是初始数据

// type _type struct {
// 	size       uintptr // 类型占用内存大小
// 	ptrdata    uintptr // 包含所有指针的内存前缀大小
// 	hash       uint32  // 类型 hash
// 	tflag      tflag   // 标记位，主要用于反射
// 	align      uint8   // 对齐字节信息
// 	fieldAlign uint8   // 当前结构字段的对齐字节数
// 	kind       uint8   // 基础类型枚举值
// 	equal func(unsafe.Pointer, unsafe.Pointer) bool // 比较两个形参对应对象的类型是否相等
// 	gcdata    *byte    // GC 类型的数据
// 	str       nameOff  // 类型名称字符串在二进制文件段中的偏移量
// 	ptrToThis typeOff  // 类型元信息指针在二进制文件段中的偏移量
// }

// 这就是为什么空接口可以作为一个范形来使用
// 接下来就可以看看go的反射

// 两个反射包中很基础的函数
// func TypeOf(i interface{}) Type
// func ValueOf(i interface{}) Value
// 用来判断数据的类型和值

// typeof的源码
// func TypeOf(i interface{}) Type {
//     eface := *(*emptyInterface)(unsafe.Pointer(&i))
//     return toType(eface.typ)
// }

// 返回了一个type接口
// 这个接口包含了诸多函数
// 可以用来确认类型的信息
// type Type interface {
//     // 所有的类型都可以调用下面这些函数

//     // 此类型的变量对齐后所占用的字节数
//     Align() int

//     // 如果是 struct 的字段，对齐后占用的字节数
//     FieldAlign() int

//     // 返回类型方法集里的第 `i` (传入的参数)个方法
//     Method(int) Method

//     // 通过名称获取方法
//     MethodByName(string) (Method, bool)

//     // 获取类型方法集里导出的方法个数
//     NumMethod() int

//     // 类型名称
//     Name() string

//     // 返回类型所在的路径，如：encoding/base64
//     PkgPath() string

//     // 返回类型的大小，和 unsafe.Sizeof 功能类似
//     Size() uintptr

//     // 返回类型的字符串表示形式
//     String() string

//     // 返回类型的类型值
//     Kind() Kind

//     // 类型是否实现了接口 u
//     Implements(u Type) bool

//     // 是否可以赋值给 u
//     AssignableTo(u Type) bool

//     // 是否可以类型转换成 u
//     ConvertibleTo(u Type) bool

//     // 类型是否可以比较
//     Comparable() bool

//     // 下面这些函数只有特定类型可以调用
//     // 如：Key, Elem 两个方法就只能是 Map 类型才能调用

//     // 类型所占据的位数
//     Bits() int

//     // 返回通道的方向，只能是 chan 类型调用
//     ChanDir() ChanDir

//     // 返回类型是否是可变参数，只能是 func 类型调用
//     // 比如 t 是类型 func(x int, y ... float64)
//     // 那么 t.IsVariadic() == true
//     IsVariadic() bool

//     // 返回内部子元素类型，只能由类型 Array, Chan, Map, Ptr, or Slice 调用
//     Elem() Type

//     // 返回结构体类型的第 i 个字段，只能是结构体类型调用
//     // 如果 i 超过了总字段数，就会 panic
//     Field(i int) StructField

//     // 返回嵌套的结构体的字段
//     FieldByIndex(index []int) StructField

//     // 通过字段名称获取字段
//     FieldByName(name string) (StructField, bool)

//     // FieldByNameFunc returns the struct field with a name
//     // 返回名称符合 func 函数的字段
//     FieldByNameFunc(match func(string) bool) (StructField, bool)

//     // 获取函数类型的第 i 个参数的类型
//     In(i int) Type

//     // 返回 map 的 key 类型，只能由类型 map 调用
//     Key() Type

//     // 返回 Array 的长度，只能由类型 Array 调用
//     Len() int

//     // 返回类型字段的数量，只能由类型 Struct 调用
//     NumField() int

//     // 返回函数类型的输入参数个数
//     NumIn() int

//     // 返回函数类型的返回值个数
//     NumOut() int

//     // 返回函数类型的第 i 个值的类型
//     Out(i int) Type

//     // 返回类型结构体的相同部分
//     common() *rtype

//     // 返回类型结构体的不同部分
//     uncommon() *uncommonType
// }

import (
	"fmt"
	"reflect"
)

type T struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

func main() {
	t := reflect.TypeOf(T{})
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag
	// v的类型为string
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)      // 2 true
	fmt.Println(x.Get("max"))         // 99
	fmt.Println(x.Lookup("optional")) //  false
	fmt.Println(y.Lookup("optional")) // yes true
	fmt.Println(z.Lookup("optional")) // yes true
}
