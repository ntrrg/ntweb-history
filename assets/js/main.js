// Theme

window.addEventListener('DOMContentLoaded', () => {
  const theme = localStorage.getItem('general-theme')

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
  localStorage.setItem('general-theme', theme)
}

function setCodeTheme(theme) {
  document.body.dataset.codeTheme = theme
  localStorage.setItem('code-theme', theme)
}

// Cookies

function checkCookiesAlert() {
  const cookiesAlert = localStorage.getItem('cookies-alert')

  if (cookiesAlert === 'agree') {
    const el = document.querySelector('#cookies')
    el.parentNode.removeChild(el)
  }
}

if (document.querySelector('#cookies') !== null) {
  document.querySelector('#cookies').addEventListener('click', (e) => {
    localStorage.setItem('cookies-alert', 'agree')
    checkCookiesAlert()
  })

  window.addEventListener('DOMContentLoaded', checkCookiesAlert)
}

// Image lazy loading

window.addEventListener('DOMContentLoaded', () => {
  const images = document.querySelectorAll('img.lazy-load')

  for (const img of images) {
    img.src = img.dataset.src 
  }
})

