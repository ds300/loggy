package main

import irc "github.com/thoj/go-ircevent"
import "time"
import "os"
import "io/ioutil"
import "os/exec"
import s "strings"
import "math/rand"
import "github.com/BurntSushi/toml"
import "fmt"

type Config struct {
  Username string
  Password string
  Server string
  Channel string
  Push bool
}

/**
 * writes a line to the log, commits it
 */
func writeLine(msg string) {
  filename := time.Now().Format("2006-01-02") + ".log"

  f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    panic(err)
  }

  defer f.Close();

  if _, err := f.WriteString(msg + "\n"); err != nil {
    panic(err)
  }
  if err := exec.Command("git", "add", "-A").Run(); err != nil {
    panic(err)
  }
  if err := exec.Command("git", "commit", "-m", msg).Run(); err != nil {
    panic(err)
  }
}

/**
 * pushes with 'git push'
 */
func push() {
  if err := exec.Command("git", "push").Run(); err != nil {
    panic(err)
  }
}

func timestamp() string {
    return time.Now().Format("2006-01-02 15:04:05")
}


/**
 * turns a filename into the text contents of the file
 */
func slurp (file string) string {
  bytes, err := ioutil.ReadFile(file)
  if err != nil {
    panic(err)
  }
  return string(bytes)
}


func main() {
  // load config
  var configFile string
  var config Config

  if len(os.Args) == 1 {
    configFile = "loggy-conf.toml"
  } else {
    configFile = os.Args[1]
  }

  if _, err := toml.Decode(slurp(configFile), &config); err != nil {
    panic(err)
  }

  // connect to irc with user/pass
  con := irc.IRC(config.Username, config.Password)
  
  if err := con.Connect(config.Server); err != nil {
    panic(err)
  }

  // listen for incoming messages
  con.AddCallback("PRIVMSG", func(e *irc.Event) {
    // only log messages from the specified channel
    if e.Arguments[0] == config.Channel {
      msg := timestamp() + " " + e.Nick + ": " + e.Message()

      writeLine(msg)
      if config.Push {
        push()
      }

      // reply to things here
      if s.Contains(e.Message(), "@" + config.Username) {
        con.Privmsg(config.Channel, "Don't talk to me. I am a robot.")
      } else if s.Contains(e.Message(), config.Username) {
        con.Privmsg(config.Channel, "*cough*")
      }
    }
  })

  // listen for incoming peoples
  con.AddCallback("JOIN", func(e *irc.Event) {
    // perhaps greet them
    if rand.Int31n(5) == 0 {
      con.Privmsg(config.Channel, "Hello " + e.Nick + "!")
    }
  })

  // listen for irc connection resolution
  con.AddCallback("001", func(e *irc.Event) { con.Join(config.Channel) })

  // start connection loop
  con.Loop()
}
