## 这是一个关于赌博的智能和合约

文件结构如下

- sol文件夹
包含智能合约的 solidity源代码，和编译后的abi 和bin文件
- utils包
实现了go语言与以太坊交互的基本函数
1. func Getclient(host string) 根据结点路径获得客户端
2. GetAddress(pkey string) 根据私钥获得地址
3. Getnouce(add common.Address, client *ethclient.Client) 获得当前账户的nouce
4. GetGasPrice(client *ethclient.Client) 获得当前预估的汽油费
5. Getauth(pkey string, client *ethclient.Client) 根据私钥获得账户签名的交易
- bet包
1. bet.go abigen生成的智能合约接口
2. userbetaction 用户操作接口
- Deploycontract(privatekey string) 发布合约
- Getbetcontract(contractaddress string) 获得合约对象
- getopts(pkey string, value int64) 获得带金额的授权交易
- Startbet(bet *Bet, pkey string) 开始一个赌局
- Inject(bet *Bet, pkey string, value int64) 往赌局中注资
- Openbet(bet *Bet, pkey string) 揭开结果
- Changeowner(bet *Bet, pkey string, next common.Address) 变更赌局拥有者
- Joinbet(bet *Bet, pkey string, value int64, expect int64)参加一个赌局
- Getbetbalance(bet *Bet, pkey string) 得到赌局的现金池
- Getbetstate(bet *Bet, pkey string) 获得赌局状态