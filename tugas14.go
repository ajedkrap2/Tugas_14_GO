package main

import "fmt"
import "os/exec"

func main() {
    var listFiles,_ = exec.Command("ls").Output()
    var path,_ = exec.Command("pwd").Output()
    var network,_ = exec.Command("ifconfig").Output()
    fmt.Println(string(listFiles), "\n\n")
    fmt.Println(string(path), "\n\n")
    fmt.Println(string(network), "\n\n")
}
