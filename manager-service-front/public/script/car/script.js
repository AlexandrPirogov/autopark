import { Request, POST } from "/manager/public/script/http/http.js"

window.onload = function () {
    var registerCarBtn = document.getElementById("submit-car-register")

    registerCarBtn.addEventListener("click", carRegisterEvent)
}

var carRegisterEvent = function () {
    var uid = document.getElementById("car-uid").value
    var type = document.getElementById("car-type").value
    var brand = document.getElementById("car-brand").value

    var car = {
        "uid": uid,
        "type": type,
        "brand": brand,
    }
    console.log(JSON.stringify(car))
    Request(POST, "/manager/api/car/register", car, function () {
        window.location = window.location
    }
    )
}