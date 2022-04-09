package main
import (
    "fmt"
    "strings"
    "encoding/base64"
)
func question(question string) string {
    fmt.Println(question)
    input := "error"
    fmt.Scanln(&input)
    answer := strings.Replace(input, " ", "_", -1) // Failsafe for spaces, work in progess
	return answer
}
func final_number(user_name string, user_os string, user_de string) string {
    combined_with_space := user_name + user_os + user_de
    combined := strings.Replace(combined_with_space, " ", "_", -2)
    encodedString := base64.URLEncoding.EncodeToString([]byte(combined))
	fmt.Printf("Send this to @Jonte (jonatan@jontes.page or via Discord): %s\n", encodedString)
    return "success"
}
func main() {
    fmt.Println("Hello, I'm Jonte. I usually don't track people, like the way I have 0 cookies and trackers on my website, but today I decided to do the unthinkable...Collect data.")
	user_name := question("What is your nickname?")
    user_os := question("What Distro (or OS if you're not on Linux) do you use?")
    user_de := question("What is your DE/WM (explorer on Windows or aqua for MacOS")
	fmt.Println("Your name is",user_name,"and your OS of choice is",user_os,"and your DE/WM is",user_de)
    final_number(user_name, user_os, user_de)
}