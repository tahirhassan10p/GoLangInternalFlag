package main
/*
#include <stdio.h>
#include <stdlib.h>
#include "header.h"
#cgo CFLAGS: -I./ 
#cgo LDFLAGS: ./libarith.a

void myprint() {
    int result;
	result = addition(1,2);
	printf("addition result is : %d\n",result);
	result = multiplication(3,2);
	printf("multiplication result is :  %d\n",result);
}
*/
import "C"

import "fmt"

func main() {
	fmt.Println("Hello world from archive")
	C.myprint()
}