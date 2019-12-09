---
title: {{ replace .Name "-" " " | title }}
author: userid
date: {{ .Date }}
description: Short description.
image: https://via.placeholder.com/350x350
type: oembed | script
ref: https://via.placeholder.com
oembed: https://example.com/oembed?format=json&url=%s
script: <script async src="https://example.com/widget.js" data-url="%s"></script>
tags:
  - tag1
comments: true
draft: true
---

