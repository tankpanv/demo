CURDIR=./
BinaryName=$(ls $CURDIR/bin/)
execDir="."
pid=`ps -ef|grep ${execDir}/bin/${BinaryName}|grep -v "grep"| awk '{print $2}'`
echo "${execDir}/bin/${BinaryName} $pid"
nohup dlv --listen=:2345 --headless=true --api-version=2  --accept-multiclient  attach ${pid} &
 ./frp_0.37.1_linux_amd64/frpc -c  ./frp_0.37.1_linux_amd64/frpc.ini
