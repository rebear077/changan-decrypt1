package canal

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	queue "github.com/FISCO-BCOS/go-sdk/structure"
	types "github.com/FISCO-BCOS/go-sdk/type"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/withlin/canal-go/client"
	pbe "github.com/withlin/canal-go/protocol/entry"
)

type Connector struct {
	conn    *client.SimpleCanalConnector
	queue   *queue.CircleQueue
	RawData []*types.RawSQLData
	Lock    sync.RWMutex
}

func NewConnector(table string) *Connector {
	connector := client.NewSimpleCanalConnector("127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)

	// err = connector.Subscribe("db_node1\\.u_t_history_settle_information")
	err = connector.Subscribe(table)
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
	queue := queue.NewCircleQueue(20)
	raw := make([]*types.RawSQLData, 0)
	return &Connector{
		conn:    connector,
		queue:   queue,
		RawData: raw,
	}
}

// 开始运行canal
func (c *Connector) Start() {
	for {
		message, err := c.conn.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(1 * time.Second)
			logrus.Println("===没有数据了===")
			continue
		}
		c.dealMessage(message.Entries)
	}
}
func (c *Connector) dealMessage(entrys []pbe.Entry) {
	for _, entry := range entrys {
		fmt.Println(entry.GetEntryType())
		if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(pbe.RowChange)
		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			// header := entry.GetHeader()
			// fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == pbe.EventType_DELETE {
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == pbe.EventType_INSERT {
					c.dealInsertMessage(rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
}
func (c *Connector) dealInsertMessage(columns []*pbe.Column) {
	rawdata := new(types.RawSQLData)
	for _, col := range columns {
		err := c.queue.Add(col.GetValue())
		if err != nil {
			logrus.Errorln(err)
			os.Exit(1)
		}
		// fmt.Println(fmt.Sprintf("%s:%s", col.GetName(), col.GetValue()))
	}
	rawdata.SQLId, _ = c.queue.Remove()
	rawdata.Num, _ = c.queue.Remove()
	rawdata.Status, _ = c.queue.Remove()
	rawdata.ID, _ = c.queue.Remove()
	rawdata.Data, _ = c.queue.Remove()
	rawdata.Key, _ = c.queue.Remove()
	rawdata.Hash, _ = c.queue.Remove()
	c.Lock.Lock()
	c.RawData = append(c.RawData, rawdata)
	c.Lock.Unlock()
	// fmt.Println(rawdata)
	//TODO:解密和存储
}
func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s:%s", col.GetName(), col.GetValue()))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
