const { default: axios } = require("axios");
var ws = require("ws");

// 初始化websocket，在登录拿到token后
function startWebSocket(user, token) {
  var sock = new ws("ws://127.0.0.1:3000/ws?token=" + token);
  let timer = null
  sock.on("open", function () {
    console.log(`${user} ->`, '连接成功');
    timer = setInterval(() => {
      randomText().then(text => {
        if(!text) return
        sock.send(JSON.stringify({
          path: '/chat',
          data: {
            text: text
          }
        }))
      })
    }, 2000)
  });
  
  sock.on("error", function(err) {
    console.error(`${user} ->`, 'error');
  });
  
  sock.on("close", function() {
    clearInterval(timer)
    console.info(`${user} ->`, '连接关闭');
  });
  
  sock.on("message", function(data) {
    var msg = JSON.parse(data);
    try {
      // 处理消息
      const path = msg['path']
      if(path == '/chat') {
      } else if(path == 'error') {
        console.error(msg['data']['msg'])
      }
    } catch(err) {
      console.error(err['message'])
    }
  })
}
// 登录
async function login(user, pass) {
  await axios.get(`http://127.0.0.1:3000/user/login?username=${user}&password=${pass}`)
  .then(({data: result}) => {
    if(result['code'] == 10200) {
      startWebSocket(result['data']['username'], result['data']['token'])
    } else {
      throw new Error(result['msg'])
    }
  }).catch(err => {
    console.log('登录失败·', err['message']);
  })
}
// Math.ceil(Math.random()*10)https://tenapi.cn/v2/yiyan
// 随机名言
async function randomText() {
  return Promise.resolve(genNickName())
  // return await axios.post('https://tenapi.cn/v2/yiyan', {format: 'json'}).then(({data}) => {
  //   return Promise.resolve(data)
  // }).catch(err => {
  //   console.log(err);
  //   return Promise.resolve(null)
  // })
}

// 上压力
const max = 10
function main() {
  for(let i=1; i<=max; i++) {
    login('u' + i, '123456')
  }
}

function genNickName() {
  // 获取指定范围内的随机数
  function randomAccess(min, max) {
      return Math.floor(Math.random() * (min - max) + max)
  }

  // 解码
  function decodeUnicode(str) {
      //Unicode显示方式是\u4e00
      str = "\\u" + str;
      str = str.replace(/\\/g, "%");
      //转换中文
      str = unescape(str);
      //将其他受影响的转换回原来
      str = str.replace(/%/g, "\\");
      return str;
  }

  function getRandomName(len) {
      let name = ""
      for (let i = 0; i < len; i++) {
          let unicodeNum = ""
          unicodeNum = randomAccess(0x4e00, 0x9fa5).toString(16)
          name += decodeUnicode(unicodeNum)
      }
      return name;
  }
  return getRandomName(4);
}



main()