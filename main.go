package main

import (
	"fmt"
	"os"

	as "github.com/aerospike/aerospike-client-go"
)

// OK

func main() {
	wd, _ := os.Getwd()
	luaPath := fmt.Sprintf("%s/udf/", wd)
	as.SetLuaPath(luaPath)

	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		panic(err)
	}

	task, err := client.RegisterUDFFromFile(nil, fmt.Sprintf("%s/udf.lua", luaPath), "udf.lua", as.LUA)
	if err != nil {
		panic(err)
	}
	<-task.OnComplete()

	foosAsValueArray := []as.Value{}
	for _, val := range []int{2, 3, 4} {
		foosAsValueArray = append(foosAsValueArray, as.NewValue(val))
	}

	statement := as.NewStatement("test", "foo")
	recordSet, err := client.QueryAggregate(
		nil, statement,
		"udf", "myUDF",
		as.NewValue(foosAsValueArray),
	)
	if err != nil {
		panic(err)
	}

	for rec := range recordSet.Records {
		fmt.Printf("%+v\n", rec)
	}
}
