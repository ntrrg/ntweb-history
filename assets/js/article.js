// Wide images

function resizeWideImages() {
  let leftMargin = '-1em'
  let bodyWidth = document.querySelector('#body').offsetWidth
  let mainWidth = document.querySelector('#body > main').offsetWidth

  if (bodyWidth !== mainWidth) {
    leftMargin = 'calc(' + (bodyWidth - mainWidth) / 2 * -1 + 'px - 1em)'
  }

  let images = document.querySelectorAll('img.wide')

  for (let img of images) {
    img.style.width = bodyWidth + 'px'
    img.style.marginLeft = leftMargin
  }
}

window.addEventListener('load', resizeWideImages)
window.addEventListener('resize', resizeWideImages)

