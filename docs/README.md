# Account 是账号系统主要职责如下

### 一、目标
#### 1.1 Web
    - Login with username & password
    - Login with phone number 
    - Logout
    - SendDigit
    - VerifyDigit
    - Login With WeChat
    - Bind/Unbind WeChat

#### 1.2 RPC
    - AuthWithWeChatOpenID
    - AUthWithAccessToken
   

#### 1.3三方登录
    - wechat login 
    

### 二、信息结构
#### User
| 字段| 类型 | 含义| 是否必须| 备注| 
|:---|:---:|:---:|:---:|:---:| 
| id| bigint|存储自增 id| 是|
| phone|string|手机号码|是|country-code + phone number|
| uid|string（64| 全局 ID |是|由发号器统一生成|
| password|string(64)|加密后的密码|是| RSA 加密后的结果|
| status | int| 状态| 是| 0，禁止； 10 正常； 后续可以扩充|
| create_at|datetime|创建时间|是||
| updated_at|datetime|更新时间|是||
| last_login|datetime|最后登录时间|是||

#### AccessToken
| 字段| 类型 | 含义| 是否必须| 备注| 
|:---|:---:|:---:|:---:|:---:| 
| id| 
| uid|
| token| string| token 由 account 生成| 是| 同一时间允许 uid 存在多个|
| type|int|token 类型| 是| 对内 or 对外|
| devide_info|json||是| 登录设备信息
| expire_at|datetime|最后登录时间|是||


#### UserProfile 
| 字段| 类型 | 含义| 是否必须| 备注| 
|:---|:---:|:---:|:---:|:---:| 
|uid| 
|avatar|
|name|
|gender|
|created_at|
|updated_at|

#### UserWechatInfo
| 字段| 类型 | 含义| 是否必须| 备注| 
|:---|:---:|:---:|:---:|:---:| 
|id|
|openid|
|unionid|
|access_token|
|refresh_token|
|expires_at|
|name|
|gender|
|province|
|city|
|country|
|avatar|
|last_updated|
|created_at|

#### UserWechatRelationship
| 字段| 类型 | 含义| 是否必须| 备注| 
|:---|:---:|:---:|:---:|:---:| 
|id|
|unionid|
|is_deleted|
    
