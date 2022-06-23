#!/bin/bash

#install|remove|start|stop|restart|status
command=$1

#dir path
project_path=$(
    cd $(dirname $0)
    pwd
)

#sh file name
project_name="${project_path##*/}"


function install {

    if [ -f /etc/systemd/system/$project_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$project_name.service <<EOF
[Unit]
Description=$project_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$project_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $project_name.service
    fi
}

function remove {
    stop
    echo "remove..."
    sudo systemctl disable $project_name.service
    sudo rm -f /etc/systemd/system/$project_name.service
    sudo systemctl daemon-reload
}

function start {
    echo "start..."
    sudo service $project_name start
}

function stop {
    echo "stop..."
    sudo service $project_name stop
}

function restart {
    echo "Restarting server.."
    sudo service $project_name restart

}

function status {
    echo "status"
    sudo service $project_name status
}

function test {
    echo $project_path
    echo $project_name
    echo $project_path/$project_name
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
