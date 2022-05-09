package main

import (
  "encoding/json"
  "fmt"
  "syscall/js"
)

type Link struct {
  CharacterStartIndex uint   `json:"characterStartIndex"`
  CharacterEndIndex   uint   `json:"characterEndIndex"`
  Href                string `json:"href"`
}

type Page struct {
  Content string `json:"content"`
  Links   []Link `json:"links"`
}

func Example() js.Func {
  example := Page {
    Content: "Hello, Shopify!",
    Links: []Link{
      {
        CharacterStartIndex: 0,
        CharacterEndIndex: 14,
        Href: "https://shopify.com",
      },
    },
  }

  return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    encoded, _ := json.Marshal(example)
    return string(encoded)
  })
}

func render(page Page) string {
  // Assume one link because multiple is complicated lol
  link := page.Links[0]
  pre := page.Content[:link.CharacterStartIndex]
  post := page.Content[link.CharacterEndIndex:]
  between := page.Content[link.CharacterStartIndex:link.CharacterEndIndex]

  return fmt.Sprintf("%s<a href=\"%s\">%s</a>%s", pre, link.Href, between, post)
}

func Render() js.Func {
  return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    if len(args) != 1 {
      return "Missing page argument"
    }

    page := Page{}

    err := json.Unmarshal([]byte(args[0].String()), &page)
    if err != nil {
      return "Invalid page argument"
    }

    return render(page)
  })
}

func main() {
  js.Global().Set("Example", Example())
  js.Global().Set("Render", Render())
  <-make(chan bool)
}
