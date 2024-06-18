package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "io"
)

func main() {
  goc()
}

func goc() { 
  var status int = 0
  var files = []string{"./current.c", "./history.c"}
  source_files := createFiles(files)
  source_current, source_history := source_files[0], source_files[1]
  reader := bufio.NewReader(os.Stdin)
  for status != 1 {
    fmt.Printf(">> "); line, err := reader.ReadString('\n')
    if err == nil {
      line = strings.TrimSuffix(line, "\n")
      addLine(source_current, source_history, line)
      status = 0
    } else {
      fmt.Printf("\n%s, exiting...\n", err)
      status = 1
    }
  }
  removeFiles(files)
}

func addLine(current, history io.Writer, line string) {
  fmt.Fprintln(current, line)
  fmt.Fprintln(history, line)
}

func removeFiles(files []string) {
  for _, file := range files {
    os.Remove(file)
  } 
}

func createFiles(files []string) (source_files []*os.File){
  for _, file := range files {
    source_file, _ := os.Create(file)
    source_files = append(source_files, source_file)
  }
  return
}
