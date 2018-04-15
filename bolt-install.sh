cd ~
rm -rf bolt

git clone https://github.com/hervit0/bolt.git

while [ ! -d bolt ]
do
  sleep 1
done

./bolt/install
# git checkout develop
./clean_up
