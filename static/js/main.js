window.addEventListener('load', () => {
  const body = document.querySelector('#body')
  const toggler = document.querySelector('#theme-toggler')
  
  toggler.addEventListener('change', () => {
    if (toggler.checked) {
      body.className = 'dark-theme'
      localStorage.setItem('dark-theme', true)
    } else {
      body.className = ''
      localStorage.removeItem('dark-theme')
    }
  })

  if (localStorage.getItem('dark-theme')) {
    toggler.dispatchEvent(new MouseEvent('click'))
  }
})

