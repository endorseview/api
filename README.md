# api
# run this commands 

docker image build -t endorseview:1.0 .
docker container run --publish 8000:8080 --detach --name bb endorseview:1.0

go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/badoux/checkmail
go get -v github.com/dgrijalva/jwt-go
go get -v github.com/dgrijalva/jwt-go/request
go get -v github.com/jinzhu/gorm
go get -v golang.org/x/crypto/bcrypt
go get -v github.com/joho/godotenv
go get -v github.com/gorilla/mux    
go get -v github.com/lib/pq
go get -v github.com/qor/audited

docker images --> get image list
docker rmi image_id  -->delete image
docker run -t -i 25c4671a1478  --> go to console of image
docker ps 
docker network create mynet
docker run --name golang --net mynet golang

# to run docker 
docker-compose up --build