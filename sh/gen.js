const delivery = require('./delivery.json')
const orders = require('./orders.json')
const payment = require('./payment.json')
const items = require('./items.json')

const rndI = (max) => {
  return Math.floor(Math.random() * max)
}

const randEl = (arr) => {
  const len = arr.length
  const idx = rndI(len)
  return arr[idx]
}

const result = []

const uidMap = {}
for (let i = 0; i < orders.length; i++) {
  const order = orders[i]
  const id = order.order_uid
  if (uidMap[id] != null) {
    throw Error("same key for " + id)
  }

  order.payment = randEl(payment)
  order.delivery = randEl(delivery)
  order.items = Array(rndI(5)).fill(null).map(e => randEl(items))
  result.push(order)
}

require('fs').writeFileSync("./all.json", JSON.stringify(result))
console.log('done')
