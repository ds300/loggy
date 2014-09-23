/* The MIT License (MIT)

Copyright (c) 2014 Everyone

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. */

package main

import irc "github.com/thoj/go-ircevent"
import "time"
import "os"
import "io/ioutil"
import "os/exec"
import "github.com/BurntSushi/toml"

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
    }
  })

  // listen for irc connection resolution
  con.AddCallback("001", func(e *irc.Event) { con.Join(config.Channel) })

  // start connection loop
  con.Loop()
}
