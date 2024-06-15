package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
  goc()
}

func goc() {
  var status int = 1
  reader := bufio.NewReader(os.Stdin)

  os.Create("./history.c")
  os.Create("./current.c")

  for status != -1 {
    line, err := reader.ReadString('\n')
    if err ==  {return}
    if err != nil {
      fmt.Println("Unable to read line")
      continue
    }
    line = strings.TrimSuffix(line, "\n")
    fmt.Println(line)
  }
  os.Remove("./history.c")
  os.Remove("./current.c")
}
