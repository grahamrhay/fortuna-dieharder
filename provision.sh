sudo apt-get update -q
sudo apt-get upgrade -qy

curl -OL https://go.googlecode.com/files/go1.2.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.2.1.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.profile
echo "export GOPATH=/vagrant" >> /home/vagrant/.profile

sudo apt-get install git dieharder -qy
