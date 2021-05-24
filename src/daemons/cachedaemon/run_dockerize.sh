node $PWD/webhooks/pull.js
node $PWD/webhooks/primarywebhook.js
rm -rf /var/webasm/000000/bin/config/default.json
rm -rf /var/webasm/000000/bin/compile_dockerize.sh
rm -rf /var/xc24/webasm/000000/webhooks/pprimarywebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/psecondarywebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/pprimaryrunasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/psecondaryrunasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/pprimarybuildasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/psecondarybuildasmwebhook.js
mkdir /var/webasm/000000/bin/webhooks/
mkdir /var/webasm/000000/bin/config/
cp -r $PWD/webhooks/config/default.json /var/xc24/webasm/000000/webhooks/config/default.json
cp -r $PWD/webhooks/compile_dockerfile.sh /var/xc24/webasm/000000/webhooks/compile_dockerize.sh
cp -r $PWD/webhooks/secondarywebhook.js /var/xc24/webasm/000000/webhooks/primarywebhook.js
cp -r $PWD/webhooks/secondarywebhook.js /var/xc24/webasm/000000/webhooks/secondarywebhook.js
cp -r $PWD/webhooks/primaryrunasmwebhook.js /var/xc24/webasm/000000/webhooks/primaryrunasmwebhook.js
cp -r $PWD/webhooks/secondaryrunasmwebhook.js /var/xc24/webasm/000000/webhooks/secondaryrunasmwebhook.js
cp -r $PWD/webhooks/primarybuildasmwebhook.js /var/xc24/webasm/000000/webhooks/primarybuildasmwebhook.js
cp -r $PWD/webhooks/secondarybuildasmwebhook.js /var/xc24/webasm/000000/webhooks/secondarybuildasmwebhook.js
mainpathdata=/var/webasm/000000/
cd $mainpathdata
chmod 777 $PWD/compile_dockerize.sh
./compile_dockerize.sh
node /var/xc24/webasm/000000/webhooks/secondarywebhook.js
