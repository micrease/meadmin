#!/bin/bash

SERVER="giftcard-server"
BASE_DIR=$PWD
INTERVAL=1

ARGS=$2
ARGC=$#
echo "Run Args $# $ARGS"

function build() {
  go build -o $SERVER main.go
  echo "$SERVER build success"
}

function start() {
  if [ $ARGC -eq 1 ]; then
    echo "failed,Miss env args,etc:sh server.sh start --env=dev|test|prod"
    exit 1
  fi

  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    echo "$SERVER already running"
    exit 1
  fi

  echo $BASE_DIR/$SERVER $ARGS
  if [ $ARGS = "--env=prod" ]; then
    nohup $BASE_DIR/$SERVER $ARGS > /www/wwwroot/api.hgkwxnfs.com/output.log 2>&1 &
  else
    nohup $BASE_DIR/$SERVER $ARGS > /www/wwwroot/api.vpemjjvb.com/output.log 2>&1 &
  fi
  echo "sleeping..." && sleep $INTERVAL

  # check status
  if [ "$(pgrep $SERVER -u $UID)" == "" ]; then
    echo "$SERVER start failed"
    exit 1
  else
    echo "start success"
  fi
}

function status() {
  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    echo $SERVER is running
  else
    echo $SERVER is not running
  fi
}

function stop() {
  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    kill $(pgrep $SERVER -u $UID)
  fi

  echo "sleeping..." && sleep $INTERVAL

  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    echo "$SERVER stop failed"
    exit 1
  else
    echo "stop success"
  fi
}

function version() {
  $BASE_DIR/$SERVER $ARGS version
}

case "$1" in
'build')
  build
  ;;
'run')
  build && stop && start
  ;;
'start')
  start
  ;;
'stop')
  stop
  ;;
'status')
  status
  ;;
'restart')
  stop && start
  ;;
'version')
  version
  ;;
*)
  echo "usage: $0 {build|run|start|stop|restart|status|version}"
  exit 1
  ;;
esac
