#!/usr/bin/env bash

GO_VERSION=1.8.3
GO_INSTALL_FOLDER=/usr/local
GOPATH=/home/vagrant/go

SQL_ROOT_PASSWORD=testpassword
SQL_USER_NAME=booker
SQL_USER_PASSWORD=$SQL_ROOT_PASSWORD
SQL_DB_NAME=booking


apt-get update

apt-get -y install git mercurial

if [ ! -d $GO_INSTALL_FOLDER/go ];then
	echo 'Downloading go'$GO_VERSION'.linux-amd64.tar.gz'
	wget –q https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz 2> /dev/null

	echo 'Unpacking go language'
	tar -C $GO_INSTALL_FOLDER -xzf go$GO_VERSION.linux-amd64.tar.gz
	rm go$GO_VERSION.linux-amd64.tar.gz

	echo 'Setting up Go Environment'
	echo 'export PATH=$PATH:'$GO_INSTALL_FOLDER'/go/bin' >> .profile
	echo 'export GOPATH='$GOPATH >> .profile
	echo 'export PATH=$PATH:$GOPATH/bin' >> .profile
	echo 'sudo chown -R vagrant:vagrant go' >> .profile #TODO: do this better

else
	echo 'Go already installed. Skipping.'
fi


echo 'Getting Go packages'
export PATH=$PATH:$GO_INSTALL_FOLDER/go/bin
export GOPATH=$GOPATH

go get -v github.com/revel/cmd/revel
go get -v github.com/go-gorp/gorp
go get -v github.com/go-sql-driver/mysql


echo 'Installing Database'
debconf-set-selections <<< 'mysql-server mysql-server/root_password password '$SQL_ROOT_PASSWORD
debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password '$SQL_ROOT_PASSWORD
apt-get install -y mysql-server

# From http://www.websightdesigns.com/posts/view/how-to-configure-an-ubuntu-web-server-vm-with-vagrant
if [ ! -f /var/log/dbinstalled ];then
	echo "CREATE USER '"$SQL_USER_NAME"'@'localhost' IDENTIFIED BY '"$SQL_USER_PASSWORD"'" | mysql -uroot -p$SQL_ROOT_PASSWORD
	echo "CREATE DATABASE "$SQL_DB_NAME | mysql -uroot -ptestpassword
	echo "GRANT ALL ON "$SQL_DB_NAME".* TO '"$SQL_USER_NAME"'@'localhost'" | mysql -uroot -p$SQL_ROOT_PASSWORD
	echo "flush privileges" | mysql -uroot -p$SQL_ROOT_PASSWORD
	touch /var/log/dbinstalled
fi
