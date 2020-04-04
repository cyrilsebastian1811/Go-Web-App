# gowebapp
Simple web application written in Go

### Build Docker Image
#### MacOS
docker build -t cyrilsebastian1811/goapp:latest .

#### Linux
sudo docker build -t cyrilsebastian1811/goapp:latest .

### Run Docker Image
#### MacOS
docker run -dt --name goapp -p 8000:8000 cyrilsebastian1811/goapp:latest

#### Linux
sudo docker run -dt --name goapp -p 8000:8000 cyrilsebastian1811/goapp:latest

### Push the docker image to you repository
#### MacOS
docker push cyrilsebastian1811/goapp:latest

#### Linux
sudo docker push cyrilsebastian1811/goapp:latest
