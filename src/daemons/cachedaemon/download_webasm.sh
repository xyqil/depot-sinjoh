node /var/xc24/webasm/000000/webhooks/primarydownloadasmwebhook.js
rm -rf /var/xc24/webasm/000000/bin/webdata/
git clone https://github.com/nemasu/asmttpd /var/xc24/webasm/000000/bin/
cp -r /var/xc24/webddata/000000/ /var/xc24/webasm/000000/bin/webdata
node /var/xc24/webasm/000000/webhooks/secondarydownloadasmwebhook.js
