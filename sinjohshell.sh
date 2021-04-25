#!/bin/bash
if [ $1 = create ]
then
  echo Creating Enviroment...
  python3 -m venv sinjohshell
  ./sinjohshell/Scripts/activate
  pip install -r requirements.txt
  chmod +x post_env_create.sh
  ./post_env_create.sh
  ls sinjohshell/Scripts
  ./sinjohshell/Scripts/deactivate
else
  echo "Run ./sinjohshell.sh create to create venv or run ./sinjohshell/Scripts/activate to launch enviroment. Launching venv is basic stuff :P"
  fi
  
