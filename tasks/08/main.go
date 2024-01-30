package main

import (
	"fmt"
)

func ChangeBit(num, newBit int64, pos uint) int64 {
	mask := int64(^(1 << pos)) //создание маски с битовым сдвигом и инвертацией числа (^010=101)
	num &= mask                //"очищаем бит"
	num |= newBit << pos       //устанваливаем новый бит в позицию
	return num
}

func main() {
	fmt.Println(ChangeBit(5, 0, 0))
	fmt.Println(ChangeBit(8, 1, 0))
}
