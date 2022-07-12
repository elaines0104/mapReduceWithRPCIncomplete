package reducer

import (
	"map-reduce-server/common"
	"map-reduce-server/wordCount"
)

func DoReduceStep(item common.ReduceStep) {
	var reduceF func(key string, values []string) string

	if item.Method == "wordcount" {
		reduceF = wordCount.WordCountReduceF

	}
	doReduceStep(item.JobName, item.ReduceStepNumber, item.NumberOfFiles, reduceF, item.Path)

}

func doReduceStep(
	jobName string,
	reduceTaskNumber int,
	numberOfFiles int,
	reduceF func(key string, values []string) string,
	path string) {

	//TODO

}
