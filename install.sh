#!/usr/bin/env bash
go get github.com/go-sql-driver/mysql
echo "install go mysql OK"
go get github.com/julienschmidt/httprouter
echo "install go httprouter OK"

go install blog
echo "install go blog OK"
