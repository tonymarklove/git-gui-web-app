jQuery(function($) {
  var conn = new WebSocket("ws://localhost:8080/ws");

  conn.onclose = function(evt) {
    console.log("Connection closed");
  };

  conn.onmessage = function(evt) {
    $("[data-target=changed-files]").html(evt.data);
  };


  $("form").on("submit", function(e) {
    e.preventDefault();

    conn.send($("textarea").val());
  });
});
