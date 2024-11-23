provider "aws" {
    region = "us-west-2"
}

resource "aws_key_pair" "keys" {
 key_name = ""
 public_key = ""
}

resource "aws_security_group" "allow_ssh" {
    name = "allow_ssh"
    description = "Allow ssh inbound traffic"
    
    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = [ "0.0.0.0/0" ]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = [ "0.0.0.0/0" ]
    }
}

resource "aws_security_group" "allow_http" {
    name = "allow_http "
    description = "Allow http inbound traffic   "
    
    ingress {
        from_port = 8000
        to_port = 8000
        protocol = "tcp"
        cidr_blocks = [ "0.0.0.0/0" ]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "tcp"
        cidr_blocks = [ "0.0.0.0/0" ]
    }
}
resource "aws_security_group" "allow_mysql" {
  name        = "allow_mysql"
  description = "Allow MySQL inbound traffic"
  
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]  
  } 
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "go_api_instance" {
  ami = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  key_name = aws_key_pair.keys.key_name

  tags = {
    Name = "Go API SERVER"
  }

  user_data = <<-EOF
        #!/bin/bash
        # Update the system and install necessary dependencies using dnf
        dnf update -y

        # Install Docker and Docker Compose
        dnf install -y docker
        dnf install -y dnf-plugins-core

        # Install Docker Compose (from the Docker Compose repository)
        dnf config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
        dnf install -y docker-compose

        # Enable Docker to start on boot and start Docker service
        systemctl enable docker
        systemctl start docker

        # Clone your repository using a DNS URL (replace with your actual Git repository or DNS-resolved URL)
        git clone https://your-dns-name-or-repo-url.com/yourusername/yourrepo.git /home/ubuntu/yourrepo

        # Navigate to your project directory
        cd /home/ubuntu/yourrepo

        # Start the application with Docker Compose
        docker-compose -f /home/ubuntu/yourrepo/docker-compose.yml up -d

  EOF

  security_groups = [ 
    aws_security_group.allow_ssh,
    aws_security_group.allow_http,
    aws_security_group.allow_mysql    
   ]
}

output "instance_ip" {
 value = aws_instance.go_api_instance.public_ip
}

