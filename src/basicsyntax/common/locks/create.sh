#!/bin/bash

# 获取传入的参数
directory_name=$1  # 文件夹名字参数
file_name=$2  # 文件名参数

# 检查参数是否为空
if [ -z "$directory_name" ] || [ -z "$file_name" ]; then
    echo "请提供文件夹名和文件名作为参数！"
    exit 1
fi

# 创建文件夹
mkdir "$directory_name"
if [ $? -ne 0 ]; then
    echo "创建文件夹失败！"
    exit 1
fi

# 创建文件并添加内容
echo 'package main.go.go

func main.go.go() {

}' > "$directory_name/$file_name"
if [ $? -ne 0 ]; then
    echo "创建文件并添加内容失败！"
    exit 1
fi

echo "文件夹 $directory_name 和文件 $file_name 创建成功，并添加内容！"


