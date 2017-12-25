Array.prototype.randomElement = function () {
    return this[Math.floor(Math.random() * this.length)]
};

var stuff = ['a puking unicorn', 'a raging leprachaun', 'a bucket full of worms', 'some frilly panties', 'a large trout', 'a feather'];


if (sender.toLowerCase().trim() === params.toLowerCase().trim() || params === "") {
    sender = "/me"
}
sendMessage(sender + ' slaps  ' + params + ' with ' + stuff.randomElement() + '!')
