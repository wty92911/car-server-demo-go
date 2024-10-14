service_name=car_server
bin=$(cd "$(dirname "$0")";pwd)
cd ${bin} || exit
rm -rf ../log/
mkdir -p ../log/
GODEBUG=gctrace=1 nohup ./${service_name} \
  >> ../log/stdout_${service_name} 2>>../log/stderr_${service_name} &