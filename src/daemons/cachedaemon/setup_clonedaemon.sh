rm -rf $PWD/installdata/
mkdir $PWD/installdata/
curl -fsSL https://get.docker.com -o $PWD/installdata/get-docker.sh
sh $PWD/installdata/get-docker.sh
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt install nodejs
node /var/webasm/000000/bin/webhooks/primaryclonedaemonsetupwebhook.js
chmod 777 $PWD/download_webasm.sh
./download_webasm.sh
node /var/webasm/000000/bin/webhooks/secondaryclonedaemonsetupwebhook.js
