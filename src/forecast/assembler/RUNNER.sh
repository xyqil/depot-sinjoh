make executesetupscript --file=$PWD/RUNNER.makefile
make executegrabberscript --file=$PWD/RUNNER.makefile
make chmodit --file=$PWD/SETUP.makefile
make runit --file=$PWD/SETUP.makefile
make announceit --file=$PWD/SETUP.makefile
make logit --file=$PWD/SETUP.makefile
