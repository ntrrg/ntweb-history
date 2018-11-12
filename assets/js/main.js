window.addEventListener('load', function () {
  if (localStorage.getItem('dark-theme')) {
    setDarkTheme(true)
  }
})

function setDarkTheme(mode) {
  if (mode) {
    body.className = 'dark-theme'
    localStorage.setItem('dark-theme', true)
  } else {
    body.className = ''
    localStorage.removeItem('dark-theme')
  }
}

