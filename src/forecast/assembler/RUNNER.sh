bash $PWD/SETUP.sh
bash $PWD/GRAB.sh
make chmodit --file=$PWD/SETUP.makefile
make runit --file=$PWD/SETUP.makefile
make announceit --file=$PWD/SETUP.makefile
make logit --file=$PWD/SETUP.makefile