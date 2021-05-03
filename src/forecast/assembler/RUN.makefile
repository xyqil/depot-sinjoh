executor:
	chmod 777 $$PWD/EXECUTE.sh
	key=$$(openssl rand -base64 242)
	echo Forecast is running............ > /home/loggeruser/log-$$key
	wall Forecast is running............
	bash $$PWD/EXECUTE.sh
	wall Forecast was ran!
builder:
	chmod 777 $$PWD/BUILD.sh
	bash $$PWD/BUILD.sh