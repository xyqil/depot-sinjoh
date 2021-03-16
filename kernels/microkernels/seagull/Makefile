builder:
	clang++-7 -pthread -std=c++17 -o main main.cpp
runner:
	./main
workflow:
	make builder
	make runner
sendthebinaryfilesomewhere:
	export data=source/tools/binarycopierandsendertool/source/
	cd $$data
	go build
	mv binarycopierandsendertool ../../../../../binarycopierandsendertool
	cd ../../../../../
	export other=source/tools/slashhexbytereturnertool/source/
	cd $$other
	go build
	mv slashhexbytereturnertool ../../../../../slashhexbytereturnertool
	cd ../../../../../
	export misc=source/tools/auditableconfigpayloaddataprinter/source/
	cd $$misc
	go build
	mv auditableconfigpayloaddataprinter ../../../../../auditableconfigpayloaddataprinter
	cd ../../../../../
	export remaining=$(./slashhexbytereturnertool -modedata=1)
	export restofdata=$(./auditableconfigpayloaddataprinter)
	./binarycopierandsendertool -outputpathdata=$$(outputvariabledata) -filenamedata=$$(filenamevariabledata) -miscdata=$$restofdata -symboldata==$$remaining
	rm -rf binarycopierandsendertool
	rm -rf slashhexbytereturnertool
	rm -rf auditableconfigpayloaddatprinter
	