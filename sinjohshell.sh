#!/bin/bash
echo "_____ _       _       _       _____ _          _ _ 
/  ___(_)     (_)     | |     /  ___| |        | | |
\ `--. _ _ __  _  ___ | |__   \ `--.| |__   ___| | |
 `--. \ | '_ \| |/ _ \| '_ \   `--. \ '_ \ / _ \ | |
/\__/ / | | | | | (_) | | | | /\__/ / | | |  __/ | |
\____/|_|_| |_| |\___/|_| |_| \____/|_| |_|\___|_|_|
             _/ |                                   
            |__/\n\n\n"
if [ $1 = create ]
then
  echo Creating Enviroment...
  python3 -m venv sinjohshell
  ./sinjohshell/Scripts/activate
  pip install -r requirements.txt
  ./post_env_create.sh
  ./sinjohshell/Scripts/deactivate
else
  echo "Run ./sinjohshell.sh create to create venv or run ./sinjohshell/Scripts/activate to launch enviroment. Launching venv is basic stuff :P"
  
