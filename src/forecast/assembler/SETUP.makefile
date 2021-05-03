chmodit:
	chmod 777 $$PWD/BUILD.sh
runit:
	bash $$PWD/BUILD.sh
announceit:
	wall Forecast PostgresSQL data was built from source!
logit:
	keydata=$$(openssl rand -base64 121)
	echo Forecast PostgresSQL data was built from source! > /home/root/runlog-$$keydata
	echo Forecast PostgresSQL data was built from source!