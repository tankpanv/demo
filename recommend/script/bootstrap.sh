#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi
BinaryName=$(ls $CURDIR/bin/)
# echo "192.168.1.114  code.huanfangsk.com" >> /etc/hosts
execDir="."
nohup ${execDir}/bin/${BinaryName}  > $KITEX_LOG_DIR/app/run.log 2>&1 &
sleep 30
pid=`ps -ef|grep ${execDir}/bin/${BinaryName}|grep -v "grep"| awk '{print $2}'`
echo "${execDir}/bin/${BinaryName} $pid"
pwd
netstat -nlp|grep 88
sed -i 's/kernel.yama.ptrace_scope = 1/kernel.yama.ptrace_scope = 0/g'  /etc/sysctl.d/10-ptrace.conf
# dlv --listen=:2345 --headless=true --api-version=2  --accept-multiclient  attach ${pid}
touch  nohup.out
tail -f nohup.out

