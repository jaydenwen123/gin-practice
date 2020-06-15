#!/bin/bash
curDir=$(pwd)
echo "[INFO] current directory is:"$curDir
for dir in `ls $curDir`
do
	if [ -d $dir ];then
		cp ./gin_form/fill_sample.sh $dir/
		echo "[INFO] finish copy the gin_form/fill_sample.sh to current directory:"$dir
	fi
done
