const ArrayForSort = []
// generate random number array
function GenerateRandomNumberArray(length) {
  for (let i = 0; i < length; i++) {
    ArrayForSort.push(Math.floor(Math.random() * length * 10000))
  }
}
GenerateRandomNumberArray(10)
console.log(ArrayForSort)

// bubble sort
// loop inside loop and compare adjacent number
// if true, swap them
// if false, continue
function BubbleSort() {
  let tempArr = [...ArrayForSort]
  for (let i = 0; i < tempArr.length; i++) {
    for (let j = 0; j < tempArr.length - 1; j++) {
      if (tempArr[j] > tempArr[j + 1]) {
        let temp = tempArr[j]
        tempArr[j] = tempArr[j + 1]
        tempArr[j + 1] = temp
      }
    }
  }
  return tempArr
}

// selection sort
// loop inside loop and compare the smallest number
// if true, store to variable and continue
// if false, continue
// finally, swap the smallest number to the first position
function SelectionSort() {
  let tempArr = [...ArrayForSort]
  for (let i = 0; i < tempArr.length; i++) {
    let lowestNum = i
    for (let j = i + 1; j < tempArr.length; j++) {
      if (tempArr[j] < tempArr[lowestNum]) {
        lowestNum = j
      }
    }
    let temp = tempArr[i]
    if (temp !== tempArr[lowestNum]) {
      tempArr[i] = tempArr[lowestNum]
      tempArr[lowestNum] = temp
    }
  }
  return tempArr
}

// insertion sort
// loop to the first position and compare with the previous number
// if true, swap them and continue
// if false, continue
function InsertionSort() {
  let tempArr = [...ArrayForSort]
  for (let i = 0; i < tempArr.length; i++) {
    let temp = tempArr[i]
    let j = i - 1
    while (j >= 0 && tempArr[j] > temp) {
      tempArr[j + 1] = tempArr[j]
      j--
    }
    tempArr[j + 1] = temp
  }
  return tempArr
}

// merge sort
// divide the array into two parts and merge them
// if the array is less than 2, return the array
function MergeSort(arr) {
  if (arr.length < 2) {
    return arr
  }
  let mid = Math.floor(arr.length / 2)
  let left = MergeSort(arr.slice(0, mid))
  let right = MergeSort(arr.slice(mid))
  return Merge(left, right)
}
function Merge(left, right) {
  let result = []
  while (left.length && right.length) {
    if (left[0] <= right[0]) {
      result.push(left.shift())
    } else {
      result.push(right.shift())
    }
  }
  return result.concat(left, right)
}

// quick sort
function QuickSort(arr) {
  if (arr.length < 2) {
    return arr
  }
  let pivot = arr[0]
  let left = []
  let right = []
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i])
    } else {
      right.push(arr[i])
    }
  }
  return QuickSort(left).concat(pivot, QuickSort(right))
}

// random-quick sort
function RandomQuickSort(arr) {
  if (arr.length < 2) {
    return arr
  }
  let pivot = arr[Math.floor(Math.random() * arr.length)]
  let left = []
  let right = []
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i])
    } else {
      right.push(arr[i])
    }
  }
  return RandomQuickSort(left).concat(pivot, RandomQuickSort(right))
}

// counting sort
function CountingSort(arr) {}
