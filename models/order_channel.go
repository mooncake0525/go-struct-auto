package models 

type Order_channel struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('扣款单uuid')" json:"uuid"` 
    PaymentOrderId string `xorm:"comment('payment order id')" json:"payment_order_id"` 
    FinancialContractNo string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    ChannelAccountNo string `xorm:"comment('项目商户号，例如拍拍贷一期在宝付开的账号')" json:"channel_account_no"` 
    GatewayAccountId string `xorm:"comment('gateway_account_id')" json:"gateway_account_id"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    Delivery int `xorm:"comment('0 wait, 1 sending, 2 querying ')" json:"delivery"` 
    StateDesc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    PlanTime string `xorm:"comment('计划执行开始时间')" json:"plan_time"` 
    RequestNo string `xorm:"comment('请求号')" json:"request_no"` 
    MerchantNo string `xorm:"comment('商户号')" json:"merchant_no"` 
    ProtocolNo string `xorm:"comment('协议支付协议号')" json:"protocol_no"` 
    GatewayId int `xorm:"comment('gateway id')" json:"gateway_id"` 
    OtherDetailJson string `xorm:"comment('分账信息')" json:"other_detail_json"` 
    CreateTime string `xorm:"comment('创建时间')" json:"create_time"` 
    LastModifyTime string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
    OrderTime string `xorm:"comment('排序时间')" json:"order_time"` 
    GroupId int `xorm:"comment('group_id')" json:"group_id"` 
}