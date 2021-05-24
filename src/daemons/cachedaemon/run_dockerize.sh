node $PWD/webhooks/pull.js
node $PWD/webhooks/primarywebhook.js
rm -rf /var/webasm/000000/bin/config/default.json
rm -rf /var/webasm/000000/bin/compile_dockerize.sh
rm -rf /var/webasm/000000/bin/webhooks/primarywebhook.js
rm -rf /var/webasm/000000/bin/webhooks/secondarywebhook.js
rm -rf /var/webasm/000000/bin/webhooks/primaryrunasmwebhook.js
rm -rf /var/webasm/000000/bin/webhooks/secondaryrunasmwebhook.js
rm -rf /var/webasm/000000/bin/webhooks/primarybuildasmwebhook.js
rm -rf /var/webasm/000000/bin/webhooks/secondarybuildasmwebhook.js
mkdir /var/webasm/000000/bin/webhooks/
mkdir /var/webasm/000000/bin/config/
cp -r $PWD/webhooks/secondarybuildasmwebhook.js /var/webasm/000000/bin/config/default.json
cp -r $PWD/webhooks/compile_dockerfile.sh /var/webasm/000000/bin/compile_dockerize.sh
cp -r $PWD/webhooks/secondarywebhook.js /var/webasm/000000/bin/webhooks/primarywebhook.js
cp -r $PWD/webhooks/secondarywebhook.js /var/webasm/000000/bin/webhooks/secondarywebhook.js
cp -r $PWD/webhooks/primaryrunasmwebhook.js /var/webasm/000000/bin/webhooks/primaryrunasmwebhook.js
cp -r $PWD/webhooks/secondaryrunasmwebhook.js /var/webasm/000000/bin/webhooks/secondaryrunasmwebhook.js
cp -r $PWD/webhooks/primarybuildasmwebhook.js /var/webasm/000000/bin/webhooks/primarybuildasmwebhook.js
cp -r $PWD/webhooks/secondarybuildasmwebhook.js /var/webasm/000000/bin/secondarybuildasmwebhook.js
mainpathdata=/var/webasm/000000/
cd $mainpathdata
chmod 777 $PWD/compile_dockerize.sh
./compile_dockerize.sh
node $PWD/webhooks/secondarywebhook.js
