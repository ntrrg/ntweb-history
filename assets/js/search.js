async function buildSearchIndex() {
  const res = await fetch('../search-index/index.json')
  const data = await res.json()

  window.idx = lunr(function () {
    this.ref('url')
    this.field('title')
    this.field('description')
    this.field('content')

    data.forEach(function (doc) {
      doc.content = atob(doc.content)
      this.add(doc)
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
    const res = await fetch(result.ref + '/index.json')
    const page = await res.json()
    let html = ''

    html += '<li>'
    html += '<div class="card">'
    html += '<a href="' + page.url + '">'
    html += '<h3>' + page.title + '</h3>'
    html += '</a>'
    html += '<p>' + page.params.description + '</p>'
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

