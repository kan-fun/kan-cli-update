version=`echo $GITHUB_REF | cut -d "/" -f 3`

go build -v -ldflags "-X main.version=$version" .

wget http://gosspublic.alicdn.com/ossutil/1.6.10/ossutil64
chmod 755 ossutil64
./ossutil64 config -e oss-cn-beijing.aliyuncs.com -i $ACCESSKEY -k $SECRETKEY

./ossutil64 rm -rf oss://kan-bin/linux --include "kan-update_*"
./ossutil64 cp -f ./kan-cli-update "oss://kan-bin/linux/kan-update_$version"

./ossutil64 cp -f ./kan-cli-update "oss://kan-start/linux/kan-update"