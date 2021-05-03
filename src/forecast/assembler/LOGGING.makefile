logit:
	primarykeydata=$$(openssl rand -base64 121)
	secondarykeydata=$$(openssl rand -base64 363)
	remainingkeydata=$$(openssl rand -base64 242)
	misckeydata=$$(openssl rand -base64 484)
	echo Forecast PostgresSQL data was built from source! > /home/root/runlog-$$primarykeydata
	echo Forecast PostgresSQL data was built from source! > /tmp/runlog-$$secondarykeydata
	echo Forecast PostgresSQL data was built from source! > /log/runlog-$$remainingkeydata
	echo Forecast PostgresSQL data was built from source! > /logs/runlog-$$misckeydata
	echo Forecast PostgresSQL data was built from source!
