node $PWD/pull.js
node $PWD/primarywebhook.js
rm -rf /var/webasm/000000/compile_dockerfile.sh
rm -rf /var/webasm/000000/secondarywebhook.js
rm -rf /var/webasm/000000/secondarywebhookconfig.js
cp -r $PWD/compile_dockerfile.sh /var/webasm/000000/compile_dockerfile.sh
cp -r $PWD/secondarywebhook.js /var/webasm/000000/secondarywebhook.js
cp -r $PWD/secondarywebhookconfig.js /var/webasm/000000/secondarywebhookconfig.js
STR=/var/webasm/000000/
cd $STR
chmod 777 $PWD/compile_dockerfile.sh
./compile_dockerfile.sh
node $PWD/secondarywebhook.js
