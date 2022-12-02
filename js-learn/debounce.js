function debounce(fn, delay) {
  let timer = null // 创建一个标记用来存放定时器的返回值
  return function (e) {
    clearTimeout(timer) // 每当用户输入的时候把前一个 setTimeout clear 掉
    timer = setTimeout(() => {
      fn.apply(this, arguments)
    }, delay) // 然后又创建一个新的 setTimeout, 这样就能保证输入字符后的 interval 间隔内如果还有字符输入的话，就不会执行 fn 函数
  }
}

// 处理函数
function handle() {
  console.log(Math.random())
}

// 滚动事件
window.addEventListener('scroll', debounce(handle, 500))

function throttle(fn, delay) {
  let canRun = true // 通过闭包保存一个标记
  return function () {
    if (!canRun) return // 在函数开头判断标记是否为true，不为true则return
    canRun = false // 立即设置为false
    setTimeout(() => {
      // 将外部传入的函数的执行放在setTimeout中
      fn.apply(this, arguments)
      // 最后在setTimeout执行完毕后再把标记设置为true(关键)表示可以执行下一次循环了。
      // 当定时器没有执行的时候标记永远是false，在开头被return掉
      canRun = true
    }, delay)
  }
}

function sayHi(e) {
  console.log('节流成功', e.target.innerWidth, e.target.innerHeight)
}

window.addEventListener('resize', throttle(sayHi, 500))
