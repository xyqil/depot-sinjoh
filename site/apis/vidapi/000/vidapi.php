<?php
header("Content-Type: application/javascript");
echo "var video = document.getElementById(\"myVideo\");\n";
echo "var btn = document.getElementById(\"myBtn\");\n";
echo "function pausebuttoninitalizer(nothing) {;\n";
echo "	if(video.paused) {;\n";
echo "		video.play();\n";
echo "		btn.innerHTML = \"Pause\";\n";
echo "		return nothing;\n";
echo "	} else {"
echo "		video.pause();\n";
echo "		btn.innerHTML = \"Play\";\n";
echo "		return nothing;\n";
echo "	};\n";
echo "};\n";
?>
