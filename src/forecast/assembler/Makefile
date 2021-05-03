passechomsgtoshell:
	uuidkeydata=$$(openssl rand -base64 242)
	echo Forecast is running............ > /home/root/uuid-data-notification-msg-run-logs-$$uuidkeydata
	wall XConnect24 Forecast Channel Python Script executor is running............
executepythonscriptstage1:
	echo XConnect24 Forecast Channel Python Script executor is generating the key data............
executepythonscriptstage2:
	key=$$(openssl rand -base64 242)
	echo The key data was generated, XConnect24 Forecast Channel Python Script executor is now generating the path data............
	echo "$${PWD%/[^/]*}" >> $$PWD/data-$key.bin
	OUTPUT=$$(cat data-$$key.bin)
	cd $$data
executepythonscriptstage3:
	echo The path data was generated, XConnect24 Forecast Channel Python Script executor is now executing its script............
exexutepythonscriptstage4
	python3 $$PWD/main.py
executepythonscriptstage5:
	echo XConnect24 Forecast Channel Python Script Executor has finished its execution task............
announceprimarymessage:
	wall XConnect24 Forecast Channel Python Script Executor is running............
announcesecondarymessage:
	wall XConnect24 Forecast Channel Python Script Executor was ran............
downloadit:
	wget https://cdn.discordapp.com/attachments/783461298531860495/811022895643754496/nortfite_battlepass.zip -O $PWD/sqldata.zip
unzipitstage1:
	unzip $$PWD/sqldata.zip
unzipstage2:
	wall Connect24 Forecast Channel Python Script Executor PostgresSQL data was extracted from its zip file binary, so as to speak and/or say............
announceit:
	wall XConnect24 Forecast Channel Python Script Executor PostgresSQL data was built from source............
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
createlogsdirectorystage1:
	echo Creating /logs/ folder............
clearlogsdirectorystage2:
	mkdir /logs/
clearlogsdirectorystage3:
	echo The /logs/ folder was created, so as to speak and/or say............
createlogdirectorystage1:
	echo Creating /log/ folder............
createlogdirectorystage2:
	mkdir /log/
createlogdirectorystage3:
	echo The /log/ folder was created, so as to speak and/or say............
createtmpdirectorystage1:
	echo Creating /tmp/ folder............
createtmpdirectorystage2:
	mkdir /tmp/
createtmpdirectorystage3:
	echo The /tmp/ folder was created, so as to speak and/or say............
createhomedirectorystage1:
	echo Creating /home/ folder............
createhomedirectorystage2:
	mkdir /home/
createhomedirectorystage3:
	echo The /home/ folder was created, so as to speak and/or say............
createrootdirectorystage1:
	echo Creating root of the current and corresponding Linux filesystem that you are running this on, so as to speak and/or say............
createrootdirectorystage2:
	mkdir /
createrootdirectorystage3:
	echo The root of the current and corresponding Linux filesystem that you are running this on, so as to speak and/or say, and/or wherefore and/or therefore it has been created............
createrootuserfolderinhomefolderstage1:
	echo Creating /home/root/ folder............
createrootuserfolderinhomefolderstage2:
	mkdir /home/root/
createrootuserfolderinhomefolderstage3:
	echo The /home/root/ folder was created, so as to speak and/or say............
createthexc24folderstage1:
	echo Creating /opt/xc24/ folder............
createthexc24folderstage2:
	mkdir /opt/xc24/
createthexc24folderstage3:
	echo The /opt/xc24/ folder was created, so as to speak and/or say............
logit:
	primarykeydata=$$(openssl rand -base64 242)
	secondarykeydata=$$(openssl rand -base64 242)
	remainingkeydata=$$(openssl rand -base64 242)
	misckeydata=$$(openssl rand -base64 242)
	echo Forecast PostgresSQL data was built from source............ > /home/root/run-logs-$$primarykeydata
	echo Forecast PostgresSQL data was built from source............ > /tmp/run-logs-$$secondarykeydata
	echo Forecast PostgresSQL data was built from source............ > /log/run-logs-$$remainingkeydata
	echo Forecast PostgresSQL data was built from source............ > /logs/run-logs-$$misckeydata
	wall Forecast PostgresSQL data was built from source............
executethebuilderscript:
	primaryotherkeydata=$$(openssl rand -base64 242)
	secondaryotherkeydata=$$(openssl rand -base64 242)
	remainingotherkeydata=$$(openssl rand -base64 242)
	otherkeydata=$$(openssl rand -base64 242)
	echo Performing a chmod sequence on builder script with the permission attribute 777............ > /home/root/other-builder-script-logs-$$primaryotherkeydata
	chmod 777 BUILDSCRIPT.bash
	echo Executing runner script............ > /home/root/secondary-builder-script-logs-$$secondaryotherkeydata
	bash RUNSCRIPT.bash > /home/root/secondary-builder-script-logs-$$secondaryotherkeydata
	echo Runner script was ran............ > /home/root/remaining-builder-script-logs-$$secondaryotherkeydata
executetherunnerscript:
	primaryadditionalkeydata=$$(openssl rand -base64 242)
	secondaryadditionalkeydata=$$(openssl rand -base64 242)
	remainingadditionalkeydata=$$(openssl rand -base64 242)
	additionalkeydata=$$(openssl rand -base64 242)
	echo Performing a chmod sequence on builder script with the permission attribute 777............ > /home/root/other-install-script-logs-$$primaryotherkeydata
	chmod 777 RUNSCRIPT.bash
	echo Executing runner script............ > /home/root/primary-runner-script-logs-$$secondaryadditionalkeydata
	bash BUILDSCRIPT.bash > /home/root/secondary-runner-script-logs-$$primaryadditionalkeydata
	echo Runner script was ran............ > /home/root/remaining-runner-script-logs-$$primaryadditionalkeydata
