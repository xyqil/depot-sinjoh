node $PWD/webhooks/pull.js
node $PWD/webhooks/primarywebhook.js
rm -rf /var/webasm/000000/bin/compile_dockerize.sh
rm -rf /var/webasm/000000/bin/secondarywebhook.js
rm -rf /var/webasm/000000/bin/secondarywebhookconfig.js
rm -rf /var/webasm/000000/bin/primaryrunasmwebhookconfig.js
rm -rf /var/webasm/000000/bin/secondaryrunasmwebhookconfig.js
rm -rf /var/webasm/000000/bin/primarybuildasmwebhookconfig.js
rm -rf /var/webasm/000000/bin/secondarybuildasmwebhookconfig.js
rm -rf /var/webasm/000000/bin/primaryrunasmwebhook.js
rm -rf /var/webasm/000000/bin/secondaryrunasmwebhook.js
rm -rf /var/webasm/000000/bin/primarybuildasmwebhook.js
rm -rf /var/webasm/000000/bin/secondarybuildasmwebhook.js
mkdir /var/webasm/000000/bin/webhooks/
cp -r $PWD/webhooks/compile_dockerfile.sh /var/webasm/000000/bin/compile_dockerize.sh
cp -r $PWD/webhooks/secondarywebhook.js /var/webasm/000000/bin/webhooks/secondarywebhook.js
cp -r $PWD/webhooks/secondarywebhookconfig.js /var/webasm/000000/bin/webhooks/secondarywebhookconfig.js
cp -r $PWD/webhooks/primaryrunasmwebhookconfig.js /var/webasm/000000/bin/webhooks/primaryrunasmwebhookconfig.js
cp -r $PWD/webhooks/secondaryrunasmwebhookconfig.js /var/webasm/000000/bin/webhooks/secondaryrunasmwebhookconfig.js
cp -r $PWD/webhooks/primarybuildasmwebhookconfig.js /var/webasm/000000/bin/webhooks/primarybuildasmwebhookconfig.js
cp -r $PWD/webhooks/secondarybuildasmwebhookconfig.js /var/webasm/000000/bin/webhooks/secondarybuildasmwebhookconfig.js
cp -r $PWD/webhooks/primaryrunasmwebhook.js /var/webasm/000000/bin/webhooks/primaryrunasmwebhook.js
cp -r $PWD/webhooks/secondaryrunasmwebhook.js /var/webasm/000000/bin/webhooks/secondaryrunasmwebhook.js
cp -r $PWD/webhooks/primarybuildasmwebhook.js /var/webasm/000000/bin/webhooks/primarybuildasmwebhook.js
cp -r $PWD/webhooks/secondarybuildasmwebhook.js /var/webasm/000000/bin/secondarybuildasmwebhook.js
mainpathdata=/var/webasm/000000/
cd $mainpathdata
chmod 777 $PWD/compile_dockerize.sh
./compile_dockerize.sh
node $PWD/secondarywebhook.js
