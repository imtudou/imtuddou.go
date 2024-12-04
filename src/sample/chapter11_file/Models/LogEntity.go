package Models

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type LogEntity struct {
	Time        string `json:"time"`        //日志生成的时间戳
	Level       string `json:"level"`       // 日志级别，如INFO, ERROR, DEBUG等
	Category    string `json:"category"`    // 日志类
	Theard      string `json:"theard"`      // 日志产生线程
	TraceId     string `json:"traceId"`     // 链路 ID
	RequestId   string `json:"requestId"`   // 请求 ID
	UserId      string `json:"userId"`      // 用户 ID
	ShopAdminId string `json:"shopAdminId"` // 商户 ID
	BizKey      string `json:"bizKey"`      // 业务Key
	Format      string `json:"format"`      // 格式化信息
	Message     string `json:"message"`     // 日志信息内容
}

// InitLogFile 初始化日至文件
func InitLogFile() {
	// 打开文件，如果文件不存在则创建文件,在文件打开时清空文件的内容
	filePath := "../../../../doc/log.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// 确保在函数结束时关闭文件
	defer file.Close()
	// 创建一个新的写入器（带缓冲区）
	writer := bufio.NewWriter(file)

	content := "%v INFO  c.x.r.c.c.GlobalControllerAspect http-nio-11075-exec-6 [%v]-%v--183613---  |请求地址:https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=7c721072-fcd5-4787-8cf6-155aa34cd6b3|请求方式:POST|请求耗时:632|请求参数:{\\\"msgtype\\\":\\\"text\\\",\\\"text\\\":{\\\"content\\\":\\\"|记录来源:.Net SqlClient Data Provider|方法来源:Kmmp.Marketing.DI.TicketStatisticsDI:Excute|记录内容:\\\\\\\"执行超时已过期。完成操作之前已超时或服务器未响应。 \\\\\\\"|堆栈信息:\\\\\\\"   在 System.Data.SqlClient.SqlConnection.OnError(SqlException exception, Boolean breakConnection, Action`1 wrapCloseInAction)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.TdsParser.ThrowExceptionAndWarning(TdsParserStateObject stateObj, Boolean callerHasConnectionLock, Boolean asyncClose)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.TdsParser.TryRun(RunBehavior runBehavior, SqlCommand cmdHandler, SqlDataReader dataStream, BulkCopySimpleResultSet bulkCopyHandler, TdsParserStateObject stateObj, Boolean& dataReady)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlDataReader.TryConsumeMetaData()\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlDataReader.get_MetaData()\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlCommand.FinishExecuteReader(SqlDataReader ds, RunBehavior runBehavior, String resetOptionsString, Boolean isInternal, Boolean forDescribeParameterEncryption, Boolean shouldCacheForAlwaysEncrypted)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlCommand.RunExecuteReaderTds(CommandBehavior cmdBehavior, RunBehavior runBehavior, Boolean returnStream, Boolean async, Int32 timeout, Task& task, Boolean asyncWrite, Boolean inRetry, SqlDataReader ds, Boolean describeParameterEncryptionRequest)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlCommand.RunExecuteReader(CommandBehavior cmdBehavior, RunBehavior runBehavior, Boolean returnStream, String method, TaskCompletionSource`1 completion, Int32 timeout, Task& task, Boolean& usedCache, Boolean asyncWrite, Boolean inRetry)\\\\\\\\r\\\\\\\\n   在 System.Data.SqlClient.SqlCommand.RunExecuteReader(CommandBehavior cmdBehavior, RunBehavior runBehavior, Boolean returnStream, String method)   在 System.Data.SqlClient.SqlCommand.ExecuteReader(CommandBehavior behavior, String method)\\\\\\\\r\\\\\\\\n   在 Dapper.SqlMapper.<QueryImpl>d__61`1.MoveNext()\\\\\\\\r\\\\\\\\n   在 System.Collections.Generic.List`1..ctor(IEnumerable`1 collection)\\\\\\\\r\\\\\\\\n   在 System.Linq.Enumerable.ToList[TSource](IEnumerable`1 source)\\\\\\\\r\\\\\\\\n   在 Dapper.SqlMapper.Query[T](IDbConnection cnn, String sql, Object param, IDbTransaction transaction, Boolean buffered, Nullable`1 commandTimeout, Nullable`1 commandType)\\\\\\\\r\\\\\\\\n   在 Kmmp.Core.Imps.IDBDataSource.ExecuteReaderReturnListT[T](String connStr, String sql, Object param) 位置 E:\\\\\\\\\\\\\\\\Repos\\\\\\\\\\\\\\\\KmmpPlat\\\\\\\\\\\\\\\\Kmmp.Plat\\\\\\\\\\\\\\\\src\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\Kmmp.Core\\\\\\\\\\\\\\\\Imps\\\\\\\\\\\\\\\\IDBDataSource.cs:行号 406\\\\\\\\r\\\\\\\\n   在 Kmmp.Core.Helper.DBDataSourceHelper.ExecuteReaderReturnListT[T](String connStr, String sql, Object param) 位置 E:\\\\\\\\\\\\\\\\Repos\\\\\\\\\\\\\\\\KmmpPlat\\\\\\\\\\\\\\\\Kmmp.Plat\\\\\\\\\\\\\\\\src\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\Kmmp.Core\\\\\\\\\\\\\\\\Helper\\\\\\\\\\\\\\\\DBDataSourceHelper.cs:行号 562\\\\\\\\r\\\\\\\\n   在 Kmmp.Marketing.DI.TicketStatisticsDI.GetTempTicketInfo(String tableName, TicketStatisticsInput input) 位置 E:\\\\\\\\\\\\\\\\Repos\\\\\\\\\\\\\\\\KmmpPlat\\\\\\\\\\\\\\\\Kmmp.Plat\\\\\\\\\\\\\\\\src\\\\\\\\\\\\\\\\Marketing\\\\\\\\\\\\\\\\Kmmp.Marketing.DI\\\\\\\\\\\\\\\\TicketStatisticsDI.cs:行号 68\\\\\\\\r\\\\\\\\n   在 Kmmp.Marketing.DI.TicketStatisticsDI.Excute(TicketStatisticsInput input) 位置 E:\\\\\\\\\\\\\\\\Repos\\\\\\\\\\\\\\\\KmmpPlat\\\\\\\\\\\\\\\\Kmmp.Plat\\\\\\\\\\\\\\\\src\\\\\\\\\\\\\\\\Marketing\\\\\\\\\\\\\\\\Kmmp.Marketing.DI\\\\\\\\\\\\\\\\TicketStatisticsDI.cs:行号 35\\\\\\\"|附加信息:优惠券统计\\\"}}\\\" | 响应参数:\\\"{\\\"errcode\\\":0,\\\"errmsg\\\":\\\"ok\\\"}"
	for i := 0; i < 1000000; i++ {
		str := fmt.Sprintf(content+"\n", time.Now().Format("2006/01/02 15:04:05.999"), strings.Replace(uuid.New().String(), "-", "", -1), rand.Int63())
		writer.WriteString(str) // 循环10000000次 耗时:1m30.5569414s
	}

	// 刷新缓冲区，确保所有内容都被写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

}

