package models 

type Order_payment struct { 
    Id int `xorm:"not null pk autoincr comment('id') INT(11)" json:"id"` 
    Uuid string `xorm:"comment('uuid')" json:"uuid"` 
    Payment_order_id string `xorm:"comment('全局唯一编号')" json:"payment_order_id"` 
    Financial_contract_no string `xorm:"comment('项目编号')" json:"financial_contract_no"` 
    State int `xorm:"comment('当前状态')" json:"state"` 
    State_desc string `xorm:"comment('当前状态的描述')" json:"state_desc"` 
    Notify_status int `xorm:"comment('回调状态 0-不用回调 1-待回调 2-回调中 3-回调成功 4-回调失败')" json:"notify_status"` 
    Notify_remark string `xorm:"comment('回调备注')" json:"notify_remark"` 
    Account_side int `xorm:"comment('借贷标记,0代付1代扣')" json:"account_side"` 
    Trade_amount float64 `xorm:"comment('交易金额')" json:"trade_amount"` 
    Recv_acc_no string `xorm:"comment('对方账号')" json:"recv_acc_no"` 
    Recv_acc_name string `xorm:"comment('对方账户名称')" json:"recv_acc_name"` 
    Recv_bank_code string `xorm:"comment('对方恒生代码')" json:"recv_bank_code"` 
    Recv_std_bank_code string `xorm:"comment('对方银行联行号')" json:"recv_std_bank_code"` 
    Recv_bank_name string `xorm:"comment('对方银行名称')" json:"recv_bank_name"` 
    Recv_bank_alias_name string `xorm:"comment('对方银行代码')" json:"recv_bank_alias_name"` 
    Recv_open_bank_pro_code string `xorm:"comment('对方开户行省份代码')" json:"recv_open_bank_pro_code"` 
    Recv_open_bank_city_code string `xorm:"comment('对方开户行城市代码')" json:"recv_open_bank_city_code"` 
    Recv_open_bank_std_code string `xorm:"comment('对方开户行联行号')" json:"recv_open_bank_std_code"` 
    Recv_open_bank_name string `xorm:"comment('对方开户行名称')" json:"recv_open_bank_name"` 
    Mobile string `xorm:"comment('手机号')" json:"mobile"` 
    Id_number string `xorm:"comment('身份证号')" json:"id_number"` 
    Remark string `xorm:"comment('备注')" json:"remark"` 
    Summary string `xorm:"comment('summary')" json:"summary"` 
    Notify_url string `xorm:"comment('回调地址')" json:"notify_url"` 
    Detail_json string `xorm:"comment('其他信息，例如分账信息')" json:"detail_json"` 
    Create_time string `xorm:"comment('创建时间')" json:"create_time"` 
    Last_modify_time string `xorm:"comment('最近修改时间')" json:"last_modify_time"` 
    Operator string `xorm:"comment('更新者，子系统名')" json:"operator"` 
}