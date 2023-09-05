export var POST = "POST";

export function Request(method, path, payload, onSuccess) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, path);
    xhr.setRequestHeader("Accept", "application/json");
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.withCredentials = true
    console.log("send " + payload)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if(xhr.status == 200) {
                onSuccess()
                //window.location = "/enterprise/enterprises"
            } else {
                console.log(xhr.status)
                //window.location = window.location
            }
        }
    };
    xhr.send(JSON.stringify(payload));
}