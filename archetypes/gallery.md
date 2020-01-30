---
title: {{ replace .Name "-" " " | title }}
author: userid
date: {{ .Date }}
description: Short description.
image: https://via.placeholder.com/350x350
mode: local | oembed | script
images:
  - src: https://via.placeholder.com/500x500
  - src: https://via.placeholder.com/350x350
    caption: This is a `350x350px` image.
  - src: https://via.placeholder.com/1280x1024
    ref: https://via.placeholder.com
    caption: This is a `1280x1024px` image.
ref: https://via.placeholder.com
oembed: https://example.com/oembed?format=json&url=%s
script: <script async src="https://example.com/widget.js" data-url="%s"></script>
tags:
  - tag1
comments: true
draft: true
---

