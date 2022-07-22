pragma solidity >=0.7.0 <0.9.0;



contract bet2 {
    address public owner;
    mapping (address=> uint) private moneys;
    bool private state;
    //0表示未参与
    //1表示压大
    //2表示压小
    mapping (address => uint) expect;
    address[] players;
    uint public nouce;

    constructor() payable public {
        owner=msg.sender;
        state=false;
        nouce=0;
    }
    function inject() payable public {
        require(msg.sender==owner,"Only owners can inject capital");
    }
    function joinbet(uint _expect) payable public {
        address sender = msg.sender;
        require(state==true,"There must be a bet in progress");
        require(_expect==1 || _expect==2,"expect must be 1 or 2");
        //未参与
        require(expect[sender]==0,"sender must not be  involved in current bet");
        expect[sender]=_expect;
        moneys[sender]=msg.value;
        players.push(msg.sender);
    }
    function changeowner(address next) public {
        require(owner==msg.sender,"The owner must be provided in order for the authority to change");
        owner=next;
    }
    function start() public {
        require(owner==msg.sender,"Only owners can start a bet");
        require(state==false);
        nouce++;
        state=true;
    }
    function open() public {
        require(owner==msg.sender,"Only possessors will prescribe results");
        uint res=1;
        for(uint i=0;i<players.length;i++){
            address player=players[i];
            if(expect[player]==res){
                payable(player).transfer(moneys[player]*2);
            }
            expect[player]=0;
        }
        payable(owner).transfer(getbalance());
    }
    function getbalance() public view returns (uint) {
        return address(this).balance;
    }
    function getstate() public view returns (bool) {
        return state;
    }
    function getplayersnum() public view returns (uint) {
        return players.length;
    }
    function getplayers(uint id) public view returns (address) {
        require(id<players.length);
        return players[id];
    }
    function getmoney(address player) public view returns (uint) {
        return moneys[player];
    }
}