provider "google" {
  credentials = file("./credentials.json") 
  project     = "suse-level3-4170"  
  region      = "asia-south1"  
}

resource "google_compute_network" "vpc_network" {
  name                    = "suse3-network"
  auto_create_subnetworks = "true"
}

resource "google_compute_instance" "vm_instance" {
  name         = "suse3-vm"
  machine_type = "e2-micro"
  zone         = "asia-south1-c"  

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name
  }
}
