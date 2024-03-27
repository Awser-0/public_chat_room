window.addEventListener('load', () => {
  const host = `192.168.194.87:3000` // 局域网，为了手机请求数据
  // const host = '127.0.0.1:3000` // 本地
  console.log($);
  let ws;
  let username = ''
  const defaultUsername = 'u99999999', defaultPassword = '123456'
  let isConnecting = false
  let isAutoLogin = true

  const userInput = $('#userInput')
  const passInput = $('#passInput')
  const loginButton = $('#loginButton')
  const closeButton = $('#closeButton')
  const sendButton = $('#sendButton')
  const textarea = $('#textarea')
  // 初始化
  userInput.val(defaultUsername)
  passInput.val(defaultPassword)
  setIsConnecting(false)

  loginButton.on('click', () => {
    login()
  })
  closeButton.on('click', () => {
    closeWS()
  })
  sendButton.on('click', () => {
    sendMsg()
  })
  textarea.on('keyup', (ev) => {
    if(ev.ctrlKey && ev.keyCode==13) {
      sendMsg()
    }
  })
  if(isAutoLogin) {
    login()
  }
  // 切换连接状态信息
  function setIsConnecting(flag) {
    if (flag == true) {
      $('#chat-box-head .state').text('(连接中)')
      loginButton.attr('disabled', true)
      closeButton.attr('disabled', false)
    } else {
      $('#chat-box-head .state').text('(未连接)')
      loginButton.attr('disabled', false)
      closeButton.attr('disabled', true)
    }
    isConnecting = flag
  }
  // 获取输入框文本，并发送
  function sendMsg() {
    const value = textarea.val()
    if(value.trim() != '') {
      const msg = JSON.stringify({
        path: '/chat',
        data: { text: value }
      })
      if(isConnecting) {
        ws.send(msg)
      } else {
        showError('未连接，无法发送消息')
        return
      }
    }
    textarea.val('')
  }

  // 添加聊天记录
  function addMsg(data) {
    const chatMsgs = $('#chat-msgs')
    const msg = $(`
      <div class="card chat-card ${data['user']==username ? 'chat-card--self' : ''}">
        <div class="card-body">
          <h5 class="card-title">${data['user']}</h5>
          <p class="card-text">${data['text']}</p>
        </div>
      </div>
    `)
    msg.on('click', (ev) => {
      const text = ev.currentTarget.querySelector('.card-text').innerText
      // return
      copy(text)
    })
    msg.on('touchstart', () => {

    })
    chatMsgs.append(msg)
    chatMsgs.scrollTop(chatMsgs.prop("scrollHeight"));

  }
  // 复制文本
  function copy(text) {
    let transfer = document.createElement('input');
    document.body.appendChild(transfer);
    transfer.value = text;  // 这里表示想要复制的内容
    transfer.focus();
    transfer.select();
    if (document.execCommand('copy')) {
        document.execCommand('copy');
    }
    transfer.blur();
    document.body.removeChild(transfer);
  }
  // 关闭websocket
  function closeWS() {
    if(ws) {
      ws.close()
    }
  }
  // 显示错误信息
  function showError(msg = '') {
    $('#chat-box-head .error').text(msg)
  }
  // 初始化webSocket，在登录成功拿到token后进行
  function initWebSocket(token) {
    ws = new WebSocket(`ws://${host}/ws?token=${token}`);
    ws.onopen = function() {
      setIsConnecting(true)
      // Web Socket 已连接上，使用 send() 方法发送数据
      // ws.send("发送数据");
      console.log("连接成功...");
      showError()
    };
    
    ws.onmessage = function (evt) { 
      showError()
      var msg = JSON.parse(evt.data);
      try {
        // 处理消息
        const path = msg['path']
        if(path == '/chat') {
          // 有人发送聊天消息
          addMsg(msg['data'])
        } else if(path == 'error') {
          showError(msg['data']['msg'])
        }
      } catch(err) {
        showError(err['message'])
      }
    };
    
    ws.onclose = function() { 
      // 关闭 websocket
      setIsConnecting(false)
      console.log("连接已关闭..."); 
    };

    ws.onerror = function(ev) { 
      // 关闭 websocket
      console.log("发生错误", ev); 
    };
  }
  // 登录
  function login() {
    $.ajax(`http://${host}/user/login`, {
      method: 'GET',
      data: {username: userInput.val(), password: passInput.val()},
      success(result) {
        console.log(result);
        try {
          if(result['code'] == 10200) {
            initWebSocket(result['data']['token'])
            username = result['data']['username']
          } else {
            showError(result.msg)
          }
        } catch(err) {
          showError(err['message'])
        }
      },
      error(err) {
        console.log(err);
        if(err?.responseJSON?.msg) {
          showError(err['responseJSON']['msg'])
        } else {
          showError('请求错误')
        }
      }}
    )
  }
})


