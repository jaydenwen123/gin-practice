#!/bin/bash
# fill the cmd line  curl sample 2 the server handler as comment
curlCmd=$1
echo "[INFO] the curl cmd is:"$curlCmd
url=$(echo $curlCmd|awk  '{for(i=1;i<=NF;i++){if (index($i,"http://")>0){print $i}}}')
echo "[INFO] the url is:"$url
uri=${url##*/}
uri=${uri%\?*}
echo "[INFO] the uri:"$uri
#找到对应的handler,然后添加注释
handler=`grep "/"$uri main.go | grep -E "\w+)" -o|grep -E "\w+" -o`
echo "[INFO] handler is:"$handler
lineNumber=`grep -E "func\s+"$handler -n main.go|cut -d ":" -f 1`
echo "[INFO] the define handler line number is:"$lineNumber
#开始动态添加注释
gsed -i   "${lineNumber}i // ${curlCmd}" main.go
echo "[INFO] the url is add to the handler"

