#!/bin/bash
if [ $1 = create ]
then
  echo Creating Enviroment...
  python3 -m venv sinjohshell
  source ./sinjohshell/bin/activate
  pip install -r requirements.txt
  source ./post_env_create.sh
  deactivate
else
  echo "Run ./sinjohshell.sh create to create venv or run ./sinjohshell/Scripts/activate to launch enviroment. Launching venv is basic stuff :P"
  fi
