export GOPATH=/home/pi/go/
CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o main .
docker build . -t door 
sudo docker run --name door --restart unless-stopped --privileged -d door 
