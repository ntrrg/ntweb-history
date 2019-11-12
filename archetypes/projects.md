---
title: {{ replace .Name "-" " " | title }}
publishdate: {{ .Date }}
date: {{ .Date }}
description: Short description.
image: https://via.placeholder.com/350x350
metadata:
  source-code: https://github.com/user/project
  license: MIT
content:
  url: https://api.github.com/repos/user/project/contents/README.md
  api: github
tags:
  - tag1
comments: true
draft: true
---

