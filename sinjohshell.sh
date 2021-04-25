#!/bin/bash
if [ $1 = create ]
then
  echo Creating Enviroment...
  python3 -m venv sinjohshell
  ls sinjohshell
  ./sinjohshell/bin/activate
  pip install -r requirements.txt
  chmod +x post_env_create.sh
  ./post_env_create.sh
  ./sinjohshell/bin/deactivate
else
  echo "Run ./sinjohshell.sh create to create venv or run ./sinjohshell/Scripts/activate to launch enviroment. Launching venv is basic stuff :P"
  fi
  
