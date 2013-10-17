$(function() {
  var conn;
  var msg = $("#msg");
  var output = $("#output");

  $("#form").submit(function() {
    if (!conn) {
      return false;
    }
    if (!msg.val()) {
      return false;
    }
    conn.send(msg.val());
    msg.val("");
    return false
  });

  conn = new WebSocket("ws://{{$}}/ws");
  conn.onclose = function(evt) {
  }
  conn.onmessage = function(evt) {
    output.html(evt.data);
  }
});
