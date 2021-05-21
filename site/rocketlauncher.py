from subprocess import run, PIPE
from config import DISCORD_WEBHOOK_URL, ROCKETLAUNCHER_USE_DISCORD
from requests import post


def send_message(message):
    if ROCKETLAUNCHER_USE_DISCORD:
        data = {}
        data["content"] = message
        data["username"] = "RocketLauncher"
        data["avatar_url"] = "https://static.thenounproject.com/png/254007-200.png"
        post(DISCORD_WEBHOOK_URL, json=data)


psaq = run(["sudo", "docker", "ps", "-aq"], capture_output=True)
psaq = psaq.stdout
psaq = str(psaq)[2:-1].split("\\n")
command = ["sudo", "docker", "stop"]
for ps in psaq:
    command.append(ps)
run(command)
send_message(":arrow_down: Pulling repository...")
run(["git", "pull"])
send_message(":white_check_mark: Pulled repository!")
send_message(":hammer: Building site...")
run(["sudo", "docker", "build", ".", "-t", "latest"])
send_message(":white_check_mark: Built site!")
send_message(":sunglasses: The site is live!")
run(
    [
        "sudo",
        "docker",
        "run",
        "-dp",
        "443:5000",
        "-v",
        "bruh-moment:/etc/db",
        "latest:latest",
    ]
)
