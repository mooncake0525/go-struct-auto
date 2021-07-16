package models 

type Order_channel struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('扣款单uuid')" json:"uuid"` 
    Payment_order_id string `xorm:"comment('payment order id')" json:"payment_order_id"` 
    Financial_contract_no string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    Channel_account_no string `xorm:"comment('项目商户号，例如拍拍贷一期在宝付开的账号')" json:"channel_account_no"` 
    Gateway_account_id string `xorm:"comment('gateway_account_id')" json:"gateway_account_id"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    Delivery int `xorm:"comment('0 wait, 1 sending, 2 querying ')" json:"delivery"` 
    State_desc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    Plan_time string `xorm:"comment('计划执行开始时间')" json:"plan_time"` 
    Request_no string `xorm:"comment('请求号')" json:"request_no"` 
    Merchant_no string `xorm:"comment('商户号')" json:"merchant_no"` 
    Protocol_no string `xorm:"comment('协议支付协议号')" json:"protocol_no"` 
    Gateway_id int `xorm:"comment('gateway id')" json:"gateway_id"` 
    Other_detail_json string `xorm:"comment('分账信息')" json:"other_detail_json"` 
    Create_time string `xorm:"comment('创建时间')" json:"create_time"` 
    Last_modify_time string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
    Order_time string `xorm:"comment('排序时间')" json:"order_time"` 
    Group_id int `xorm:"comment('group_id')" json:"group_id"` 
}