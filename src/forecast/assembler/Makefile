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
chmoddlscript:
	chmod 777 $$PWD/DOWNLOAD.sh
chmodgrabscript:
	chmod 777 $$PWD/GRAB.sh
chmodrunscript:
	chmod 777 $$PWD/RUNNER.sh
chmodsetupscript:
	chmod 777 $$PWD/SETUP.sh
chmodbuildscript:
	chmod 777 $$PWD/BUILD.sh
deleteunzippedsqldatanotinroot:
	if [ -d "${$$PWD/grabbedsqldata}" ]; then /
        rm -rf ${$$PWD/grabbedsqldata}; /
  fi /
deletezippedsqldatanotinroot:
	if [ -d "${$$PWD/grabbedsqldata}" ]; then /
  	rm -rf ${$$PWD/bkps}; /
  fi /
grabsqldata:
	bash $$PWD/dl.sh
	wall dl.sh was ran!
clearsqldatainroot:
	if [ -d "${$$PWD/license.txt}" ]; then /
        rm -rf ${$$PWD/license.txt}; /
  fi /
		if [ -d "${$$PWD/worldcities.csv}" ]; then /
        rm -rf ${$$PWD/worldcities.csv}; /
  fi /
		if [ -d "${$$PWD/worldcities.xlsx}" ]; then /
        rm -rf ${$$PWD/worldcities.xlsx}; /
  fi /
		if [ -d "${$$PWD/sqldata.zip}" ]; then /
        rm -rf ${$$PWD/sqldata.zip}; /
  fi /
createsqldatafolder:
	mkdir $$PWD/grabbedsqldata
createbackupfolder:
	mkdir $$PWD/bkps
createzipbackupfolder:
	mkdir $$PWD/bkps/zip/
createsqldatazipbackupfolder:
	mkdir $$PWD/bkps/zip/sqldata.zip/
createsqldatanodefolder:
	mkdir $$PWD/bkps/zip/sqldata.zip/000000000000/
moveunzippedsqldatatofolder:
	cp -r $$PWD/license.txt $$PWD/grabbedsqldata/license.txt
	cp -r $$PWD/worldcities.csv $$PWD/grabbedsqldata/worldcities.csv
	cp -r $$PWD/worldcities.xlsx $$PWD/grabbedsqldata/worldcities.xlsx
movezippedsqldatatofolder:
	cp -r $$PWD/sqldata.zip $$PWD/bkps/zip/sqldata.zip/000000000000/sqldata.zip
