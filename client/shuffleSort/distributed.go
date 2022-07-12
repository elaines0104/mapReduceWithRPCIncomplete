package shuffleSort

import (
	"map-reduce-client/client"
)

func DoMap(jobName string,
	files []string,
	numberOfMapOutput int,
	method string,
	path string,
	column string,
	clients []client.Client) {
	// TODO: RPC Call

}

func DoReduce(
	jobName string,
	numberOfMapOutput int,
	numberOfFiles int,
	method string,
	path string,
	clients []client.Client) {

	// TODO: RPC Call

}
