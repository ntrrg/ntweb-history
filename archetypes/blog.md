---
title: {{ replace .Name "-" " " | title }}
author: userid
publishdate: {{ .Date }}
date: {{ .Date }}
description: Short description.
image: https://via.placeholder.com/350x350
content:
  url: https://api.github.com/repos/user/project/contents/README.md
  api: github
tags:
  - tag1
series:
  - Serie 1
  - Serie 2
hiddenseries:
  - Serie 2
math: false
mermaid: false
comments: true
toc: true
draft: true
---

