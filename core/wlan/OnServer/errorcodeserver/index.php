<?php
$modedata = $_POST["modeinteger"]
header("Content-Type: text/plain");
if (modedata == 0) {
  echo 403;
} elseif (modedata == 1) {
  echo 404;
} elseif (modedata == 2) {
  echo 410;
} elseif (modedata == 3) {
  echo 500;
} elseif (modedata == 4) {
  echo 501;
} elseif (modedata == 5) {
  echo 200;
} elseif (modedata == 6) {
  echo 502;
} elseif (modedata == 7) {
  echo 503;
} elseif (modedata == 8) {
  echo 504;
} elseif (modedata == 9) {
  echo 505;
} else {
  die();
}
?>
