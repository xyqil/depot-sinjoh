<?php
	require_once "random_compat/lib/random.php";
	require_once "php-csprng/support/random.php";
	header("Content-Type: text/plain");
	http_response_code(410);
	$rng = new CSPRNG();
	for ($x = 0; $x < 100; $x++)
	{
		$result = $rng->GetInt(0, 16384);
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
