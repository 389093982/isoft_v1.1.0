package entry

// 调度者
type Dispatcher struct {
	TrackingId string                 // 调度者执行时的 trackingId
	TmpDataMap map[string]interface{} // 调度者发送过来的临时数据
}

// 接收者
type Receiver struct {
	TmpDataMap map[string]interface{} // 被调度者发送过来的临时数据
}

