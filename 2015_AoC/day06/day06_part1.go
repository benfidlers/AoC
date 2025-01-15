package main

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
  "strconv"
)

type pos struct {
    x int
    y int
}

const fileName = "input.txt"

func parse(text string) (string, pos, pos) {
  re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
  matches := re.FindStringSubmatch(text)

  action := matches[1]
  p1_x, _ := strconv.Atoi(matches[2])
  p1_y, _ := strconv.Atoi(matches[3])
  p1 := pos{x: p1_x, y: p1_y}
  p2_x, _ := strconv.Atoi(matches[4])
  p2_y, _ := strconv.Atoi(matches[5])
  p2 := pos{x: p2_x, y: p2_y}
  return action, p1, p2
}

func main() {
  file, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  scanner := bufio.NewScanner(file)

  lights := map[pos]bool {}
  for scanner.Scan() {
    action, p1, p2 := parse(scanner.Text())
    for i := p1.x; i <= p2.x; i++ {
      for j := p1.y; j <= p2.y; j++ {
        p := pos{x:i, y:j}
        if action == "turn on" {
          lights[p] = true 
        } else if action == "turn off" {
          lights[p] = false
        } else if action == "toggle" {
          lights[p] = !lights[p]
        }
      }
    }
  }
  count := 0
  for _, value := range lights {
    if value {
      count++
    }
  }
  fmt.Println(count)
}
