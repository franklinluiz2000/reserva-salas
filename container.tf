# Criação do container para implantação
resource "docker_container" "reserva" {
  name         = "reserva"
  image        = docker_image.reserva.image_id
  ports {
    internal = 8080
    external = 8080
  }

  volumes {
    host_path      = "/home/aluno/reserva-salas/dados"
    container_path = "/app/dados"
  }

  restart = "always"
}
