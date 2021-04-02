<?php
header("Content-Type: application/javascript");
echo "function initalizemii(){\n";
echo "  var s = document.createElement(\"script\");\n";
echo "  s.src = \"https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js\\n";
echo "  s.onload = function(e){initalizeotherscripts()};\n";
echo "  document.head.appendChild(s);\n";
echo "  return Math.random().toString(36).substring(2);\n";
echo "};\n";
echo "function initalizeotherscripts(){\n";
echo "  $.getScript('apis/nullapi/000/nullapi.php', function() {\n";
echo "    console.log(\"apis/nullapi/000/nullapi.php was loaded\")\n";
echo "  });\n";
echo "  $.getScript('apis/vidapi/000/vidapi.js', function() {\n";
echo "    console.log(\"apis/vidapi/000/vidapi.php was loaded\")\n";
echo "});\n";
echo "  $.getScript('https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js', function() {\n";
echo "    console.log(\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js was loaded\")\n";
echo "  });\n";
echo "  document.title = \"XConnect24\"\n";
echo "  initalizethecss();\n";
echo "  return console.log(\"multiplexers/000/css/multiplexmii.php was loaded\");\n";
echo "};\n";
echo "function initalizethecss(){\n";
echo "  return $('body').append('<link rel=\"stylesheet\" href=\"multiplexers/000/css/multiplexmii.php\" type=\"text/css\"></div></body><div class=\"d\" id=\"sessionseed\" hidden>'+Math.random().toString(36).substring(2)+'</div></html>');\n";
echo "}\n";
echo "window.onload = initalizemii();\n";
?>
