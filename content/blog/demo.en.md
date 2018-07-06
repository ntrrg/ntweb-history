---
title: Demo page
date: 2028-07-05T18:35:00-04:00
image: /uploads/gopher.png
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

# Headers

# h1 Heading
## h2 Heading
### h3 Heading
#### h4 Heading
##### h5 Heading
###### h6 Heading


---

**This is bold text**

__This is bold text__

*This is italic text*

_This is italic text_

~~Deleted text~~

This is text with inline math \\(\sum\_{n=1}^{\infty} 2^{-n} = 1\\) and with math blocks:

$$
\sum\_{n=1}^{\infty} 2^{-n} = 1
$$

> Block quotes are
> written like so.
>
> They can span multiple paragraphs,
> if you like.

Some text, and some `code` and then a nice plain [link with title](https://github.com/davidhampgonsalves/davidhampgonsalves.com-hugo "title text!").

and then

+ Create a list by starting a line with `+`, `-`, or `*`
+ Sub-lists are made by indenting 2 spaces:
  - Marker character change forces new list start:
    * Ac tristique libero volutpat at
+ Very easy!

vs.

1. Lorem ipsum dolor sit amet
2. Consectetur adipiscing elit
3. Integer molestie lorem at massa

## Code

Inline `code`

```go
package main

import "fmt"

func main() {
  fmt.Println("hello, world")
}
```

<kbd>Ctrl</kbd> + <kbd>Alt</kbd> + <kbd>Del</kbd>

| Heading | Another heading |
| ------  | --------------- |
|  text   |      text       |
|  text   |      text       |
|  text   |      text       |

| Heading | Another heading |
| :----:  | :-------------: |
|  text   |      text       |
|  text   |      text       |
|  text   |      text       |

| Heading | Another heading |
| -----:  | --------------: |
|  text   |      text       |
|  text   |      text       |
|  text   |      text       |

