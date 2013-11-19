package contest

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	SMALL    = "small"
	LARGE    = "large"
	SAMPLE   = "sample"
	IN       = ".in"
	OUT      = ".out"
	SOLUTION = ".sol"
)

var (
	cin    chan string
	csin   chan string
	cout   chan string
	exit   chan int
	Small  bool
	Large  bool
	Sample bool
	std    bool
)

func init() {
	flag.BoolVar(&Small, "small", false, "small ")
	flag.BoolVar(&Large, "large", false, "large ")
	flag.BoolVar(&Sample, "sample", false, "sample ")
	flag.BoolVar(&std, "std", true, "std ")
}

func Test(solve func() string) {
	flag.Parse()
	cin = make(chan string)
	cout = make(chan string)
	exit = make(chan int)

	var fin, fout, fsol string

	switch {
	case Small == true:
		fin = SMALL + IN
		fout = SMALL + OUT
		fsol = SMALL + SOL
		break
	case Large == true:
		fin = LARGE + IN
		fout = LARGE + OUT
		fsol = LARGE + SOL
		break
	case Sample == true:
		fin = SAMPLE + IN
		fout = SAMPLE + OUT
		fsol = SAMPLE + SOL
		break
	case std == true:
		go ReopenStdin(os.Stdin)
		go ReopenStdout(os.Stdout)
		return
	default:
		panic("Select input!")
	}

	filein, err := os.Open(fin)
	if err != nil {
		panic(err)
	}

	fileout, err := os.Create(fout)
	if err != nil {
		panic(err)
	}

	go ReopenStdin(filein)
	go ReopenStdout(fileout)

	Run(solve)
}

func ReopenStdin(fin *os.File) {
	defer fin.Close()
	reader := bufio.NewReaderSize(fin, 1024*1024)

	var line []byte
	var err error
	for {
		line, _, err = reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				close(cin)
				return
			}
			fmt.Println(err)
		}
		cin <- string(line)
	}
}

func ReopenStdout(fout *os.File) {
	defer fout.Close()
	writer := bufio.NewWriterSize(fout, 1024)
	for {
		s, ok := <-cout
		if !ok {
			exit <- 1
			return
		}
		writer.WriteString(s)
		writer.Flush()
	}
}

func Run(solve func() string) {
	nTasks := ReadInt()
	for iTask := 0; iTask < nTasks; iTask++ {
		result := solve()
		Solve(iTask+1, result)
	}
	CloseStdout()
	Exit()
}

func Stoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Solve(i int, s string) {
	cout <- fmt.Sprintf("Case #%d: %s\n", i, s)
}

func CloseStdout() {
	close(cout)
}

func Exit() {
	<-exit
}

func ReadString() string {
	return <-cin
}

func ReadStringSlice() []string {
	return strings.Split(<-cin, " ")
}

func ReadInt() int {
	s, _ := <-cin
	i, _ := strconv.Atoi(s)
	return i
}

func ReadIntSlice() []int {
	args := strings.Split(<-cin, " ")
	out := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		n := Stoi(args[i])
		out[i] = n
	}
	return out
}

func ReadIntPair() (int, int) {
	args := strings.Split(<-cin, " ")
	return Stoi(args[0]), Stoi(args[1])
}
