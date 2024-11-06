terraform {
  # https://developer.hashicorp.com/terraform/tutorials/docker-get-started/docker-build#:~:text=provider%20source%20documentation
  required_providers {
    docker = {
      source = "kreuzwerker/docker" //registry.terraform.io/kreuzwerker/docker
      version = "~> 3.0.1"
    }
  }
}

# https://developer.hashicorp.com/terraform/tutorials/docker-get-started/docker-build#:~:text=VPC%20IDs.%20Our-,providers%20reference,-documents%20the%20required
provider "docker" {}

resource "docker_network" "monitoring_network" {
  name = "monitoring_network"
}

# Loki Service
# https://registry.terraform.io/providers/cybershard/docker/latest/docs/resources/container
resource "docker_container" "loki" {
  name  = "loki"
  image = "grafana/loki:3.1.0"
  ports {
    internal = 3100
    external = 3100
  }
  networks_advanced {
    name = docker_network.monitoring_network.name
  }
  command = ["-config.file=/etc/loki/local-config.yaml"]
}

# Fluent-bit Service
resource "docker_container" "fluent_bit" {
  name  = "fluent-bit"
  image = "fluent/fluent-bit:3.1.3"
  depends_on = [docker_container.loki]
  volumes {
    host_path = abspath("./logs")
    container_path = "/var/logs"
  }
  volumes {
    host_path = abspath("./fluent-bit.yml")
    container_path = "/fluent-bit/etc/fluent-bit.yml"
  }
  volumes {
    host_path = abspath("./parsers.conf")
    container_path = "/fluent-bit/etc/parsers.conf"
  }
  networks_advanced {
    name = docker_network.monitoring_network.name
  }
  command = ["-c", "/fluent-bit/etc/fluent-bit.yml"]
}

# Grafana Service
resource "docker_container" "grafana" {
  name  = "grafana"
  image = "grafana/grafana:11.1.0"
  depends_on = [docker_container.loki]
  ports {
    internal = 3000
    external = 3000
  }
  env = toset([
    "GF_AUTH_ANONYMOUS_ENABLED=true",
    "GF_AUTH_ANONYMOUS_ORG_ROLE=Admin",
    "GF_AUTH_DISABLE_LOGIN_FORM=true",
    "GF_FEATURE_TOGGLES_ENABLE=traceqlEditor,metricsSummary"
  ])
  volumes {
    host_path = abspath("./grafana/provisioning")
    container_path = "/etc/grafana/provisioning"
  }
  volumes {
    host_path = abspath("./grafana/dashboards")
    container_path = "/var/lib/grafana/dashboards"
  }
  networks_advanced {
    name = docker_network.monitoring_network.name
  }
}

# Tempo Service
# Memcached Service
resource "docker_container" "memcached" {
  name  = "memcached"
  image = "memcached:1.6.29"
  ports {
    internal = 11211
    external = 11211
  }
  env = toset([
    "MEMCACHED_MAX_MEMORY=64m",
    "MEMCACHED_THREADS=4"
  ])
  networks_advanced {
    name = docker_network.monitoring_network.name
  }
}

# Tempo Service
resource "docker_container" "tempo" {
  name  = "tempo"
  image = "grafana/tempo:latest"
  depends_on = [docker_container.memcached]
  command = ["-config.file=/etc/tempo.yaml"]
  ports {
    internal = 14268
    external = 14268
  }
  ports {
    internal = 3200
    external = 3200
  }
  ports {
    internal = 9095
    external = 9095
  }
  ports {
    internal = 4317
    external = 4317
  }
  ports {
    internal = 4318
    external = 4318
  }
  ports {
    internal = 9411
    external = 9411
  }
  volumes {
    host_path = abspath("./tempo.yaml")
    container_path = "/etc/tempo.yaml"
  }
  volumes {
    host_path = abspath("./tempo-data")
    container_path = "/var/tempo"
  }
  networks_advanced {
    name = docker_network.monitoring_network.name
  }
}
