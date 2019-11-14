// Theme

window.addEventListener('DOMContentLoaded', () => {
  const theme = sessionStorage.getItem('theme')

  if (theme) {
    setTheme(theme)
  }

  const codeTheme = sessionStorage.getItem('code-theme')

  if (codeTheme) {
    setCodeTheme(codeTheme)
  }
})

function setTheme(theme) {
  document.body.dataset.theme = theme
  sessionStorage.setItem('theme', theme)
}

function setCodeTheme(theme) {
  document.body.dataset.codeTheme = theme
  sessionStorage.setItem('code-theme', theme)
}

// Image lazy loading

window.addEventListener('DOMContentLoaded', () => {
  let images = document.querySelectorAll('img.lazy-load')

  for (let img of images) {
    img.src = img.dataset.src 
  }
})

