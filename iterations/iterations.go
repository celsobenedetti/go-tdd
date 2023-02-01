package iterations

func Repeat(str string, iterations int) (result string) {
    for i := 0; i < iterations; i++ {
        result += str
    }
    return
}
