oc login -u system:admin


#image=${DEVOPSCONSOLE_IMAGE} #THIS NEEDS TO BE CHANGED
image="avni16/devopsconsole-operator:0.1.0" 
output_csv="final.csv" 
namespace="olm" 
dir=$(realpath "$(dirname "${BASH_SOURCE}")/../..")

manifest_path=${dir}/manifests/devopsconsole/devopsconsole.package.yaml 
devopsconsole_version=`cat ${manifest_path} | grep "currentCSV"| cut -d "." -f2- | cut -d "v" -f2 | tr -d '[:space:]'`
devopsconsole_csv=`cat ${manifest_path} | grep "currentCSV" | cut -d ":" -f2 | tr -d '[:space:]'` 
devopsconsole_csv_yaml="${dir}/manifests/devopsconsole/${devopsconsole_version}/${devopsconsole_csv}.clusterserviceversion.yaml"

oc create namespace ${namespace}
oc project ${namespace}
oc apply -f https://raw.githubusercontent.com/operator-framework/operator-lifecycle-manager/master/deploy/upstream/quickstart/olm.yaml

sed -e "s,REPLACE_IMAGE,${image}," -i ${devopsconsole_csv_yaml} 
sed -e "s,placeholder,${namespace}," -i ${devopsconsole_csv_yaml} 
oc apply -f ${devopsconsole_csv_yaml}
oc get crds