const config = require('config');
const WebHookURLData = config.get('PrimaryBuildASMWebhook.WebhookURL');
const webhook = require("webhook-discord")
const Hook = new webhook.Webhook(WebHookURLData)
Hook.info("6100m\'s asmttpd builder","Building......")
