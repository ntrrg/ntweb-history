---
title: Demo page
author: ntrrg
publishdate: 2028-07-05T18:35:00-04:00
date: 2028-07-05T18:35:00-04:00
image: images/ntrrg.png
description: This is a demo page to see the Markdown styles.
tags:
  - tag1
  - tag2
  - tag3
  - tag4
  - tag5
math: true
mermaid: true
toc: true
draft: true
---

# Paragraphs

Any line with blank lines before and after it is a paragrah,
consequent lines are joined.

You need a blank line for a new paragraph.

# Separators

---

# Heading (h1) 

## Heading (h2) 

### Heading (h3) 

#### Heading (h4) 

##### Heading (h5) 

###### Heading (h6) 

# Text decoration

**This is bold text**

__This is bold text__

*This is italic text*

_This is italic text_

~~This is dashed text~~

<https://nt.web.ve>

[This is a link](https://nt.web.ve)

[This is a link with a title](https://nt.web.ve "This is the title!").

[ntweb]: https://nt.web.ve

[ntweb][]

[My site][ntweb]

# Lists

* Create a list by starting a line with `+`, `-`, or `*`
* Sub-lists are made by indenting 2 spaces:
  * This is a sublist
* And everything become normal again

1. This is
2. an ordered
3. list

* [ ] This is
* [x] a task list

This
: is a definition list.

Term:
: definition, you can add the `:` in the term.

# Quotes

This paragraph has a footnote[^1].

[^1]: And here is the footnote.

> Block quotes are
> written like so.
>
> They can span multiple paragraphs, if you like.
>
> And **Markdown**!.
>
> -- The Author

# Tables

| Heading | Another heading |
| ------- | --------------- |
| text    | text            |
| text    | text            |
| text    | text            |

| Heading | Another heading |
| :-----: | :-------------: |
|  text   |      text       |
|  text   |      text       |
|  text   |      text       |

| Heading | Another heading |
| ------: | --------------: |
|    text |            text |
|    text |            text |
|    text |            text |

# Images

!["Alternative text"](images/ntrrg.png "This is the title")

# Math formulas

This is a smart fraction 1/2, this is text with inline math
\\(\sum\_{n=1}^{\infty} 2^{-n} = 1\\) and this is a math block:

$$
\sum\_{n=1}^{\infty} 2^{-n} = 1
$$

# Code

Inline `code`.

```go {linenos=true,hl_lines=["1", "5-7"],linenostart=0}
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```

# Shortcodes

## Keyboard

{{< kbd "Ctrl" >}} + {{< kbd "Alt" >}} + {{< kbd "Del" >}}

## Images

### Inline

Inline image with {{< img src="images/hugo.png" style="height: 1em;" >}}

### Wide

{{< img src="images/merida.jpg" class="wide" >}}

### Block (center aligned)

{{< img src="images/merida.jpg" class="block" >}}

### Left aligned

{{< img src="images/ntrrg.png" class="align-left" >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

### Right aligned

{{< img src="images/ntrrg.png" class="align-right" >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Figures

### Wide

{{< figure src="images/merida.jpg" class="wide" title="Title" caption="Figure caption. Supports **Markdown**." >}}

### Block (center aligned)

{{< figure src="images/merida.jpg" class="block" title="Title" caption="Figure caption. Supports **Markdown**." >}}

### Left aligned

{{< figure src="images/ntrrg.png" class="align-left" title="Title" caption="Figure caption. Supports **Markdown**." >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

### Right aligned

{{< figure src="images/ntrrg.png" class="align-right" title="Title" caption="Figure caption. Supports **Markdown**." >}}

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Go Playground

{{< go-playground >}}
```go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```
{{< /go-playground >}}

## Notes

{{< note >}}
This is a note.
{{< /note >}}

{{< note "My title" >}}
This is a **note** with a custom title.
{{< /note >}}

## Links of interest

{{< loi >}}
* <https://nt.web.ve>
* <https://nt.web.ve/en/>
* <https://nt.web.ve/es/>
{{< /loi >}}

## Details

{{< details >}}
**Lorem ipsum** dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{< /details >}}

{{< details summary="Custom text" open=true >}}
**Lorem ipsum** dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{< /details >}}

# Charts

{{< mermaid "Chart caption. Supports **Markdown**." >}}
```mermaid
graph TD
  A[Christmas] -->|Get money| B(Go shopping)
  B --> C{Let me think}
  C -->|One| D[Laptop]
  C -->|Two| E[Phone]
  C -->|Three| F[fa:fa-car Car]
```
{{< /mermaid >}}

## Snippets

{{< snippet "files/hello.go" "go" >}}

{{< snippet "files/hello.go" "go {linenos=true,hl_lines=[\"1\", \"5-7\"],linenostart=0}" >}}

{{< import "files/imports.en.md" >}}

## Cards

{{< card "/blog/demo/" >}}

