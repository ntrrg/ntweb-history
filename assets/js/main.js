window.addEventListener('load', function () {
  var body = document.querySelector('#body')
  var toggler = document.querySelector('#theme-toggler')

  toggler.addEventListener('change', function () {
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

