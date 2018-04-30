### building a
### notifications bot

---

### create a telegram bot

tell the [@botfather](http://t.me/botfather):
> /newbot

---

### start the fn server

```sh
fn start
```

---

### start the fn ui

```sh
docker run --rm -it --link fnserver:api -p 4000:4000 -e "FN_API_URL=http://api:8080" fnproject/ui
```

---

### create a new application

```sh
fn init --runtime node tbot-eg
```

---

### add dependencies

```sh
cd tbot-eg;
npm i node-telegram-bot-api --save;
```

---

### add code

```js
const fdk = require('@fnproject/fdk');
const TelegramBot = require('node-telegram-bot-api');
const telegramBotToken = process.env['TELEGRAM_BOT_TOKEN'];
const telegramChatId = process.env['TELEGRAM_CHAT_ID'];

fdk.handle(function(input){
  const bot = new TelegramBot(telegramBotToken);
  bot.sendMessage(telegramChatId, `Hello`);
  return 'ok';
});

```

---

### get the telegram chat
Talk to the bot and visit the URL to get the chat ID:

```
https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/getUpdates
```

---

### test run the code

```sh
fn run -e TELEGRAM_BOT_TOKEN="${TELEGRAM_BOT_TOKEN}" -e TELEGRAM_CHAT_ID="${TELEGRAM_CHAT_ID}"
```

---

### add environment variable requirements

open the `func.yaml` and add:

```yaml
expects:
  config:
  - name: TELEGRAM_BOT_TOKEN
    required: true
  - name: TELEGRAM_CHAT_ID
    required: true
```

---

### deploy the code

```sh
fn deploy --app example --local
```

---

### add the environment variables

go to http://localhost:4000/#/app/example

---

### run it in production

go to http://localhost:8080/r/example/tbot-eg

---

### making it useful

```js
const fdk = require('@fnproject/fdk');
const TelegramBot = require('node-telegram-bot-api');
const telegramBotToken = process.env['TELEGRAM_BOT_TOKEN'];
const telegramChatId = process.env['TELEGRAM_CHAT_ID'];

fdk.handle(function(input){
  const bot = new TelegramBot(telegramBotToken);
  const inputData = JSON.parse(input);
  bot.sendMessage(telegramChatId, inputData.message || `Hello`);
  return 'ok';
});
```