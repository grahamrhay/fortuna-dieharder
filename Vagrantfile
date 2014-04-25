# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "debian7"
  config.vm.provision "shell", path: "provision.sh"
  config.vm.provider "virtualbox" do |v|
      v.memory = 1024
      v.cpus = 2
  end
end
