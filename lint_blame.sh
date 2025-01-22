#!/bin/bash

golangci-lint run ./... --out-format=line-number > .lint_output.txt

# 遍历 .lint_output.txt 中的每一行
while IFS= read -r line; do
    # 从行中提取文件名和行号
    file_name=$(echo "$line" | awk -F ':' '{print $1}')
    line_number=$(echo "$line" | awk -F ':' '{print $2}')

    # 如果文件名和行号存在，则运行 git blame
    if [[ -n "$file_name" && -n "$line_number" ]]; then
        blame_output=$(git blame -L "$line_number","$line_number" "$file_name")
        echo "----------------"
        echo " Git: $blame_output"
        echo "Lint: $line"
    fi
done < .lint_output.txt

# 删除临时文件
rm .lint_output.txt
