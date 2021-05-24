stringdata="/var/xc24/webasm/000000/bin/"
node /var/xc24/webasm/000000/webhooks/primarydownloadasmwebhook.js
rm -rf /var/xc24/webasm/000000/bin/webdata/
git clone https://github.com/nemasu/asmttpd /var/xc24/webasm/000000/bin/
cp -r $PWD/build_webasm.sh /var/xc24/webasm/000000/bin/build_webasm.sh
cd $stringdata
chmod 777 $PWD/build_webasm.sh
./build_webasm.sh
cp -r /var/xc24/webddata/000000/ /var/xc24/webasm/000000/bin/webdata
node /var/xc24/webasm/000000/webhooks/secondarydownloadasmwebhook.js
