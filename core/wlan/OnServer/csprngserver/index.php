<?php
	require_once "random_compat/lib/random.php";
	require_once "php-csprng/support/random.php";
	require_once "revnode/xml/xml.php";
	session_start();
	mysqli_report(MYSQLI_REPORT_ERROR | MYSQLI_REPORT_STRICT);
	$link = new mysqli('localhost', 'username', 'password', 'db_name');
	$link->set_charset('utf8mb4'); // always set the charset
	$name = $_GET["username"];
	$stmt1 = $link->prepare("SELECT primaryintegerdata FROM csprngserver WHERE values=? limit 1");
	$stmt1->bind_param('s', $name);
	$stmt1->execute();
	$result = $stmt1->get_result();
	$value1 = $result->fetch_object();
	$_SESSION['myid'] = $value1->id;
	$link->set_charset('utf8mb4'); // always set the charset
	$name = $_GET["username"];
	$stmt2 = $link->prepare("SELECT secondaryintegerdata FROM csprngserver WHERE values=? limit 1");
	$stmt2->bind_param('s', $name);
	$stmt2->execute();
	$result = $stmt2->get_result();
	$value2 = $result->fetch_object();
	$_SESSION['myid'] = $value2->id;
	$link->set_charset('utf8mb4'); // always set the charset
	$name = $_GET["username"];
	$stmt3 = $link->prepare("SELECT configurationdata FROM csprngserver WHERE values=? limit 1");
	$stmt3->bind_param('s', $name);
	$stmt3->execute();
	$result = $stmt3->get_result();
	$value3 = $result->fetch_object();
	$_SESSION['myid'] = $value3->id;
	$url1 = $value2+"/index.php?integerdata="+$value3
	$ch1 = curl_init();  
	curl_setopt($ch1, CURLOPT_URL, $url1);  
	curl_setopt($ch1, CURLOPT_HEADER, $value1);  
	curl_setopt($ch1, CURLOPT_RETURNTRANSFER, true);  
	$value1 = curl_exec($ch1);  
	curl_close($ch1);
	$url2 = $value2+"/index.php?integerdata="+$data1;
	$url3 = $value2+"/index.php?integerdata="+$data2;
	$ch2 = curl_init();  
	curl_setopt($ch2, CURLOPT_URL, $url2);  
	curl_setopt($ch2, CURLOPT_HEADER, $value1);  
	curl_setopt($ch2, CURLOPT_RETURNTRANSFER, true);  
	$value2 = curl_exec($ch2);  
	curl_close($ch2);
	$ch3 = curl_init();  
	curl_setopt($ch3, CURLOPT_URL, $url3);  
	curl_setopt($ch3, CURLOPT_HEADER, $value1);  
	curl_setopt($ch3, CURLOPT_RETURNTRANSFER, true);  
	$value3 = curl_exec($ch3);  
	curl_close($ch1);
	$random = new Rych\Random\Random();
	$randomnumber1 = $random->getRandomInteger($value2, $value3);
	$randomnumber2 = $random->getRandomInteger($value2, $value3);
	header("Content-Type: text/plain");
	$rng = new CSPRNG();
	for ($x = $randomnumber1; $x < $randomnumber2; $x++)
	{
		$result = $rng->GetInt($value2, $value3);
	}
	try {
		$string = random_bytes($result);
	} catch (TypeError $e) {
		// Well, it's an integer, so this IS unexpected.
		die("An unexpected error has occurred"); 
	} catch (Error $e) {
		// This is also unexpected because 32 is a reasonable integer.
		die("An unexpected error has occurred");
	} catch (Exception $e) {
		// If you get this message, the CSPRNG failed hard.
		die("Could not generate a random string. Is our OS secure?");
	}
	echo bin2hex($string);
	exit;
?>
