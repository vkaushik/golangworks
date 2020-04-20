package countingElements

func countElements(arr []int) int {
    count := 0
    register := [1001]int{}

    for _, element := range arr {
        register[element]++
    }

    for index := 0; index < 1000; index++ {
        if register[index + 1] != 0 {
            count = count + register[index]
        }
    }

    return count
}