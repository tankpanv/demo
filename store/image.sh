# project="recommend"
# psm="federation.recommend.store"
project="demo"
psm="federation.demo.store"
go mod tidy
if [[ ${project} == "" || ${psm} == "" ]];then
    echo "please fill in project and psm"
    echo "like:"
    echo  'project="recommend"'
    echo  'psm="federation.recommend.strategy"'
    exit 1
fi
# **** build start**** #
bash build.sh
if [[ $? != 0 ]];then
    echo "build fail"
    exit 1
fi
# **** build end**** #
env=$1
mod=$2
if [[ $env != "prod" && $env != "ppe" && $env != "test" ]];then
    echo "\$1 should be prod or ppe or test"
    echo "bash image.sh ppe debug"
    exit 1
fi
last_version=5

service_name=${psm//\./-}
echo $service_name

let v=$last_version+1
version=1.1.${v}
# env=ppe

workspace=`pwd`



if [[ $mod == "debug" ]];then
    cp -rf /programs/frp_0.37.1_linux_amd64/ ./output
fi
cd output

sudo nerdctl -n k8s.io  build  -t $service_name:$env-$version .
#--no-cache
sudo nerdctl -n k8s.io tag $service_name:$env-$version core.harbor.domain/$project/$service_name:$env-$version
sudo nerdctl -n k8s.io tag $service_name:$env-$version core.harbor.domain/$project/$service_name:$env
sudo nerdctl -n k8s.io  push --insecure-registry core.harbor.domain/$project/$service_name:$env-$version
sudo nerdctl -n k8s.io  push --insecure-registry core.harbor.domain/$project/$service_name:$env
# cd /home/ubuntu/workspace/k8s/$project/$projectapi && sudo  bash start.sh bash
file=$workspace/$0
echo $env-$version
sed -i "s/last_version=${last_version}/last_version=${v}/g" $file
echo $file
# cd $workspace
