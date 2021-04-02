<?php
// Kept for legacy purposes
$data0 = "<h1>Done!</h1>"
$data1 = shell_exec('export data=source/main/');
$data2 = shell_exec('cd $data');
$data3 = shell_exec('rm -rf main2.go');
$data4 = shell_exec('rm -rf main');
$data5 = shell_exec('rm -rf ./main');
$data6 = "pre"
$data7 = "<"
$data8 = ">"
$data9 = "/pre"
header("Content-Type:text/html")
echo $data7+$data6+$data8+$data1+$data7+$data9+$data8+$data7+$data6+$data8+$data2+$data7+$data9+$data8+$data7+$data6+$data8+$data3
echo $data7+$data9+$data8+$data7+$data6+$data8+$data4+$data7+$data9+$data8+$data7+$data6+$data8+$data5+$data7+$data9+$data8+$data0
?>
