package shuffleSort

import (
	"map-reduce-client/client"
)

func DoMap(jobName string,
	files []string,
	numberOfMapOutput int,
	useCase string,
	path string,
	column string,
	clients []client.Client) {
	// TODO: RPC Call

}

func DoReduce(
	jobName string,
	numberOfMapOutput int,
	numberOfFiles int,
	useCase string,
	path string,
	clients []client.Client) {

	// TODO: RPC Call

}
