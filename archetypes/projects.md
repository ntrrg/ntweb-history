---
title: {{ replace .Name "-" " " | title }}
publishdate: {{ .Date }}
date: {{ .Date }}
description: Short description.
image: https://via.placeholder.com/350x350
metadata:
  client: "[Example Inc.](https://example.com)"
  role: Software Developer
  source-code: https://github.com/user/project
  license: MIT
content:
  url: https://api.github.com/repos/user/project/contents/README.md
  api: github
tags:
  - tag1
math: false
mermaid: false
comments: true
toc: true
draft: true
---

