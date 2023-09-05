import { Request, POST } from "/manager/public/script/http/http.js"

window.onload = function () {
    var login = document.getElementById("submit-manager-login")

    login.addEventListener("click", authManager)
}

var authManager = function () {
    var login = document.getElementById("manager-login").value
    var password = document.getElementById("manager-password").value

    var manager = {
        "login": login,
        "pwd": password,
    }
    console.log(JSON.stringify(manager))
    Request(POST, "/manager/api/login", manager, function () {
        window.location = "/manager/"
    }
    )
}