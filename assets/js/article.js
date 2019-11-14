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

