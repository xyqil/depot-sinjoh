bench:
	go test -bench
test:
	go test
all:
	go build -ldflags="-s -w" -o bin/gamestats
testhashes:
	echo $$GOPATH > tokenfile
	export GOPATH=$$PWD
	go test hashes_test.go
	export GOPATH="$$(cat tokenfile)"
	rm -rf tokenfile
cleanup:
	echo Cleaning up misspelled files...
	echo Please wait...
	if [ -d "utilities" ]; then rm -rf "utilities"; fi;
	if [ -d ".tmpfile" ]; then rm -rf ".tmpfile"; fi;
	if [ -d "tmpfile" ]; then rm -rf "tmpfile"; fi;
	if [ -d "utilities:" ]; then rm -rf "utilities:"; fi;
	if [ -d ".utilities:" ]; then rm -rf ".utilities:"; fi;
	if [ -d ".utility:" ]; then rm -rf ".utility:"; fi;
	if [ -d "utility:" ]; then rm -rf "utility:"; fi;
	if [ -d ".utility" ]; then rm -rf ".utility"; fi;
	if [ -d "utility" ]; then rm -rf "utility"; fi;
	if [ -d "gitignore" ]; then rm -rf ".installkeyfile"; fi;
	if [ -d "installkeyfile" ]; then rm -rf ".installkeyfile"; fi;
	if [ -d "keyfile" ]; then rm -rf ".keyfile"; fi;
	if [ -d "file" ]; then rm -rf "file"; fi;
	if [ -d ".file" ]; then rm -rf ".file"; fi;
	if [ -d "ignore" ]; then rm -rf "ignore"; fi;
	if [ -d ".ignore" ]; then rm -rf ".ignore"; fi;
	if [ -d ".ignor" ]; then rm -rf ".ignor"; fi;
	if [ -d "ignor" ]; then rm -rf "ignor"; fi;
	if [ -d "installkey" ]; then rm -rf "installkey"; fi;
	if [ -d ".installkey" ]; then rm -rf ".installkey"; fi;
	if [ -d "install" ]; then rm -rf "install"; fi;
	if [ -d ".install" ]; then rm -rf ".install"; fi;
	if [ -d "instal" ]; then rm -rf "instal"; fi;
	if [ -d ".instal" ]; then rm -rf ".instal"; fi;
	if [ -d "tokenfile" ]; then rm -rf "tokenfile"; fi;
	if [ -d ".tokenfile" ]; then rm -rf ".tokenfile"; fi;
	if [ -d "tokenfil" ]; then rm -rf "tokenfil"; fi;
	if [ -d ".tokenfil" ]; then rm -rf ".tokenfil"; fi;
	if [ -d "token" ]; then rm -rf "token"; fi;
	if [ -d ".token" ]; then rm -rf ".token"; fi;
	if [ -d "tokn" ]; then rm -rf "tokn"; fi;
	if [ -d ".tokn" ]; then rm -rf ".tokn"; fi;
	if [ -d "tok" ]; then rm -rf "tok"; fi;
	if [ -d ".tok" ]; then rm -rf ".tok"; fi;
respawntype1:
	echo "utilities/" > .gitignore
respawntype2:
	echo "utilities/" >> .gitignore
despawn:
	if [ -d ".gitignore" ]; then rm -rf "gitignore"; fi;
	if [ -d ".installkeyfile" ]; then rm -rf ".installkeyfile"; fi;
clearscreentype1:
	clear
clearscreentype2:
	cls
setupenviroment:
	mkdir utilities/
	cp -r hashes.go utilities/main.go
installtest:
	openssl rand -base64 64 >> .installkeyfile
	export key=$(cat .installkeyfile)
	mv "utilities" "$key"
	echo Module name should be set to $key
	sed 's/#   \"utilities\"/  \"$key\"/' hashes_test.go
	go install $key
	rm -rf .installkeyfile
	go test hashes_test.go
