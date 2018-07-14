---
title: Demo page
date: 2028-07-05T18:35:00-04:00
image: /images/logo.png
description: This is a demo page to see the styles from NtWeb.
categories:
  - demo
tags:
  - tag1
  - tag2
  - tag3
  - tag4
  - tag5
math: true
---

{{< toc >}}

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

# Images

![test image](/images/logo.png)

# Lists

* Create a list by starting a line with `+`, `-`, or `*`
* Sub-lists are made by indenting 2 spaces:
  * This is a sublist
* And everything become normal again

1. This is
2. an ordered
3. list

This
: is a definition list.

Term:
: definition, you can add the `:` in the term.

# Quotes

<!--lint disable no-undefined-references no-shortcut-reference-link-->
This paragraph has a footnote[^1].

[^1]: And here is the footnote.
<!--lint enable no-undefined-references no-shortcut-reference-link-->

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

# Math formulas

This is a smart fraction 1/2, this is text with inline math
\\(\sum\_{n=1}^{\infty} 2^{-n} = 1\\) and this is a math blocks:

$$
\sum\_{n=1}^{\infty} 2^{-n} = 1
$$

# Code

Inline `code`.

```go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```

# Theme

## Keyboard

<kbd>Ctrl</kbd> + <kbd>Alt</kbd> + <kbd>Del</kbd>

## Notes

{{% note %}}
This is a note.
{{% /note %}}

{{% note "My title" %}}
This is a note with a custom title.
{{% /note %}}

## Links of interest

{{% loi %}}
* <https://nt.web.ve>
* <https://nt.web.ve/en>
* <https://nt.web.ve/es>
{{% /loi %}}

## Go Playground

{{% go-playground "sfPZJf_jZLF" %}}
```go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```
{{% /go-playground %}}

