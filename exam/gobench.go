package exam

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
)

var (
	// Benchmarks setting whether to run the benchmarks or just the tests
	Benchmarks bool
)

// GoBench runs `go test` on provided path and adds suggestions to the response
func GoBench(_ *astrav.Folder, r *extypes.Response, _ string, skip bool) bool {
	if !Benchmarks || skip {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Brown("SKIPPED"))
		return true
	}

	res, state := bench()

	if state.Success() {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Green("OK"))
	} else {
		fmt.Println(aurora.Gray("benchmarks:\t"), aurora.Red("FAIL"))
	}

	fmt.Println(res)

	if state.Success() {
		return true
	}
	r.AppendTodoTpl(gtpl.PassTests)
	return false
}

func bench() (string, *os.ProcessState) {
	cmd := exec.Command("go", "test", "-v", "--bench", ".", "--benchmem", "-test.run=^$")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go test: ", err)
	}

	return string(b), cmd.ProcessState
}