// AnalysisFile 读文件并生成josn文件
func AnalysisFile() {

	/*			解析日志行中的以下字段：
	time: 日志生成的时间戳
	level: 日志级别，如INFO, ERROR, DEBUG等
	category: 日志类
	theard: 日志产生线程
	traceId: 链路 ID
	requestId: 请求 ID
	userId: 用户 ID
	shopAdminId: 商户 ID
	bizKey: 业务Key
	format: 格式化信息
	message: 日志信息内容

		2024-10-08 14:46:11,324 INFO  c.x.r.c.c.GlobalControllerAspect http-nio-11075-exec-6 [0b32823e17283699713003920ec4ea]-210407978682429440--183613---This is message
	*/
	/*	msg := " 2024/10/17 18:06:20.219 INFO  c.x.r.c.c.GlobalControllerAspect http-nio-11075-exec-6 [0b32823e17283699713003920ec4ea]-210407978682429440--183613--- |请求地址:https://nacos-kf-new.cloud.kemai.cn/nacos/v1/auth/login|请求方式:POST|请求耗时:0|请求参数:\"username=nacos&password=nacos\"|响应参数:null"
		WriteJsonFile(msg)*/

	// 打开文件
	file, err := os.Open("../../../../doc/log.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 函数退出时关闭文件
	defer file.Close()

	// 使用*reader 对象读取 带缓冲区
	reader := bufio.NewReader(file)

	/*		for i := 0; i < 3; i++ {
			line, _, err := reader.ReadLine()
			if err == io.EOF { // io.EOF 表示文件结尾
				break
			}
			msg := string(line)
			HandleData(msg)
		}*/

	var sb strings.Builder
	//var logs []Models
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF { // io.EOF 表示文件结尾
			break
		}
		msg := string(line)
		sb = ConvertTojson(msg, sb)
		//logs = append(logs, ConvertToLogEntity(msg))
		if i%100000 == 0 {
			//序列化并写入JSON文件
			//logMsg, _ := json.Marshal()

			jsonstr := sb.String()
			if jsonstr != "" {
				WriteFile("../../../../doc/parsed_logs.json", jsonstr)
				sb.Reset()
			}
		}
		i++
	}

}

