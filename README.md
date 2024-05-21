# simple_blockchain

## usecase
- example for simple payment transaction

## leveldb install
- sudo apt update
- sudo apt-get install libsnappy-dev wget curl build-essential cmake gcc sqlite3
- VER=$(curl -s https://api.github.com/repos/google/leveldb/releases/latest | grep tag_name |  cut -d '"' -f 4)
- wget https://github.com/google/leveldb/archive/${VER}.tar.gz -O leveldb.tar.gz
- tar xvf leveldb.tar.gz
- cd leveldb*/
- mkdir -p build && cd build
- cmake -DCMAKE_BUILD_TYPE=Release .. && cmake --build 

## get started
- go mod tidy
- go mod vendor
- import postman_collection.json to postman
- change port if needed
- go run 

## reference
- https://youtu.be/RZ9MjCR4YW8?si=up6ySOnUS4Acq0P2
- https://youtu.be/mYlHT9bB6OE?si=wXWbPm9StCOv1YqK
- https://computingforgeeks.com/how-to-install-leveldb-on-ubuntu-linux/


## note
- not inclue consensus
