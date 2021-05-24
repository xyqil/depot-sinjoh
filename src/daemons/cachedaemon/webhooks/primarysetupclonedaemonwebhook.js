const config = require('config');
const WebHookURLData = config.get('PrimaryWebhook.WebhookURL');
const webhook = require("webhook-discord")
const Hook = new webhook.Webhook(WebHookURLData)
Hook.info("6100m\'s clonedaemon setup tool","Setting up daemon......")
