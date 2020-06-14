package main

import (
	//support automatic consumer-group rebalancing and offset tracking
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("go_kafka")

func main() {
	logger.Infof("vim-go")
	groupID := "group-1"
	topicList := "topic_1"
	logger.Debugf("the groupID is %s, topicList is %s", groupID, topicList)
}
