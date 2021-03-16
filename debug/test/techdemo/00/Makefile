buildmii:
	if [ -d "concepttester00/" ]; then rm -rf concepttester00/; fi
	if [ -d "concepttester00/concepttester00" ]; then rm -rf concepttester00; fi
	if [ -d "concepttester00/00" ]; then rm -rf concepttester00; fi
	if [ -d "concepttester00" ]; then rm -rf concepttester00; fi
	if [ -d "00" ]; then rm -rf concepttester00; fi
	mkdir concepttester00; cp -r main.go concepttester00/main.go
	cd concepttester00
	go build
	mv concepttester00 ../
	cd ../
	rm 00
nuke:
	mv main.go /home/bkp
run:
	./concepttester00
	rm ./concepttester00
testwithnomove:
	go build
	./00
	rm 00
testwithmove:
	cd concepttester00
	./concepttester00
cleanup:
	rm -rf concepttester00
erase:
	rmdir concepttester00
remover:
	rm -rf tmp/*
clearscreentype1:
	clear
clearscreentype2:
	cls
goback:
 cd ../
intoroot:
 cd /