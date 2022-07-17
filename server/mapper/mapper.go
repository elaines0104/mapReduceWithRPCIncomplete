package mapper

import (
	"hash/fnv"
	"log"
	"map-reduce-server/common"
	"map-reduce-server/wordCount"
)

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func DoMapStep(item common.MapStep) {
	var mapF func(file string, contents string) []common.KeyValue
	log.Printf("UseCase %s", item.UseCase)

	if item.UseCase == "wordcount" {
		mapF = wordCount.WordCountMapF

	}
	doMapStep(item.JobName, item.MapStepNumber, item.File, item.NumberOfMapOutput, mapF, item.Path, item.Column)

}

func doMapStep(
	jobName string,
	mapTaskNumber int,
	inFile string,
	numberOfMapOutput int,
	mapF func(file string, contents string) []common.KeyValue,
	path string,
	column string) {
	//TODO

}
