from subprocess import Popen
from config import DISCORD_WEBHOOK_URL, ROCKETLAUNCHER_USE_DISCORD
from requests import post

def send_message(message):
    if ROCKETLAUNCHER_USE_DISCORD:
        data = {}
        data['content'] = message
        data['username'] = 'RocketLauncher'
        data['avatar_url'] = 'https://static.thenounproject.com/png/254007-200.png'
        requests.post(DISCORD_WEBHOOK_URL, body=data)

send_message(':arrow_down: Pulling repository...')
Popen(['git', 'pull'])
send_message(':white_check_mark: Pulled repository!')
send_message(':hammer: Building site...')
Popen(['sudo','docker','build','.','-t','latest'])
send_message(':white_check_mark: Built site!')
send_message(':sunglasses: Starting site...')
Popen(['sudo','docker','run','-dp','443:5000', '-v', 'bruh-moment:/etc/db', 'latest'])
send_message(':rocket: The site is live!')