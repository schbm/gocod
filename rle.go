package gocod

import (
	"bufio"
	"os"
)

const (
	CHUNK_SIZE = 8
)

func RleStream(in <-chan byte, out chan<- byte) error {
	var mask byte = 0b10000000

	var currPattern bool = false
	var occurences uint8 = 0
	for in := range in { //iterate over channel

		//shift the mask from 0 up to 7
		var tempMask uint8 = mask
		for maskShift := byte(0); maskShift < 8; maskShift++ {
			tempMask := tempMask >> maskShift       //1000 0000, 0100 0000
			var currBit bool = (in & tempMask) == 1 // get the current bit
			//check the current bit
			if currBit == currPattern { //if the bit is the same add an occurence
				//if chunk size -1 is reached write to the output
				occurences++
			} else {
				//write to out buffer
				//if buffer is full write to out
				//----------------------
				currPattern = currBit
				occurences = 1
			}
		}
	}
	return nil
}

func ReadFile(source string) error {

	file, err := os.Open(source)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	buff := make([]byte, 1024) //1024 bytes
	reader.Read(buff)

	return nil
}
