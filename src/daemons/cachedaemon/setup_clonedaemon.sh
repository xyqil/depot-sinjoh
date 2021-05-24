rm -rf $PWD/installdata/
mkdir $PWD/installdata/
curl -fsSL https://get.docker.com -o $PWD/installdata/get-docker.sh
sh $PWD/installdata/get-docker.sh
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt install nodejs
node $PWD/webhooks/primaryclonedaemonsetupwebhook.js
chmod 777 build_webasm.sh
./build_webasm.sh
chmod 777 download_webasm.sh
./download_webasm.sh
node $PWD/webhooks/secondaryclonedaemonsetupwebhook.js
