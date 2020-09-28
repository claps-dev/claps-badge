var svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
svg.style.height='160px';
var x = document.getElementsByClassName("claps-donate-button");
var c=x[0].getAttribute("data-button-id").replace(/#/g, "%23");//get请求中，16进制颜色的#号需要转义
c=c.split("|");//分割字符串获取参数
var pnode=x[0].parentNode;//父节点
var a="http://127.0.0.1:8081/?project="+c[0]+"&money="+c[1]+"&size="+c[2]+"&bgcolor="+c[3]+"&color="+c[4];//拼接get请求url
fetch(a).then(response => response.text())//解析为可读数据
    .then(data => svg.innerHTML=data)//执行结果是 resolve就调用then方法
    .catch(err => console.log("Oh, error", err))//执行结果是 reject就调用catch方法
pnode.insertBefore(svg,x[0]);


