mkdir /var/webasm/000000/bin/webhooks/
node /var/xc24/webasm/000000/webhooks/pull.js
node /var/xc24/webasm/000000/webhooks/primarywebhook.js
rm -rf /var/webasm/000000/bin/config/default.json
rm -rf /var/webasm/000000/bin/compile_dockerize.sh
rm -rf /var/xc24/webasm/000000/webhooks/primarywebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/secondarywebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/primaryrunasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/secondaryrunasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/primarybuildasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/secondarybuildasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/primarydownloadasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/secondarydownloadasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/primarysetupasmwebhook.js
rm -rf /var/xc24/webasm/000000/webhooks/secondarysetupasmwebhook.js
mkdir /var/webasm/000000/bin/config/
cp -r $PWD/webhooks/config/default.json /var/xc24/webasm/000000/webhooks/config/default.json
cp -r $PWD/webhooks/compile_dockerfile.sh /var/xc24/webasm/000000/webhooks/compile_dockerize.sh
cp -r $PWD/webhooks/secondarywebhook.js /var/xc24/webasm/000000/webhooks/primarywebhook.js
cp -r $PWD/webhooks/secondarywebhook.js /var/xc24/webasm/000000/webhooks/secondarywebhook.js
cp -r $PWD/webhooks/primaryrunasmwebhook.js /var/xc24/webasm/000000/webhooks/primaryrunasmwebhook.js
cp -r $PWD/webhooks/secondaryrunasmwebhook.js /var/xc24/webasm/000000/webhooks/secondaryrunasmwebhook.js
cp -r $PWD/webhooks/primarybuildasmwebhook.js /var/xc24/webasm/000000/webhooks/primarybuildasmwebhook.js
cp -r $PWD/webhooks/secondarybuildasmwebhook.js /var/xc24/webasm/000000/webhooks/secondarybuildasmwebhook.js
cp -r $PWD/webhooks/primarydownloadasmwebhook.js /var/xc24/webasm/000000/webhooks/primarydownloadasmwebhook.js
cp -r $PWD/webhooks/secondarydownloadasmwebhook.js /var/xc24/webasm/000000/webhooks/secondarydownloadasmwebhook.js
cp -r $PWD/webhooks/primarysetupasmwebhook.js /var/xc24/webasm/000000/webhooks/primarysetupasmwebhook.js
cp -r $PWD/webhooks/secondarysetupasmwebhook.js /var/xc24/webasm/000000/webhooks/secondarysetupasmwebhook.js
mainpathdata=/var/webasm/000000/
cd $mainpathdata
chmod 777 $PWD/compile_dockerize.sh
./compile_dockerize.sh
node /var/xc24/webasm/000000/webhooks/secondarywebhook.js
