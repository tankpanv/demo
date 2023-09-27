env=$1
version=$2
is_create=$3
#project="stock"
#psm="wealth.stock.mainstore"
project="demo"
psm="federation.demo.store"
if [[ $project == "" || $psm == "" ]];then
	echo "please fill in project and psm"
	exit 1
fi
service_name=${psm//\./-}
deployment_service_name=${service_name}-${env}
if [[ $env != "ppe" && $env != "prod" && $env != "test" ]];then
	echo "env = ppe or prod or test"
	echo "like:"
	echo "bash update.sh ppe 1.1.1"
fi

env_type=${envn%%_*} 
if [[ $env_type == "" ]];then
	env_type=$env
fi
if [[ $version == "" ]];then
	echo "version is emtpy"
	exit 1
fi
deployment_file=deployment_$env.yaml
destination_rule_file=destination_rule.yaml
virtual_server=virtual_server.yaml 
image="core.harbor.domain\/$project\/$service_name:$env-$version"
if [[ ${is_create} == "create" ]];then
	cp template_deployment.yaml $deployment_file
	sed -i "s/{IMAGE}/${image}/g" $deployment_file
fi
if [[ ! -f $deployment_file ]];then
	echo "file not exist ,please use create command"
	echo "bash update.sh prod 1.1.1 create"
	exit
fi
if [[ ! -f $destination_rule_file ]];then
	cp template_destination_rule.yaml $destination_rule_file
fi
if [[ ! -f $virtual_server ]];then
	cp template_virtual_server.yaml $virtual_server
	sed -i "s/{SERVICE_NAME}/${service_name}/g" $virtual_server
fi
login_file="login.sh"
if [[ ! -f $login_file ]];then
	cp template_login.sh $login_file
	sed -i "s/{SERVICE_NAME}/${service_name}/g" $login_file
fi
echo $image
sed -i "s/image: core.harbor.domain\/${project}\/${service_name}:${env}-.*/image: core.harbor.domain\/${project}\/${service_name}:${env}-${version}/g" $deployment_file
sed -i "s/{SERVICE_NAME}/${service_name}/g" $deployment_file
sed -i "s/{DEPLOYMENT_SERVICE_NAME}/${deployment_service_name}/g" $deployment_file
sed -i "s/{ENV}/${env}/g" $deployment_file
sed -i "s/{ENV_TYPE}/${env_type}/g" $deployment_file

# destination_rule
sed -i "s/{SERVICE_NAME}/${service_name}/g" $destination_rule_file
if [[  $(sed -n "/version: ${env}/p" ${destination_rule_file}) == '' ]];then
	sed -i "/- name: test/i\\    - name: ${env}" ${destination_rule_file}
	sed -i "/- name: test/i\\      labels:" ${destination_rule_file}
	sed -i "/- name: test/i\\        version: ${env}" ${destination_rule_file}
	kubectl apply -f $destination_rule_file
	kubectl apply -f $virtual_server
fi
if [[ ${is_create} == "create" ]];then
	kubectl apply -f $destination_rule_file
	kubectl apply -f $virtual_server
fi
kubectl apply -f $deployment_file
kubectl -n default rollout status deployment ${deployment_service_name}
