package main

import "fmt"

func buble_sort(a []int){
	for i:=len(a);i>0;i--{
		for j:=0;j<i-1;j++{
			if(a[j]>a[j+1]){
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}
func insert_sort(a []int){
	for i:=1; i<len(a); i++{
		tmp:=a[i]
		j := i-1
		for tmp < a[j]&&j>=0{
			a[j+1]=a[j]
			if(j!=0) {
				j--
			}else {
				break
			}
		}
		//if j {
		//
		//}
		if(tmp>a[j]) {
			a[j+1] = tmp
		}else{
			a[j] = tmp
		}

	}
}
func main(){
	b := [...]int{5,7,4,3,9,8,6}
	//buble_sort(b[:])
	insert_sort(b[:])
	fmt.Println(b)
}
