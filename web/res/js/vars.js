var ws = new WebSocket('ws://' + location.href.split('/')[2] + '/ws')
ws.onopen = function () {
    console.log("ws open")
}
ws.onclose = function () {
    window.close()
}