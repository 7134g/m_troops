
#go build -ldflags "-s -w" -o ./gen.exe .


#./gen -p gen -t ./template -mongo UserActivity
./gen -p gen -t ./template -mysql article
