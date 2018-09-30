package ui

func Index() string {
  return `
<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <style>
  .row {
    display: flex; 
    flex-direction: row;
    flex: 1;
  }

  .button {
    flex: 1;
  }

  #buttons {
    display: flex;
    flex-direction: column;
    position:fixed;
    padding:0;
    margin:0;

    top:0;
    left:0;

    width: 100%;
    height: 100%;
  }
  </style>
</head>
<body>
  <div id="buttons">
  </div>


 <script>

  function addListener(element, eventName, handler) {
    if (element.addEventListener) {
      element.addEventListener(eventName, handler, false);
    }
    else if (element.attachEvent) {
      element.attachEvent('on' + eventName, handler);
    }
    else {
      element['on' + eventName] = handler;
    }
  }

  function removeListener(element, eventName, handler) {
    if (element.addEventListener) {
      element.removeEventListener(eventName, handler, false);
    }
    else if (element.detachEvent) {
      element.detachEvent('on' + eventName, handler);
    }
    else {
      element['on' + eventName] = null;
    }
  }


  const handle = function(button) {
    return function() {

      var r = new XMLHttpRequest();
      r.open("POST", "/handle", true);
      r.onreadystatechange = function () {
        if (r.readyState != 4 || r.status != 200) return;
        console.log("Success: " + r.responseText);
      };
      r.send("button=" + button);
    };
  };
  const descriptions = [
       [{
         "label": "Hello", 
         "id": "hello", 
        },
        {
         "label": "World",
         "id": "world"
        }
       ],
       [{
         "label": "Ich",
         "id": "ich", 
        },
        {
         "label": "Bin",
         "id": "bin"
        },
        {
         "label": "Hamo",
         "id": "hamo"
        }]
    ]

  const html = descriptions.map( 
    function(row) {
      return  "<div class='row'>\n" + 
        row.map( 
        function (button) {
          return "<button class='button' id='" + button["id"] + "'>" + button['label'] + "</button>"; 
        }
      ).join("\n") + 
        "</div>";
     }
  ).join("\n");

  const root = 	document.getElementById("buttons");
  root.innerHTML = html; 

  const buttons = Array.from(document.getElementsByTagName("button")); 

  buttons.forEach( function(button) {
    addListener(button, "click", 
      function(e) { 
        handle(e.target.id)(); 
      });   
  }); 

  </script>

</body>
</html>




  `
}
