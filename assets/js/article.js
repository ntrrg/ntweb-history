// Wide images

function resizeWideImages() {
  let leftMargin = '-1em'
  const bodyWidth = document.querySelector('#body').offsetWidth
  const mainWidth = document.querySelector('#body > main').offsetWidth

  if (bodyWidth !== mainWidth) {
    leftMargin = 'calc(' + (bodyWidth - mainWidth) / 2 * -1 + 'px - 1em)'
  } 
  const images = document.querySelectorAll('img.wide')

  for (const img of images) {
    img.style.width = bodyWidth + 'px'
    img.style.marginLeft = leftMargin
  }
}

window.addEventListener('DOMContentLoaded', resizeWideImages)
window.addEventListener('resize', resizeWideImages)

// Go Playground

function setupGoPlaygroundLinks(baseURL, content) {
  const links = document.querySelectorAll('a.go-playground-link')

  for (const link of links) {
    link.addEventListener('click', getGoPlaygroundLink)
  }
}

async function getGoPlaygroundLink(e) {
  const link = e.target
  const code = link.dataset.code
  const isRunnable = link.href[link.href.length - 1] === "#"

  if (code === undefined || !isRunnable) {
    return
  }

  if (isRunnable) {
    e.preventDefault()
  }

  link.innerText = i18n.GO_PLAYGROUND_SENDING

  const response = await fetch(goPlaygroundURL + '/share', {
    method: 'POST',
    body: atob(code)
  })

  if (response.ok) {
    const id = await response.text()
    link.href = goPlaygroundURL + '/p/' + id
    link.innerText = i18n.GO_PLAYGROUND_OPEN
  } else {
    link.innerText = i18n.GO_PLAYGROUND_ERROR
  }
}

window.addEventListener('DOMContentLoaded', setupGoPlaygroundLinks)
