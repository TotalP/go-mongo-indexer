package main

import (
	"fmt"
	"go-mongo-indexer/container"
)

func main() {
	var fIDataSourceDao container.FIDataSourceDao
	fIDataSourceDao.SetUp()
	printFIDataSourceDto(fIDataSourceDao.FetchAll())
	defer fIDataSourceDao.Shutdown()
}

func printFIDataSourceDto(dtoSlice []container.FIDataSourceDto) {
	for _, value := range dtoSlice {
		fmt.Printf("%+v\n", value)
	}
}
