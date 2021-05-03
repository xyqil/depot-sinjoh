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