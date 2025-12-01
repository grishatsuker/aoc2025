package main

import (
	"strconv"
    "fmt"
    "os"
    "bufio"
)

func sum_times_safe_hits_zero(inputs []string) int{
	var res int = 0
	var current int = 50
	for i := 0; i < len(inputs); i++ {
		rotation := inputs[i]
		direction := rotation[:1]
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil{
			fmt.Println("Error, skipping?")
			continue
		}
		over_zero := clicks / 100
		res += over_zero

		clicks_after := clicks - (over_zero * 100)

		switch direction{
			case "R":
				current += clicks_after
				if current > 99 {
					res += 1
					current -= 100
				}
			case "L":
				if current == 0 {
					res -= 1
				}
				current -= clicks_after
				if current == 0 {
					res += 1
				}
				if current < 0 {
					res +=1
					current += 100
				}
			default:
				fmt.Println("Something bad happened")

		}

		fmt.Printf("current: %d, res: %d, rotation: %s, direction: %s, clicks: %d over zero: %d, clicks_after: %d\n", current, res, rotation, direction, clicks, over_zero, clicks_after)

	}
	return res
}

func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		fmt.Println("Opening the file failed")
		return
	}
	defer file.Close()

	var lines []string
 	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
        return
	}
	result := sum_times_safe_hits_zero(lines)
	fmt.Printf("Result is: %d", result)

}
