jQuery(function($) {
  var conn = new WebSocket("ws://localhost:8080/ws");

  conn.onclose = function(evt) {
    console.log("Connection closed");
  };

  conn.onmessage = function(evt) {
    eventDispatch(JSON.parse(evt.data));
  };

  var events = {
    "raw": function(evt) {
      gitConsole.log(evt);
    },
    "changedFiles": function(evt) {
      var ul = $("[data-target=changed-files]");
      changedFilesUl.empty();

      evt.changedFiles.forEach(function(file) {
        var li = $("<li>").html(file);
        changedFilesUl.append(li);
      });
    },
    "fileDiff": function(evt) {
      $("[data-target=file-viewer]").html(evt.data);
    }
  };

  function eventDispatch(evt) {
    if (!events[evt.type]) {
      alert("Unrecognised event type");
      return;
    }

    return events[evt.type](evt);
  }


  var changedFilesUl = $("[data-target=changed-files]");

  changedFilesUl.on("click", function(e) {
    e.preventDefault();

    var li = $(e.target).closest("li");
    if (!li) {
      return;
    }

    var file = li.text();

    var command = {
      type: "diffFile",
      file: file
    };

    conn.send(JSON.stringify(command));
  });


  var commitForm = $("[data-target=commit-form]");

  commitForm.on("submit", function(e) {
    e.preventDefault();

    var command = {
      type: "commit",
      file: $("[data-target=commit-message]").val()
    };

    conn.send(JSON.stringify(command));
  });


  // Git console
  var gitConsole = (function() {
    var gitConsole = $("[data-target=git-console]");
    var form = gitConsole.find("form");
    var commandLine = gitConsole.find("input");
    var output = gitConsole.find("pre");

    form.on("submit", function(e) {
      e.preventDefault();

      var command = {
        type: "raw",
        command: {
          command: commandLine.val()
        }
      };

      conn.send(JSON.stringify(command));
    });

    return {
      log: function(evt) {
        output.html(evt.data);
      }
    };
  })();
});
