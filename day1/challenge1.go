package challenge1

import(
	"fmt"
)

func RunChallenge1(arr := [1998]int64){
	for (i:=0;i<1999;i++){
		if(arr[i]>arr[i-1]){
			fmt.Println(arr[i]+" (increased)");
			counter := counter+1;
		}else{
			fmt.Println(arr[i]+" (decreased)");
		}
	}
	fmt.Printf("%d", counter);
}
