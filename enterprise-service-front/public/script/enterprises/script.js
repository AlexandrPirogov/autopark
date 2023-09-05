import {Request, POST} from "/enterprise/public/script/http/http.js"

window.onload = function () {
    var taskCreateSubmitBtn = document.getElementById("submit-enterprise-create")
    
    taskCreateSubmitBtn.addEventListener("click", enterpriseCreateEvent)
}

var enterpriseCreateEvent = function() {
    var title = document.getElementById("enterprise-title").value
    var enterprise = {
        "title" : title
    } 

    Request(POST, "/enterprise/api/enterprises/register", enterprise, function() {
        console.log("nothing")
        }
    )
}