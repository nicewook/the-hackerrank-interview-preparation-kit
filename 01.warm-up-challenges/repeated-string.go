package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
    firstCharCnt := make([]int64, len(s))
    var cnt int64
    for i, r := range s {
        if r == 'a' {
            cnt++
        }
        firstCharCnt[i] = cnt
    }

    totalCount := cnt * (n/int64(len(s)))
    rest := n%int64(len(s))
    if rest != 0 {
        totalCount += firstCharCnt[rest-1]
    }
    return totalCount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
