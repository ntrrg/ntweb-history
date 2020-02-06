function Base64Encode(str) {
  const encoder = new TextEncoder()
  return base64js.fromByteArray(encoder.encode(str))
}

function Base64Decode(str) {
  const decoder = new TextDecoder()
  return decoder.decode(base64js.toByteArray(str));
}

async function buildSearchIndex() {
  const res = await fetch('../search-index/index.json')
  const data = await res.json()
  window.idxData = new Map()

  window.idx = lunr(function () {
    this.ref('url')
    this.field('type')
    this.field('title')
    this.field('description')
    this.field('content')

    data.forEach(function (doc) {
      doc.content = Base64Decode(doc.content)
      this.add(doc)
      window.idxData.set(doc.url, doc)
    }, this)
  })
}

async function search(q) {
  window.searching = true

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

  window.searching = false
}

window.addEventListener('DOMContentLoaded', () => {
  const params = new URLSearchParams(window.location.search)
  const q = params.get('q')

  if (q) {
    search(q)
    document.querySelector('#search-box input').value = q
  }
})

document.querySelector('#search-box').addEventListener('submit', (e) => {
  const q = e.target.querySelector('input').value

  if (!window.searching && q) {
    search(q)
  }

  e.preventDefault()
})

