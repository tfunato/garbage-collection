// ==UserScript==
// @name           FlashtoJava
// @namespace      gizmode
// @include        http://www.gizmodo.jp/2010/10/macflash.html?utm_source=atom&utm_medium=rss
// ==/UserScript==

var serializer = new XMLSerializer();
var xml = serializer.serializeToString(document.body);
var str = xml.replace(/Flash/g, '<font color="red">Java</font>');
document.title = "新MacからJavaが消えた！ : ギズモード・ジャパン";
document.body.innerHTML = str;

