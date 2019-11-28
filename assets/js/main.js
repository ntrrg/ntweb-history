// Theme

window.addEventListener('DOMContentLoaded', () => {
  const theme = localStorage.getItem('ntweb-theme')

  if (theme) {
    setTheme(theme)
  }

  const codeTheme = localStorage.getItem('ntweb-code-theme')

  if (codeTheme) {
    setCodeTheme(codeTheme)
  }
})

function setTheme(theme) {
  document.body.dataset.theme = theme
  localStorage.setItem('ntweb-theme', theme)
}

function setCodeTheme(theme) {
  document.body.dataset.codeTheme = theme
  localStorage.setItem('ntweb-code-theme', theme)
}

// Image lazy loading

window.addEventListener('DOMContentLoaded', () => {
  let images = document.querySelectorAll('img.lazy-load')

  for (let img of images) {
    img.src = img.dataset.src 
  }
})

