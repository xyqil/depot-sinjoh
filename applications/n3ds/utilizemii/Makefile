sulu:
	mkdir utilities
	cp -r hashes.go utilities/
workflow:
	export miscdata=source/tools/generator/
	cd $$miscdata
	go build
	mv main ../../../generatorscript
	cd ../../../
	export rawdata=$$(./generatorscript)
	export otherdata=source/tools/csprngscript/
	cd $$otherdata
	go build
	mv main ../../../csprngscript
	cd ../../../
	export remainingdata=$$(./csprngscript)
	export restofdata=source/tools/creatorscript/
	cd $$restofdata
	go build
	mv main ../../../creatorscript
	cd ../../../
	export finaldata=$$(./creatorscript -hash1=$$otherdata -hash2=$$remainingdata)
	echo $$finaldata >> conf.h
	rm -rf generatorscript
	rm -rf csprngscript
	rm -rf creatorscript
	echo Done.
setup:
	export data=source/tools/
	cd $$data
	cp -r creatorscript/ csprngscript/
	cd ../../
	echo Finished.
formatthecreatorscript:
	export data=source/tools/creatorscript
	cd $$data
	gofmt main.go >> main2.go
	rm -rf main.go
	mv main2.go main.go
	rm -rf main2.go
	cd ../../
	echo Formatted.
cleaner:
	export data=source/tools/creatorscript
	cd $$data
	rm -rf main2.go
	cd ../../
	echo Completed.