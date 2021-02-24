package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	//Create an OSSClient instance.
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4G1ymHvHH8pUe1HmHFSL", "SZYRvjdKzjJIfcLlDbFMAJTGYDz38P")
	if err != nil {
		fmt.Println("Error1:", err)
		os.Exit(-1)
	}

	err = client.CreateBucket("mybucket1273", oss.ACL(oss.ACLPublicRead))
	if err != nil {
		fmt.Printf("Error in bucket creation: %v", err)
		os.Exit(-1)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////////

}
