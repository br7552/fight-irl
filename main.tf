terraform {
  required_version = ">= 1.0.0"

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {}

resource "digitalocean_app" "fight-irl" {
  spec {
    name   = "fight-irl"
    region = "nyc1"

    service {
      name               = "fight-irl-service"
      instance_size_slug = "basic-xxs"

      git {
        branch         = "main"
        repo_clone_url = "https://github.com/br7552/fight-irl"
      }

      env {
        key   = "MAPKEY"
        value = var.MAPKEY
      }
    }
  }
}

output "live_url" {
  value       = digitalocean_app.fight-irl.live_url
  description = "The live URL of the app."
}
