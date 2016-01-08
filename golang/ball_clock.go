package main

import (
	"bufio"
    "fmt"
	"log"
	"io"
    "os"
	"path/filepath"
	"strconv"
)

const (	MAKE_CLOCK=iota
		GET_BALL
		MINUTE
		EMPTY_MINUTE
		FIVE_MINUTE
		EMPTY_FIVE_MINUTE
		HOUR
		EMPTY_HOUR
		PRINT_DAYS )

func main() {
	slNumberOfBalls := make([]int, 0)	// each input line becomes one element
	numberOfBalls,err := ReadBallClock(slNumberOfBalls)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	slNumberOfBalls=numberOfBalls
	var nBalls int = 0					// holds one slNumberOfBalls element

	slBallClock := make([]int, 0)		// clock
	slBall := make([]int, 0)			// ball in motion
	slMinute := make([]int, 0)			// minute storage
	slFiveMinute := make([]int, 0)		// five minute storage
	slHour:= make([]int, 0)				// hour storage

	var nBallState int = MAKE_CLOCK		// beginning program state
	var nIterations int = 0				// count of completed clock cycles

	//var nLoops int = 0

	DONE:
	for {
		switch nBallState {
			case 0: //MAKE_CLOCK
				nIterations=0								// back to zero
				if len(slNumberOfBalls) > 0 {
					if len(slBallClock) > 0	{				// remove previous clock
						slBallClock = slBallClock[:0]
					}
					nBalls=slNumberOfBalls[0]
					if nBalls == 0 {
						break DONE
					}
					slNumberOfBalls = slNumberOfBalls[1:]
					for i:=0; i<nBalls; i++ {
						slBallClock = append(slBallClock, i+1)
					}
					nBallState = GET_BALL
				} else {
					break DONE
				}
			case 1:	//GET_BALL
				ballClock,ball := GetBall(&nBallState,slBallClock,slBall)
				slBallClock=ballClock
				slBall=ball
			case 2:  //MINUTE
				ball,minute := Minute(&nBallState,slBall,slMinute)
				slBall=ball
				slMinute=minute
			case 3: //EMPTY_MINUTE
				ballClock,minute := EmptyMinute(&nBallState,slBallClock,slMinute)
				slBallClock=ballClock
				slMinute=minute
			case 4: //FIVE_MINUTE
				ball,fiveMinute := FiveMinute(&nBallState,slBall,slFiveMinute)
				slBall=ball
				slFiveMinute=fiveMinute
			case 5: //EMPTY_FIVE_MINUTE 
				ballClock,fiveMinute := EmptyFiveMinute(&nBallState,slBallClock,slFiveMinute)
				slBallClock=ballClock
				slFiveMinute=fiveMinute
			case 6: //HOUR
				ball,hour := Hour(&nBallState,slBall,slHour)
				slBall=ball
				slHour=hour
			case 7:	//EMPTY_HOUR
				ballClock,hour,ball := EmptyHour(&nBallState,&nIterations,slBallClock,slHour,slBall)
				slBallClock=ballClock
				slHour=hour
				slBall=ball
			case 8: //PRINT_DAYS
				fmt.Println(nBalls,"balls cycle after",(nIterations+1)/2,"days.")
				nBallState = MAKE_CLOCK
		}
	}
}

func GetBall(nBallState* int,slBallClock []int,slBall []int) ([]int,[]int) {
	if len(slBall) > 0 {
		slBall = slBall[1:]
	}
	slBall = append(slBall,slBallClock[0])
	slBallClock = slBallClock[1:]
	*nBallState=MINUTE
	return slBallClock,slBall
}

func Minute(nBallState* int,slBall []int,slMinute []int) ([]int,[]int) {
	if len(slMinute) < 4 {
		slMinute = append(slMinute, slBall[0])
		*nBallState = GET_BALL
	} else {
		*nBallState = EMPTY_MINUTE
	}
	return slBall,slMinute
}

func EmptyMinute(nBallState* int,slBallClock []int,slMinute[]int) ([]int,[]int) {
	var nMin int = len(slMinute)
	if nMin > 0 {
		slBallClock = append(slBallClock,slMinute[nMin-1])
		slMinute = slMinute[:nMin-1]
	} else {
		*nBallState = FIVE_MINUTE
	}
	return slBallClock,slMinute
}

func FiveMinute(nBallState* int,slBall []int,slFiveMinute []int) ([]int,[]int) {
	if len(slFiveMinute) < 11 {
		slFiveMinute = append(slFiveMinute, slBall[0])
		*nBallState = GET_BALL
	} else {
		*nBallState = EMPTY_FIVE_MINUTE
	}
	return slBall,slFiveMinute
}

func EmptyFiveMinute(nBallState* int,slBallClock []int,slFiveMinute[]int) ([]int,[]int)  {
	var nMin int = len(slFiveMinute)
	if nMin > 0 {
		slBallClock = append(slBallClock,slFiveMinute[nMin-1])
		slFiveMinute = slFiveMinute[:nMin-1]
	} else {
		*nBallState = HOUR
	}
	return slBallClock,slFiveMinute
}

func Hour(nBallState* int,slBall []int,slHour []int) ([]int,[]int) {
	if len(slHour) < 11 {
		slHour = append(slHour, slBall[0])
		*nBallState = GET_BALL
	} else {
		*nBallState = EMPTY_HOUR
	}
	return slBall,slHour
}

func EmptyHour(nBallState* int,nIterations* int,slBallClock []int,slHour[]int,slBall []int) ([]int,[]int,[]int) {
	var nHour int = len(slHour)
	if nHour > 0 {
		slBallClock = append(slBallClock,slHour[nHour-1])
		slHour = slHour[:nHour-1]
	} else {
		slBallClock = append(slBallClock, slBall[0])
		if  OrderTest(slBallClock) {
			*nBallState = PRINT_DAYS
		} else {
			*nIterations = *nIterations + 1
			*nBallState = GET_BALL
		}
	}
	return slBallClock,slHour,slBall
}

func OrderTest(slBallClock []int)(bool) {
	for i := range slBallClock {
		if slBallClock[i] != i+1 {
			return false
		}
	}
	return true
}

func ReadBallClock(slNumberOfBalls []int) ([]int,error) {
	inFilename,err := fileNameFromCommandLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile := os.Stdin
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}
	reader := bufio.NewReader(inFile)

	eof := false

	for !eof {
		var line string
		line, err = reader.ReadString('\n')

		if err == io.EOF {
			err = nil // io.EOF isn't really an error
			eof = true // this will end the loop at the next iteration
		} else if err != nil {
			fmt.Println("Error")
			return slNumberOfBalls,err // finish immediately for real errors
		}

		var nSize int
		nSize = len(line)

		if nSize > 0 {
			line = line[:(len(line)-2)]
			nNumBalls,err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			if (nNumBalls < 27 || nNumBalls > 127) && (nNumBalls != 0) {	// check input
				fmt.Println("Error : ",nNumBalls," is invalid. Number must be greater than 26 or less than 128.")
				os.Exit(2)
			}
			slNumberOfBalls = append(slNumberOfBalls, nNumBalls)
		}
	}
	return slNumberOfBalls,err
}

func fileNameFromCommandLine() (inFilename string, err error) {
	if len(os.Args) == 1 {
		err = fmt.Errorf("usage: %s [<]infile.txt",filepath.Base(os.Args[0]))
		return "", err
	} else if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt",filepath.Base(os.Args[0]))
		return "", err
	}
	if len(os.Args) > 1 {
		inFilename = os.Args[1]
	}
	return inFilename, nil
}