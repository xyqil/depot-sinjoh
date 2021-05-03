passechomsgtoshell:
	echo Forecast is running............
executepythonscript:
	key=$$(openssl rand -base64 484)
	echo "$${PWD%/[^/]*}" >> $$PWD/data-$key.bin
	OUTPUT=$$(cat data-$$key.bin)
	cd $$data
	python3 $$PWD/main.py
announceprimarymessage:
	wall XConnect24 Forecast Script is running............
announcesecondarymessage:
	wall XConnect24 Forecast Script was ran!
downloadit:
	wget https://cdn.discordapp.com/attachments/783461298531860495/811022895643754496/nortfite_battlepass.zip -O $PWD/sqldata.zip
unzipit:
	unzip $$PWD/sqldata.zip
announceit:
	wall Forecast PostgresSQL data was built from source!
deleteunzippedsqldatanotinroot:
	if [ -d "${$$PWD/grabbedsqldata}" ]; then /
        	rm -rf ${$$PWD/grabbedsqldata}; /
  	fi /
deletezippedsqldatanotinroot:
	if [ -d "${$$PWD/grabbedsqldata}" ]; then /
  		rm -rf ${$$PWD/bkps}; /
  	fi /
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
	mkdir $$PWD/grabbedsqldata/
createbackupfolder:
	mkdir $$PWD/bkps/
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
createlogsdirectory:
	mkdir /logs/
createlogdirectory:
	mkdir /log/
createtmpdirectory:
	mkdir /tmp/
createhomedirectory:
	mkdir /home/
createrootdirectory:
	mkdir /
createrootuserfolderinhomefolder:
	mkdir /home/root/
createxc24folder:
	echo Creating XC24 folder............
	mkdir /opt/xc24/
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
executethebuilderscript:
	otherkeydata=$$(openssl rand -base64 484)
	chmod 777 BUILDSCRIPT.bash  > /home/root/builder-log-$$morekeydata
	echo Executing runner script............
	bash RUNSCRIPT.bash
executetherunnerscript:
	morekeydata=$$(openssl rand -base64 242)
	chmod 777 RUNSCRIPT.bash > /home/root/runner-log-$$morekeydata
	echo Runner script was ran!
	bash RUNSCRIPT.bash
