Array.prototype.randomElement = function () {
    return this[Math.floor(Math.random() * this.length)]
};
var stuff = ['unicorn', 'leprachaun', 'bucket full of worms', 'nice little thingy'];

if(sender === "fantaraa"){
    sendMessage('moe_nana gives herself a ' + stuff.randomElement() + '. She has a lot of fun with it. Look away. Creep.')
}else {
    sendMessage('moe_nana gives a ' + stuff.randomElement() + ' to ' + sender + '. Have fun with it, ' + sender + '!')
}