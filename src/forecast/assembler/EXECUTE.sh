make executor --file=RUNNER.makefile
data=`cat data-variable.bin`
cd $data
python3 $PWD/main.py
