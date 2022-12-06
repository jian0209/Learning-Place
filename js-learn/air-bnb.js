/**
 * https://github.com/airbnb/javascript
 */

// 1.1
const foo = [1, 2]
const bar = foo

bar[0] = 9
console.log('1.1>', foo[0], bar[0]) // => 9, 9

// 2.1
// bad
var a = 1
var b = 1

// good
const a1 = 1
const b1 = 1

function getKey(k) {
  return `a key named ${k}`
}

// 3.2
// bad
const obj = {
  id: 5,
  name: 'San Francisco',
}
obj[getKey('enabled')] = true

// good
const obj2 = {
  id: 5,
  name: 'San Francisco',
  [getKey('enabled')]: true,
}

// 3.3
// bad
const atom = {
  value: 1,
  addValue: function (value) {
    return atom.value + value
  },
}

// good
const atom2 = {
  value: 1,
  addValue(value) {
    return atom.value + value
  },
}
console.log('3.3>', atom2.addValue(123))

// 3.5
const anakinSkywalker = 'Anakin Skywalker'
const lukeSkywalker = 'Luke Skywalker'

// bad
const obj3 = {
  episodeOne: 1,
  twoJediWalkIntoACantina: 2,
  lukeSkywalker,
  episodeThree: 3,
  mayTheFourth: 4,
  anakinSkywalker,
}

// good
const obj4 = {
  lukeSkywalker,
  anakinSkywalker,
  episodeOne: 1,
  twoJediWalkIntoACantina: 2,
  episodeThree: 3,
  mayTheFourth: 4,
}

// 3.8
// very bad
const original = { a: 1, b: 2 }
const copy = Object.assign(original, { c: 3 }) // will mutates original
delete copy.a

// bad
const original2 = { a: 1, b: 2 }
const copy2 = Object.assign({}, original2, { c: 3 })

// good
const original3 = { a: 1, b: 2 }
const copy3 = { ...original3, c: 3 }

function _arrFrom(obj, cb) {
  if (!cb) {
    if (typeof obj === 'string') {
      return obj.split('')
    }
  } else {
    if (obj.length && obj.length > 0) {
      let tempArr = []
      for (let i = 0; i < obj.length; i++) {
        tempArr.push(cb(obj[i]))
      }
      return tempArr
    }
  }
  return 'not the suitable format type'
}

// 4.7
// good
;[1, 2, 3].map((x) => {
  const y = x + 1
  return y
})
;[(1, 2, 3)].map((x) => x + 1)
;[([0, 1], [2, 3], [4, 5])].reduce((acc, item, index) => {
  const flatten = acc.concat(item)
  return flatten
})

// bad
const inbox = [
  {
    subject: 'Mockingbird',
    author: 'Not Harper Lee',
  },
  {
    subject: 'Mockingbird',
    author: 'Harper Lee',
  },
]
const bad = inbox.filter((msg) => {
  const { subject, author } = msg
  if (subject === 'Mockingbird') {
    return author === 'Harper Lee'
  } else {
    return false
  }
})
const good = inbox.filter((msg) => {
  const { subject, author } = msg
  if (subject === 'Mockingbird') {
    return author === 'Harper Lee'
  }
  return false
})

// 5.1
// bad
function getFullName(user) {
  const firstName = user.firstName
  const lastName = user.lastName
  return `${firstName} ${lastName}`
}

// good
function getFullName2(user) {
  const { firstName, lastName } = user
  return `${firstName} ${lastName}`
}

// best
function getFullName3({ firstName, lastName }) {
  return `${firstName} ${lastName}`
}

// 5.2
const arr = [1, 2, 3, 4]
// bad
const first = arr[0]
const second = arr[1]

// good
const [first2, second2] = arr

// 5.3
// bad
// function processInput(input) {
//   return [left, right, top, bottom]
// }
// // caller needs to think about the order of return data
// const [left, __, top] = processInput('something')

// good
// function processInput2(input) {
//   return { left2, right2, top2, bottom2 }
// }
// const { left2, top2 } = processInput2('something')

// 6.3
// bad
function sayHi(name) {
  return 'How are you, ' + name + '?'
}

// good
function sayHi2(name) {
  return `How are you, ${name}`
}

// 7.4
// bad
if ('current user') {
  function test() {
    console.log('nope')
  }
}

// good
let test
if ('current user') {
  test = () => {
    console.log('yup')
  }
}

function _concat(...args) {
  return args.join('')
}

// 7.13
// bad
// never reassign parameter
function f1(a) {
  a = 1
  // ...
}
function f2(a) {
  if (!a) a = 1
  // ...
}

// good
function f3(a) {
  const b = a || 1
  // ...
}
function f4(a = 1) {
  // ...
}

// 9.1
// bad
function Queue(contents = []) {
  this.queue = [...contents]
}
Queue.prototype.pop = function () {
  const value = this.queue[0]
  this.queue.splice(0, 1)
  return value
}

// good
class Queue2 {
  constructor(contents = []) {
    this.queue = [...contents]
  }

  pop() {
    const value = this.queue[0]
    this.queue.splice(0, 1)
    return value
  }
}

// 9.2
// use extends for inheritance
class PeekableQueue extends Queue2 {
  constructor() {
    super()
  }

  peek() {
    return this.queue[0]
  }
}

// 15.7
// bad
const foo2 = a ? a : b
const bar2 = c ? true : false
const quux = a !== null ? a : b

// good
const foo3 = a || b
const bar3 = !!c
const quux2 = a ?? b
