package main

import (
  "fmt"
  "time"
  "github.com/tebeka/selenium"
)

var code = `
package main
import "fmt"

func main() {
    fmt.Println("Hello WebDriver!\n")
}
`

func main() {
    // Connect to the selenium server
    caps := selenium.Capabilities{"browserName": "firefox"}
    wd, err := selenium.NewRemote(caps, "")
    if err != nil {
        fmt.Println(err)
    }
    defer wd.Quit()

    // Get simple playground interface
    wd.Get("http://play.golang.org/?simple=1")

    // Enter code in textarea
    elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
    elem.Clear()
    elem.SendKeys(code)

    // Click the run button
    btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
    btn.Click()

    // Get the result
    div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")

    output := ""
    // Wait for run to finish
    for {
        output, _ = div.Text()
        if output != "Waiting for remote server..." {
            break
        }
        time.Sleep(time.Millisecond * 100)
    }

    fmt.Printf("Got: %s\n", output)
}
