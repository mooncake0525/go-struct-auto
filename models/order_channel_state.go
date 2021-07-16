package models 

type Order_channel_state struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('扣款单uuid')" json:"uuid"` 
    OrderChannelId string `xorm:"comment('order channel id')" json:"order_channel_id"` 
    PaymentOrderId string `xorm:"comment('payment_order_id')" json:"payment_order_id"` 
    FinancialContractNo string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    ChannelAccountNo string `xorm:"comment('项目商户号，例如拍拍贷一期在宝付开的账号')" json:"channel_account_no"` 
    GatewayAccountId string `xorm:"comment('gateway_account_id')" json:"gateway_account_id"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    StateDesc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    RequestNo string `xorm:"comment('请求号')" json:"request_no"` 
    PlanTime string `xorm:"comment('计划执行开始时间')" json:"plan_time"` 
    EventTime string `xorm:"comment('发生时间')" json:"event_time"` 
    LastModifyTime string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    CompleteTime string `xorm:"comment('完成时间')" json:"complete_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
    EventCode string `xorm:"comment('触发事件code')" json:"event_code"` 
    EventInfo string `xorm:"comment('触发事件')" json:"event_info"` 
}