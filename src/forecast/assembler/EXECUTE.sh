echo Forecast is running............
key=$(openssl rand -base64 484)
echo "${PWD%/[^/]*}" >> $PWD/data-$variable.bin
data=`cat data-variable.bin`
cd $data
python3 $PWD/main.py