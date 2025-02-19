# Name Picker

This was built as a tool to randomly pick a name from a list. If a number is provided, the randomization will run that many times. The final name with the most selections is picked and provided.

Running 100,000,000 randomizations on a Ryzen 7 5800x and 32Gb of ram took a couple of seconds.

## How to 'Install'

You will have to download the file, put it somewhere and run a command through a command-line.

* Download the name picker [Here](https://github.com/MarkDPierce/cult-of-sleep/releases/download/v0.0.1/cult-name-picker.exe) or you can visit the release [page](https://github.com/MarkDPierce/cult-of-sleep/releases/tag/v0.0.1)
* Currently the `cult-name-picker.exe` downloaded into my "Downloads" folder. I will move this somewhere simple to run and access now.
  * I will be using my c:\ drive for this demo. If you have no clue what this means.... Good luck. 
  * Create a folder calling it what ever you want. For me I'll use 'name-picker'
  * Go to your "Downloads" folder and move `cult-name-picker.exe` into your new folder. `c:\name-picker` for me.
  * Go to your `c:\name-picker` folder and create a file called `names.txt`
    * The first line should be a number indicating how many runs you want to execute.
    * Every new line should be a name you wish to randomly pick from.
        Example:
        ```
        1
        Coma Crew
        Slumber Squad 
        Rest in Keys
        ```
At this point, the application should be ready to be ran.

## Picking a Name
You need to know how to open a command prompt or powershell terminal to unlock this achievement. Honestly, its as easy as clicking start, and typing in the word command and clicking the best match.

* Navigate to our folder we created in the last section
  * `cd c:\name-picer`
* Ensure your names.txt exists, has a number, and has a list of name
* Run it. `cult-name-picker.exe`
* Hate the result and run again.