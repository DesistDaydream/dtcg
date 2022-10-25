let FormEle = document.querySelector("form")
let DeckEle = FormEle.querySelector("textarea[name=deck]")
let StockOutEle = FormEle.querySelector("button[name=commit]")
let tbodyCardsList = document.getElementById("cards_list")
let trRespPrice = document.getElementById("resp_price")

StockOutEle.onclick = function (event) {
  event.preventDefault()

  // 清理表格。防止上次的查询残留
  while (tbodyCardsList.hasChildNodes()) {
    tbodyCardsList.removeChild(tbodyCardsList.lastChild)
  }
  while (trRespPrice.hasChildNodes()) {
    trRespPrice.removeChild(trRespPrice.lastChild)
  }

  // 创建请求
  let xhr = new XMLHttpRequest()
  xhr.open("POST", "http://localhost:2205/api/v1/deck/price")
  xhr.send(
    JSON.stringify(
      (reqBody = {
        deck: String(DeckEle.value),
        envir: "chs",
      })
    )
  )
  xhr.onload = function () {
    let resp = JSON.parse(xhr.responseText)
    console.log(resp)

    // 生成第一张表
    let minPrice = document.createElement("td")
    let avgPrice = document.createElement("td")
    minPrice.innerText = resp.min_price.toFixed(2)
    avgPrice.innerText = resp.avg_price.toFixed(2)
    trRespPrice.appendChild(minPrice)
    trRespPrice.appendChild(avgPrice)

    // 生成第二张表
    for (let i = 0; i < resp.data.length; i++) {
      let tr = document.createElement("tr")
      let tdScName = document.createElement("td")
      let tdCount = document.createElement("td")
      let tdSerial = document.createElement("td")
      let tdMinPrice = document.createElement("td")
      let tdAvgPrice = document.createElement("td")
      tdScName.innerText = resp.data[i].sc_name
      tdCount.innerText = resp.data[i].count
      tdSerial.innerText = resp.data[i].serial
      tdMinPrice.innerText = resp.data[i].min_price.toFixed(2)
      tdAvgPrice.innerText = resp.data[i].avg_price.toFixed(2)
      tr.appendChild(tdScName)
      tr.appendChild(tdCount)
      tr.appendChild(tdSerial)
      tr.appendChild(tdMinPrice)
      tr.appendChild(tdAvgPrice)
      tbodyCardsList.appendChild(tr)
    }
  }
}
