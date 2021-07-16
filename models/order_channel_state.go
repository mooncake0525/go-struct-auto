package models 

type Order_channel_state struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('扣款单uuid')" json:"uuid"` 
    Order_channel_id string `xorm:"comment('order channel id')" json:"order_channel_id"` 
    Payment_order_id string `xorm:"comment('payment_order_id')" json:"payment_order_id"` 
    Financial_contract_no string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    Channel_account_no string `xorm:"comment('项目商户号，例如拍拍贷一期在宝付开的账号')" json:"channel_account_no"` 
    Gateway_account_id string `xorm:"comment('gateway_account_id')" json:"gateway_account_id"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    State_desc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    Request_no string `xorm:"comment('请求号')" json:"request_no"` 
    Plan_time string `xorm:"comment('计划执行开始时间')" json:"plan_time"` 
    Event_time string `xorm:"comment('发生时间')" json:"event_time"` 
    Last_modify_time string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    Complete_time string `xorm:"comment('完成时间')" json:"complete_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
    Event_code string `xorm:"comment('触发事件code')" json:"event_code"` 
    Event_info string `xorm:"comment('触发事件')" json:"event_info"` 
}