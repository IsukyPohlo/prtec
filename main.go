package main

import (
      "fmt"
      "strings"
)

func main() {

      out1 := Encrypt("dcj", " I love prOgrAmming!")

      fmt.Println(out1) // dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!

      out2 := ZeroSum([]int{3, 4, -7, 5, -6, 2, 5, -1, 8}) // 5 , 8

      //out := ZeroSum([]int{1, 2, -3, 3, 1}) // 3, 1

      //out := ZeroSum([]int{1, 2, 3, -3, 4}) // 1, 2, 4

      fmt.Println(out2)
}

func Encrypt(key, message string) string {

      if message == "" {
            return ""
      }

      if key == "" {
            key = "dcj"
      }

      vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

      for _, y := range vowels {
            message = strings.Replace(message, string(y), key+string(y), -1)
      }

      return message
}

func ZeroSum(input []int) []int {

      if len(input) == 0 {
            return nil
      }

      // add element at beginning
      input = append([]int{0}, input...)

      fmt.Println(input)

      sumMap := map[int]int{}
      sum := 0
      output := []int{}
      lastCut := 0

      for idx, val := range input {

            sum = sum + val

            if _, ok := sumMap[sum]; !ok {

                  sumMap[sum] = idx
                  continue
            }

            val := sumMap[sum]

            // save new segment
            newSeg := input[lastCut : val+1]
            output = append(output, newSeg...)

            // update lastCut
            lastCut = idx + 1

            // reset sum map
            sumMap = map[int]int{}

      }

      // save from lastcut to end
      lastSegment := input[lastCut:]

      // remove added element and concat
      output = append(output[1:], lastSegment...)

      return output
}
