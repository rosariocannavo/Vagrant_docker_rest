Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.network "private_network", type: "dhcp"
  config.vm.synced_folder "./vagrant/app_js", "/vagrant/app_js"
  config.vm.synced_folder "./vagrant/app_go", "/vagrant/app_go"
  config.vm.network "forwarded_port", guest: 8080, host: 8080       #vm to host port forwarding
  config.vm.network "forwarded_port", guest: 9090, host: 9090
  config.vm.provider "virtualbox" do |vb|
    
    vb.memory = "1024"
    vb.cpus = 2
  end

config.vm.provision "docker" do |docker|
    # First container
    docker.build_image "/vagrant/app_js", args: "-t node_image"
    docker.run "node_image", args: "--name node_container -d -p 8080:8080"  #container to vm port forwarding
    
    # Second container
    docker.build_image "/vagrant/app_go", args: "-t go_image"
    docker.run "go_image", args: "--name go_container -d -p 9090:9090"
  end

end