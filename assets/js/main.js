// Theme

window.addEventListener('load', () => {
  const theme = localStorage.getItem('theme')

  if (theme) {
    setTheme(theme)
  }

  const codeTheme = localStorage.getItem('code-theme')

  if (codeTheme) {
    setCodeTheme(codeTheme)
  }
})

function setTheme(theme) {
  document.body.dataset.theme = theme
  localStorage.setItem('theme', theme)
}

function setCodeTheme(theme) {
  document.body.dataset.codeTheme = theme
  localStorage.setItem('code-theme', theme)
}

// Image lazy loading

/*
window.addEventListener('load', () => {
  let images = document.querySelectorAll('img.lazy-load')

  for (let img of images) {
    img.src = img.dataset.src 
  }
})
*/

