# TwitchTimedMessages
This program lets you send Twitch messages in intervals.
## How to use
- You will need to have Go installed. You can get it here: [go.dev](https://go.dev/dl/)
- Execute the following:

  ```
  git clone https://github.com/Sterbehilfe/TwitchTimedMessages.git
  cd TwitchTimedMessages/
  ./build.sh
  ```
- You can now find an executable and a JSON file in the "Build" folder.
- Now insert your twitch username and your OAuth token into it and define messages you would like to send. The interval's unit is milliseconds. Like the following:
  
  ```json
  {
    "Username": "strbhlfe",
    "OAuthToken": "oauth:abcdefghijklmnopqr0123456789",
    "Messages": [
      {
        "Content": "Hello :)",
        "Channel": "strbhlfe",
        "Interval": 60000
      },
      {
        "Content": "This is a test.",
        "Channel": "okayegteatime",
        "Interval": 12000000
      }
    ]
  } 
  ```
- Start the executable file.
