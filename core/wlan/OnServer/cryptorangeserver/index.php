<?php
require_once "revnode/xml/xml.php";
$integerdata = $_POST["numberdata"]
header("Content-Type: text/plain");
if ($integerdata == 1) {
  echo 0;
} elseif ($integerdata == 2) {
  echo 16384;
} elseif ($integerdata == 3) {
  $xml = new xml('<?xml version="1.0" encoding="ISO-8859-1"?>
<primarynode>1</primarynode>
<secondarynode>16384</secondarynode>
</breakfast_menu>');
  echo $xml->data
} else {
  die();
}
?>
