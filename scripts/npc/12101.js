/* Author: Xterminator
    NPC Name: 		Rain
    Map(s): 		Maple Road : Amherst (1010000)
    Description: 		Talks about Amherst
*/
var status = -1;

function start() {
    action(1, 0, 0);
}

function action(mode, type, selection) {
    status++;
    if (mode == 0) {
        status -= 3
        action(1, 0, 0);
    } else if(mode == 1){
        if (status == 0) {
            cm.sendNext("这是一个名为 #b彩虹村#k 的小镇，坐落于冒险岛大陆的最北边。这里是通往冒险岛大陆的起点。我很庆幸这附近只有一些弱小的怪物。");
        } else if (status == 1) {
            cm.sendNextPrev("如果你想变得更强，你需要去一个名叫 #b南港#k 的地方。在那停着一个前往 #b金银岛#k 的巨大船只。与这个小岛相比，它的大小无与伦比。");
        } else if (status == 2) {
            cm.sendPrev("在金银岛，你可以选择你的职业。是 #b勇士部落#k 吗？我听说那是一个荒凉的小镇，那里住着勇士。高原。。。那是什么样的地方？");
        } else if (status == 3) {
            cm.dispose();
        }
    }else{
        cm.dispose();
    }
}