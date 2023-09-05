import { Request, POST } from "/manager/public/script/http/http.js"

window.onload = function () {
    var taskCreateSubmitBtn = document.getElementById("submit-manager-register")

    taskCreateSubmitBtn.addEventListener("click", managerRegisterEvent)
}

var managerRegisterEvent = function () {
    var name = document.getElementById("manager-name").value
    var surname = document.getElementById("manager-surname").value
    var login = document.getElementById("manager-login").value
    var password = document.getElementById("manager-password").value

    var manager = {
        "name": name,
        "surname": surname,
        "login": login,
        "pwd": password,
        "e_title": title
    }
    console.log(JSON.stringify(manager))
    Request(POST, "/enterprise/api/enterprises/" + title + "/register/manager", manager, function () {
        window.location = window.location
    }
    )
}