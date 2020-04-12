package main

//import packages
import (
    "fmt"
    "os"
    "os/user"
    "log"
    "net/http"
)
//end import packages


func main() {


//Start Logging ability
  user, err := user.Current()
    if err != nil {
        panic(err)
    }

//Create the file to write the log out to. This will be where you run the executable
  f, err := os.OpenFile("WebAttack.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()

logger := log.New(f, "prefix", log.LstdFlags)
//End of logging ability - individual write to log lines are below


//Get the target from the user
  fmt.Print("Enter the domain to attack (example: https://www.domain.com):  ")
  var domain string
  fmt.Scanln(&domain)

  //log who is running the testing
  logger.Println("This test ran by Username: " + user.Username)
  logger.Println("Home Dir: " + user.HomeDir)
  logger.Println("********************************************")

//start url discovery array - These are the droids we are going to look for
dirsAndFiles := []string{"admin", "test", "testing", "demo", "dev", "development", "secure", "adm", "bak", "backup", "back", "logs", "log", "archive", "old", "~root", "inc", "js", "global", "local", "de", "en", "isapi", "wp-admin", "mail", "readme", "README", "ToDo", "todo", "changes", "Changes", "install.txt", "Install.txt", "config.inc", "install.inc", "password.inc", "connection.inc", "global.js", "local.js", "menu.js", "toolbar.js", "adovbs.inc", "database.inc", "db.inc", "sql", "SQL", "template", "templates", "temp", "tmp", "cgi-bin", "bin", "src", "robots.txt"}
//En of array

//Loop through the array
 for _, hackDirsandFiles := range  dirsAndFiles {
  result := domain + "/" + hackDirsandFiles

//Send the http request for the target URL and append the directory or file from the array above
  resp, err := http.Get(result)
  if err !=nil {
    log.Fatal(err)
  }

//Print to screen and write to log file
  fmt.Println("Testing", result)
  logger.Println("Testing", result)
  fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
  logger.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

  if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
      fmt.Println("HTTP Status is in the 2xx range")
      logger.Println("!!!!!!!!!!HTTP Status is in the 2xx range!!!!!!!!!!")
      fmt.Println("-------------------------------")
      logger.Println("-------------------------------")
      fmt.Println("                               ")
      logger.Println("                               ")

          } else {
      fmt.Println("Doesn't Seem to Exist!")
      logger.Println("Doesn't Seem to Exist!")
      fmt.Println("-------------------------------")
      logger.Println("-------------------------------")
      fmt.Println("                                  ")
      logger.Println("                               ")

  }
}
logger.Println("This is the end of the test ran by: " + user.Username)
logger.Println("Home Dir: " + user.HomeDir)
logger.Println("*****************************************************")
}
