package main

import "fmt"

func testMap(a map[string]map[string]string){
	_,ok := a["zhangsan"]
	if !ok{
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["passward"] = "1234566"
	a["zhangsan"]["nickname"] = "pangpang"
}
func testMap2()  {
	a:= make(map[string]map[string]string)
	a["key1"] = make(map[string]string)
	a["key1"]["1号"] = "zhangsan"
	a["key1"]["2号"] = "lisi"
	a["key2"] = make(map[string]string)
	a["key2"]["小马哥"] = "100000000"
	fmt.Println(a)
	for k,v := range a{
		fmt.Println(k)
		for k1,v1:=range v{
			fmt.Println(v[k1],v1)
			delete(v,k1)
		}
	}

	fmt.Println(a)
	delete(a,"key1")
	fmt.Println(a)

}
//map中key是无序的，对map排序必须对key进行排序，再遍历
 func main(){
	a := make(map[string]map[string]string)
	testMap(a)
	fmt.Println(a)
	fmt.Println(a["zhangsan"])
	a["lisi"] = make(map[string]string)
	a["lisi"]["passward"] = "2347646"
	fmt.Println(a)
	testMap2()

}
