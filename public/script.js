$(document).ready(function() {
  var entryContentElement = $("#note-input");

  var appendNoteList = function(data) {
    if (data == null) {
      return
    }
    $("#Notes ul").empty();
    $.each(data, function(key, val) {
      $("#Notes ul").append('<li><div class="col s12 m6 l4"><div class="card-panel purple accent-4 z-depth-5"><span class="white-text">'+val+'<p><div class="switch"><label><input type="checkbox" name="deleteCheck" value="1"><span class="lever"></span></label></div></p></span></div></div></li>');
    });
  }

  var handleSubmission = function(e) {
    e.preventDefault();
    var entryValue = entryContentElement.val()
    if (!entryValue || entryValue.length <= 0){
        entryContentElement.parent().addClass("has-error").removeClass("has-success");
        return false;
    }

    entryContentElement.val("")
    entryContentElement.parent().removeClass("has-error").addClass("has-success");
    $.getJSON("insert/note/" + entryValue, appendNoteList);
  }

  var handleDeletion = function(e){
    e.preventDefault();

    var checkboxes = document.getElementsByName("deleteCheck");
    for (var i=0; i < checkboxes.length; i++) {
     if (!checkboxes[i].checked){
       continue
     }
     var checkbox = checkboxes[i];
     $.getJSON("delete/note/" + $(checkbox).closest('li').text(), appendNoteList);
    }
  }

  $("#note-submit").click(handleSubmission);
  $("#note-delete").click(handleDeletion);

  // Poll every five second.
  (function fetchNotes() {
    $.getJSON("read/note").done(appendNoteList).always(
      function() {
        setTimeout(fetchNotes, 50000);
      });
  })();
});

$(document).ready(function() {
   $.getJSON("version", function(data) {
    if (data == null) {
      return
    }
    $("#footer-version").text("Version: " + data["version"]);
  });
});
