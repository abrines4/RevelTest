#!/usr/bin/env bash

GO_VERSION=1.8.3
GO_INSTALL_FOLDER=/usr/local

GOPATH=/home/vagrant/go

apt-get update

apt-get -y install git mercurial

if [ ! -d $GO_INSTALL_FOLDER/go ];then
	echo 'Downloading go$GO_VERSION.linux-amd64.tar.gz'
	wget –q https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz

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

echo 'Installing Database'
debconf-set-selections <<< 'mysql-server mysql-server/root_password password testpassword'
debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password testpassword'
apt-get update
apt-get install -y mysql-server