// ConvertToLogEntity 将读取的内容转成 LogEntity
func ConvertToLogEntity(msg string) LogEntity {
	//msg := "2024/10/17 18:08:27.067 INFO  c.x.r.c.c.GlobalControllerAspect http-nio-11075-exec-6 [0b32823e17283699713003920ec4ea]-210407978682429440--183613--- |请求地址:https://nacos-kf-new.cloud.kemai.cn/nacos/v1/auth/login|请求方式:POST|请求耗时:0|请求参数:"username=nacos&password=nacos"|响应参数:null"
	arr := strings.Split(msg, "]-")
	if len(arr) == 0 {
		return LogEntity{}
	}

	beginStr := arr[0]
	endStr := arr[1]

	beginArr := strings.Split(beginStr, " ")
	endArr := strings.Split(endStr, "-")

	/*	for i, v := range beginArr {
			fmt.Printf("i=%v,%v \n", i, v)
		}

		for i, v := range endArr {
			fmt.Printf("i=%v,%v \n", i, v)
		}*/

	logEntity := LogEntity{
		Time:        "",
		Level:       "",
		Category:    "",
		Theard:      "",
		TraceId:     "",
		RequestId:   "",
		UserId:      "",
		ShopAdminId: "",
		BizKey:      "",
		Format:      "",
		Message:     "",
	}

	if len(beginArr) > 0 {
		logEntity.Time = fmt.Sprintf("%v %v", beginArr[0], beginArr[1])
		logEntity.Level = beginArr[2]
		logEntity.Category = beginArr[4]
		logEntity.Theard = beginArr[5]
		logEntity.TraceId = strings.Replace(beginArr[6], "[", "", -1)
	}

	if len(endArr) > 0 {
		logEntity.RequestId = endArr[0]
		logEntity.UserId = endArr[1]
		logEntity.ShopAdminId = endArr[2]
		logEntity.BizKey = endArr[3]
		logEntity.Format = endArr[4]
		if len(endArr[5:]) > 0 { // 提取从第6个元素（索引从0开始）至末尾的子切片
			logEntity.Message = strings.Join(endArr[5:], "-")
		}
	}
	return logEntity
}

// ConvertTojson 将读取的内容转成 json 字符串
func ConvertTojson(msg string, sb strings.Builder) strings.Builder {
	//msg := "2024/10/17 18:08:27.067 INFO  c.x.r.c.c.GlobalControllerAspect http-nio-11075-exec-6 [0b32823e17283699713003920ec4ea]-210407978682429440--183613--- |请求地址:https://nacos-kf-new.cloud.kemai.cn/nacos/v1/auth/login|请求方式:POST|请求耗时:0|请求参数:"username=nacos&password=nacos"|响应参数:null"
	arr := strings.Split(msg, "]-")
	if len(arr) == 0 {
		return strings.Builder{}
	}

	beginStr := arr[0]
	endStr := arr[1]
	beginArr := strings.Split(beginStr, " ")
	endArr := strings.Split(endStr, "-")

	// {"Time":"2024/10/18 11:39:58.381","Level":"INFO","Category":"c.x.r.c.c.GlobalControllerAspect","Theard":"http-nio-11075-exec-6","TraceId":"2881f2a411c642ef8121e0142f30e9d0","RequestId":"973829208395239109","UserId":"","ShopAdminId":"183613","BizKey":"","Format":"","Message":" |请求地址:https://nacos-kf-new.cloud.kemai.cn/nacos/v1/auth/login|请求方式:POST|请求耗时:0|请求参数:\"username=nacos\u0026password=nacos\"|响应参数:null"}
	sb.WriteString("{")
	if len(beginArr) > 0 {

		time := fmt.Sprintf(`"time":"%v",`, fmt.Sprintf("%v %v", beginArr[0], beginArr[1]))
		sb.WriteString(time)
		level := fmt.Sprintf(`"level":"%v",`, beginArr[2])
		sb.WriteString(level)
		category := fmt.Sprintf(`"category":"%v",`, beginArr[4])
		sb.WriteString(category)
		theard := fmt.Sprintf(`"theard":"%v",`, beginArr[5])
		sb.WriteString(theard)
		traceId := fmt.Sprintf(`"traceId":"%v",`, strings.Replace(beginArr[6], "[", "", -1))
		sb.WriteString(traceId)
	}

	if len(endArr) > 0 {
		requestId := fmt.Sprintf(`"requestId":"%v",`, endArr[0])
		sb.WriteString(requestId)
		userId := fmt.Sprintf(`"userId":"%v",`, endArr[1])
		sb.WriteString(userId)
		shopAdminId := fmt.Sprintf(`"shopAdminId":"%v",`, endArr[2])
		sb.WriteString(shopAdminId)
		bizKey := fmt.Sprintf(`"bizKey":"%v",`, endArr[3])
		sb.WriteString(bizKey)
		format := fmt.Sprintf(`"format":"%v",`, endArr[4])
		sb.WriteString(format)
		message := fmt.Sprintf(`"message":""`)
		if len(endArr[5:]) > 0 { // 提取从第6个元素（索引从0开始）至末尾的子切片

			replacestr := strings.Join(endArr[5:], "-")
			//strings.Replace(replacestr, "\"",`\"`,-1 )
			message = fmt.Sprintf(`"message": "%v"`, replacestr)
			sb.WriteString(message)
		} else {
			sb.WriteString(message)
		}
	}
	sb.WriteString("},")
	return sb
}

func WriteFile(path string, str string) {
	// 打开文件，如果文件不存在则创建文件，并追加内容到文件末尾
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// 创建一个新的写入器（带缓冲区）
	writer := bufio.NewWriter(file)

	_, err1 := writer.WriteString(str) // 循环10000000次 耗时:1m30.5569414s
	if err1 != nil {
		fmt.Println("Error WriteString:", err1)
	}
	// 刷新缓冲区，确保所有内容都被写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
	}
}
