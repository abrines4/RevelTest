#!/usr/bin/env bash

GO_VERSION=1.8.3
GO_INSTALL_FOLDER=/usr/local

GOPATH=/home/vagrant/go


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
debconf-set-selections <<< 'mysql-server mysql-server/root_password password testpassword'
debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password testpassword'
apt-get install -y mysql-server

# From http://www.websightdesigns.com/posts/view/how-to-configure-an-ubuntu-web-server-vm-with-vagrant
if [ ! -f /var/lof/dbinstalled ];then
	echo "CREATE USER 'booker'@'localhost' IDENTIFIED BY 'testpassword'" | mysql -uroot -ptestpassword
	echo "CREATE DATABASE booking" | mysql -uroot -ptestpassword
	echo "GRANT ALL ON booking.* TO 'booker'@'localhost'" | mysql -uroot -ptestpassword
	echo "flush privileges" | mysql -uroot -ptestpassword
	touch /var/log/dbinstalled
fi
