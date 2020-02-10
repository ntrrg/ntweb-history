function base64Encode(str) {
  const encoder = new TextEncoder()
  return base64js.fromByteArray(encoder.encode(str))
}

function base64Decode(str) {
  const decoder = new TextDecoder()
  return decoder.decode(base64js.toByteArray(str));
}

function startLoader(interval = 100) {
  const el = document.querySelector('#search-box .loader')

  if (typeof el.dataset.interval !== 'undefined') {
    return
  }

  let count = 0

  el.dataset.interval = setInterval(() => {
    switch (count % 4) {
      case 0:
        el.innerText = '|'
        break
      case 1:
        el.innerText = '/'
        break
      case 2:
        el.innerText = '-'
        break
      case 3:
        el.innerText = '\\'
        break
    }

    count++
  }, interval)
}

function stopLoader() {
  const el = document.querySelector('#search-box .loader')

  if (typeof el.dataset.interval === 'undefined') {
    return
  }

  clearInterval(el.dataset.interval)
  el.innerText = ''
  el.removeAttribute('data-interval')
}

async function buildSearchIndex() {
  const input = document.querySelector('#search-box input')
  input.setAttribute('disabled', '')
  startLoader()

  const res = await fetch('../search-index/index.json')
  const data = await res.json()
  window.idxData = new Map()

  window.idx = lunr(function () {
    this.ref('url')
    this.field('type')
    this.field('title')
    this.field('description')
    this.field('content')
    this.field('tags')

    data.forEach(function (doc) {
      doc.content = base64Decode(doc.content)
      this.add(doc)
      window.idxData.set(doc.url, doc)
    }, this)

    input.removeAttribute('disabled')
    stopLoader()
  })
}

async function search(q) {
  if (window.idx === undefined) {
    await buildSearchIndex()
  }

  const resultsEl = document.querySelector('#search-results')
  const results = idx.search(q)

  resultsEl.innerHTML = ''

  for (const result of results) {
    const page = window.idxData.get(result.ref)
    let html = ''

    html += '<li>'
    html += '<div class="card">'
    html += '<a href="' + page.url + '">'
    html += '<h3>' + page.title + '</h3>'
    html += '</a>'
    html += '<p>' + page.description + '</p>'
    html += '</div>'
    html += '</li>'

    resultsEl.innerHTML += html
  }
}

window.addEventListener('DOMContentLoaded', () => {
  const params = new URLSearchParams(window.location.search)
  const q = params.get('q')

  if (q) {
    document.querySelector('#search-box input').value = q
    search(q)
  }
})

document.querySelector('#search-box input').addEventListener('keyup', (e) => {
  const q = e.target.value

  if (q) {
    search(q)
  }
})

document.querySelector('#search-box').addEventListener('submit', (e) => {
  const q = e.target.querySelector('input').value

  if (q) {
    search(q)
  }

  e.preventDefault()
})

