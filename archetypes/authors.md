---
id: {{ replace .Name "-" " " }}
name: John Doe
position: CEO and Co-Founder
image: https://via.placeholder.com/350x350
social:
  website:
    label: Website
    text: example.com
    url: https://example.com
    icon: /images/social/fa-globe.svg
    weight: 1
  telegram:
    label: Telegram
    text: "@{{ replace .Name "-" " " }}"
    url: https://t.me/{{ replace .Name "-" " " }}
    icon: /images/social/fa-telegram.svg
    weight: 2
  facebook:
    label: Facebook
    text: John Doe
    url: https://www.facebook.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-facebook.svg
    weight: 3
  twitter:
    label: Twitter
    text: "@{{ replace .Name "-" " " }}"
    url: https://twitter.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-twitter.svg
    weight: 4
  instagram:
    label: Instagram
    text: "@{{ replace .Name "-" " " }}"
    url: https://instagram.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-instagram.svg
    weight: 5
  snapchat:
    label: Snapchat
    text: "@{{ replace .Name "-" " " }}"
    url: https://snapchat.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-snapchat.svg
    weight: 6
  linkedin:
    label: LinkedIn
    text: Jhon Doe
    url: https://www.linkedin.com/in/{{ replace .Name "-" " " }}
    icon: /images/social/fa-linkedin.svg
    weight: 7
  flickr:
    label: Flickr
    text: {{ replace .Name "-" " " }}
    url: https://flickr.com/people/{{ replace .Name "-" " " }}
    icon: /images/social/fa-flickr.svg
    weight: 8
  500px:
    label: 500px
    text: {{ replace .Name "-" " " }}
    url: https://500px.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-500px.svg
    weight: 9
  medium:
    label: Medium
    text: "@{{ replace .Name "-" " " }}"
    url: https://medium.com/@{{ replace .Name "-" " " }}
    icon: /images/social/fa-medium.svg
    weight: 10
  docker:
    label: Docker Hub
    text: "{{ replace .Name "-" " " }}"
    url: https://hub.docker.com/u/{{ replace .Name "-" " " }}
    icon: /images/social/fa-docker.svg
    weight: 11
  github:
    label: GitHub
    text: "@{{ replace .Name "-" " " }}"
    url: https://github.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-github.svg
    weight: 12
  gitlab:
    label: GitLab
    text: "@{{ replace .Name "-" " " }}"
    url: https://gitlab.com/{{ replace .Name "-" " " }}
    icon: /images/social/fa-gitlab.svg
    weight: 13
  bitbucket:
    label: Bitbucket
    text: "{{ replace .Name "-" " " }}"
    url: https://bitbucket.org/{{ replace .Name "-" " " }}
    icon: /images/social/fa-bitbucket.svg
    weight: 14
  email:
    label: Email
    text: {{ replace .Name "-" " " }}@example.com
    url: mailto://{{ replace .Name "-" " " }}@example.com
    icon: /images/social/fa-email.svg
    weight: 15
  phone-ve:
    label: Phone (Venezuela)
    text: +58 424-1234567
    url: tel://+584241234567
    icon: /images/social/fa-phone-square.svg
    weight: 16
  phone-us:
    label: Phone (US)
    text: +1 520-123-4567
    url: tel://+15201234567
    icon: /images/social/fa-phone-square.svg
    weight: 17
authortype: Person
---

