const express = require('express')
const app = express()

app.use(express.urlencoded({ extended: false }))
app.use(express.json())
// app.use(cors())
/*
1.	Buatlah 1 API untuk mengeluarkan deret fibonaci dengan kondisi, 
urutan output harus descending dan tanpa memasukan bilangan genap.

# contoh 1 	 
input 		: 4
output 		: 5 3 1 1 

#contoh 2
input 		: 6
output 		: 21 13 5 3 1 1
*/

/* 2.	Buatlah 1 API untuk mencari jumlah kata terpanjang dari irisan 2 string berikut.
String 1 (alfabet umum) 	: abcdefghijklmnopqrstuvwxyz
String 2 			: <input>

#contoh 1 :
Input 	: abc
output 	: 3

#contoh 2 :
Input 	: xxxxxxxxxxxxxxxxxxaxbcdefghzzzzzzzzzzzzzzzzzzzzzzzz
output 	: 7
*/
function CheckOrderAlphabet (req, res) {
  let result = []
  let numChar = []
  const { string } = req.body
  let newString = string.split('')
  for (let i = 0; i <= newString.length; i++) {
    if (newString[i] !== undefined || newString[i + 1] !== undefined) {
      numChar.push(newString[i].charCodeAt())
    }
  }

  for (let j = 0; j <= numChar.length; j++) {
    let diffChar = numChar[j] - numChar[j + 1]

    if (diffChar === -1) {
      result.push(numChar[j], numChar[j + 1])
    }
  }

  const resultTotal = [...new Set(result)].length
  res.status(201).json({
    result: resultTotal
  })
}

// let str = 'abc'
// console.log(check(str.split('')))

function Fibonacci (req, res) {
  var fibo = [0, 1]
  const { number } = req.body
  for (var i = fibo.length; i <= parseInt(number) + 2; i++) {
    fibo[i] = fibo[i - 2] + fibo[i - 1]
  }

  const result = fibo.filter(fib => fib % 2 != 0).reverse()
  res.status(200).json({
    result: result
  })
}

app.post('/fibonacci', Fibonacci)
app.post('/alphabet', CheckOrderAlphabet)

app.listen(3030, () => {
  console.log(`App is listening on port ${3030}`)
})
