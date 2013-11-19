package main

// import (
// 	. "bufio"
// 	"io"
// 	"os"
// 	. "strconv"
// )

// type StringReader struct {
// 	content string
// 	offset  int
// }

// func (s *StringReader) Close() error { return nil }

// func (s *StringReader) Read(p []byte) (int, error) {
// 	n := copy(p, s.content[s.offset:])

// 	if n == 0 {
// 		return 0, io.EOF
// 	}

// 	s.offset += n
// 	return n, nil
// }

// func NewStringReader(content string) *StringReader {
// 	return &StringReader{content, 0}
// }

// func RunDefault(solve func(nextLine func() string) string) {
// 	file, _ := os.Open(os.Args[1])
// 	defer file.Close()
// 	Run(file, os.Stdout, solve)
// }

// func Run(input io.ReadCloser, output io.WriteCloser, solve func(nextLine func() string) string) {
// 	reader := NewReaderSize(input, 1024*1024)
// 	writer := NewWriterSize(output, 1024)
// 	line, _, _ := reader.ReadLine()

// 	nTasks, _ := Atoi(string(line))

// 	for iTask := 0; iTask < nTasks; iTask++ {
// 		result := solve(func() string {
// 			line, _, _ := reader.ReadLine()
// 			return string(line)
// 		})

// 		writer.WriteString("Case #" + Itoa(iTask+1) + ": " + result + "\n")
// 		writer.Flush()
// 	}
// }

// func Stoi(s string) int {
// 	i, _ := Atoi(s)
// 	return i
// }
