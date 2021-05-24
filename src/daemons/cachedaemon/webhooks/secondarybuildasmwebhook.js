const config = require('config');
const WebHookURLData = config.get('PrimaryWebhook.WebhookURL');
const webhook = require("webhook-discord")
const Hook = new webhook.Webhook(WebHookURLData)
Hook.success("6100m\'s asmttpd builder","Built!")
