// Task list

window.addEventListener('DOMContentLoaded', () => {
  let selector = '.markdown ul li input[type="checkbox"]'
  const inputs = document.querySelectorAll(selector)

  for (const input of inputs) {
    const list = input.parentNode.parentNode
    list.classList.add('task-list')
  }
})

// Wide images

function resizeWideImages() {
  let leftMargin = '-1em'
  const bodyWidth = document.querySelector('body').offsetWidth
  const mainWidth = document.querySelector('body > main').offsetWidth

  if (bodyWidth !== mainWidth) {
    leftMargin = 'calc(' + (bodyWidth - mainWidth) / 2 * -1 + 'px - 1em)'
  } 

  const images = document.querySelectorAll('img.wide, figure.wide img')

  for (const img of images) {
    img.style.width = bodyWidth + 'px'
    img.style.marginLeft = leftMargin
  }
}

window.addEventListener('DOMContentLoaded', resizeWideImages)
window.addEventListener('resize', resizeWideImages)

// Go Playground

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

window.addEventListener('DOMContentLoaded', () => {
  const links = document.querySelectorAll('a.go-playground-link')

  for (const link of links) {
    link.addEventListener('click', getGoPlaygroundLink)
  }
})

// Mermaid

window.addEventListener('load', () => {
  if (typeof mermaid === 'undefined') {
    return
  }

  mermaid.mermaidAPI.initialize({
    startOnLoad: false,
    fontFamily: 'monospace',
    theme: 'forest',
    flowchart: {
      useMaxWidth: false
    },
    sequence: {
      useMaxWidth: false
    },
    gantt: {
      useMaxWidth: false
    }
  })

  const charts = document.querySelectorAll('.mermaid-chart code')

  for (const chart of charts) {
    const id = 'mermaid-' + Math.floor(Math.random() * 1000)

    mermaid.mermaidAPI.render(id, chart.innerText, (data) => {
      const target = chart.parentNode.parentNode.querySelector('.mermaid')
      target.innerHTML = data
    })
  }
})

