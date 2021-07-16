package models 

type Order_payment struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('uuid')" json:"uuid"` 
    PaymentOrderId string `xorm:"comment('全局唯一编号')" json:"payment_order_id"` 
    FinancialContractNo string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    StateDesc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    NotifyStatus int `xorm:"comment('回调状态 0-不用回调 1-待回调 2-回调中 3-回调成功 4-回调失败')" json:"notify_status"` 
    NotifyRemark string `xorm:"comment('回调备注')" json:"notify_remark"` 
    AccountSide int `xorm:"comment('借贷标记,0代付1代扣')" json:"account_side"` 
    TradeAmount float64 `xorm:"comment('交易金额')" json:"trade_amount"` 
    RecvAccNo string `xorm:"comment('对方账号')" json:"recv_acc_no"` 
    RecvAccName string `xorm:"comment('对方账户名称')" json:"recv_acc_name"` 
    RecvBankCode string `xorm:"comment('对方恒生代码')" json:"recv_bank_code"` 
    RecvStdBankCode string `xorm:"comment('对方银行联行号')" json:"recv_std_bank_code"` 
    RecvBankName string `xorm:"comment('对方银行名称')" json:"recv_bank_name"` 
    RecvBankAliasName string `xorm:"comment('对方银行代码')" json:"recv_bank_alias_name"` 
    RecvOpenBankProCode string `xorm:"comment('对方开户行省份代码')" json:"recv_open_bank_pro_code"` 
    RecvOpenBankCityCode string `xorm:"comment('对方开户行城市代码')" json:"recv_open_bank_city_code"` 
    RecvOpenBankStdCode string `xorm:"comment('对方开户行联行号')" json:"recv_open_bank_std_code"` 
    RecvOpenBankName string `xorm:"comment('对方开户行名称')" json:"recv_open_bank_name"` 
    Mobile string `xorm:"comment('手机号')" json:"mobile"` 
    IdNumber string `xorm:"comment('身份证号')" json:"id_number"` 
    Remark string `xorm:"comment('备注')" json:"remark"` 
    Summary string `xorm:"comment('summary')" json:"summary"` 
    NotifyUrl string `xorm:"comment('回调地址')" json:"notify_url"` 
    DetailJson string `xorm:"comment('其他信息，例如分账信息')" json:"detail_json"` 
    CreateTime string `xorm:"comment('创建时间')" json:"create_time"` 
    LastModifyTime string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
}