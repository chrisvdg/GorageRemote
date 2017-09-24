$(document).ready(function(){

    // create websocket
    var ws = new WebSocket("wss://" + window.location.host + "/api/actionsocket");

    ws.onopen = function()
    {
       console.log("websocket is open")
    }

    ws.onclose = function()
    { 
       console.log("websocket closed..."); 
    };

    // Action commands
    //multifunction button (1 button for open/close)
    $("#btnMulti").click(function() {
        ws.send("multi");
    });

    // open button
    $("#btnOpen").click(function() {
        ws.send("open");
    });

    // close button
    $("#btClose").click(function() {
        ws.send("close");
    });

    // stop button
    $("#btStop").click(function() {
        ws.send("stop");
    });
});