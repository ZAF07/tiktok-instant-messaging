package main

import (
	"fmt"

	"github.com/ZAF07/tiktok-instant-messaging/rpc-server/config"
)

func main() {
	rp := config.InitAppConfig()
	fmt.Println("RPC PORT: ", rp.GetPortRPC())
	fmt.Println("RPC METHOD: ", rp.GetMethodRPC())
	fmt.Println("RPC GRACE: ", rp.GetMaxConGraceRPC())
	fmt.Println("RPC AGE: ", rp.GetMaxConAgeRPC())
	fmt.Println("RPC DATASTORE NAME: ", rp.GetDatastoreName())
	fmt.Println("RPC DATASTORE USER: ", rp.GetDatastoreUser())
	fmt.Println("RPC DATASTORE PASSWORD: ", rp.GetDatastorePassword())
	fmt.Println("RPC DATASTORE HOST: ", rp.GetDatastoreHost())
	fmt.Println("RPC DATASTORE ENVIRONMENT: ", rp.GetDatastoreEnv())
	fmt.Println("RPC DATASTORE SSL: ", rp.GetDatastoreSSL())
	fmt.Println("RPC DATASTORE MAX CONN: ", rp.GetDatastoreMaxCons())
	fmt.Println("RPC DATASTORE PORT: ", rp.GetDatastorePort())

}
