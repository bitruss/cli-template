#!/bin/bash

#install|remove|start|stop|restart|status
command=$2

#service name
service_name=$1

#dir path
project_path=$(
    cd $(dirname $0)
    pwd
)




function install {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start {
    echo "start..."
    sudo service $service_name start
}

function stop {
    echo "stop..."
    sudo service $service_name stop
}

function restart {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status {
    echo "status"
    sudo service $service_name status
}

function test {
    echo "project_path:"$project_path
    echo "service_name:"$service_name
    echo "exe_path:"$project_path/$service_name
}

case "$command" in
install)
    install
    ;;

remove)
    remove
    ;;
start)
    start
    ;;
stop)
    stop
    ;;
restart)
    restart
    ;;
status)
    status
    ;;
test)
    test
    ;;
*)
    echo "Usage: sudo $0 {install|remove|start|stop|restart|status}"
    ;;
esac
