# FSD-Metar-GetMessage
Here is a simple code designed for multiple programming languages to fetch vatsim metar interface data. It addresses certain shortcomings in the FSD program. Please modify the weather section in the 'fsd.conf' file to '[weather] source=file'. This adjustment will enable you to access the necessary metar information within Euroscope.

## How To Use

### Python

- The python version requires you to install os, requests, and time modules, otherwise it cannot run
````python
pip install os
````
````python
pip install requests
````
````python
pip install time
````
- After installing these modules, run the following code to execute FSD-Metar-GetMessage.py. It will fetch metar content for you and save it to metar.txt. (By default, it runs every 1800 seconds, and you can modify it according to your personal needs.)
````python
python FSD-Metar-GetMessage.py
````
### PHP

#### Task Scheduler in Windows:

1. Open the "Task Scheduler" application.
2. Click on "Create Basic Task" and follow the prompts to set up a new task.
3. Choose a name and description for the task.
4. Select "Daily" and set the recurrence to occur every 1 day(s).
5. Choose a start date and time for the task.
6. Select "Start a Program" and click "Next."
7. Browse and locate the program/script you want to run (curl.exe or wget.exe) in the "Program/script" field.
8. In the "Add arguments (optional)" field, provide the URL path of your FSD-Metar-GetMessage.php script.
9. Click "Next" and review your settings. Click "Finish" to create the task.
10. Right-click on the newly created task, select "Properties," go to the "Triggers" tab, and modify the trigger to recur every 30 minutes instead of daily.

#### Cron Job in Linux:

1. Open the terminal.
2. Run crontab -e to edit the cron jobs for the current user.
3. Add a new line to the crontab with the following format:
````javascript
*/30 * * * * curl https://your-server-path/FSD-Metar-GetMessage.php
````
Replace https://your-server-path/FSD-Metar-GetMessage.php with the actual URL path to your FSD-Metar-GetMessage.php script.<br>
4. Save and exit the crontab editor.

#### Notice:

Both of these methods will schedule the task to run every 30 minutes, triggering the specified URL path of FSD-Metar-GetMessage.php. Make sure you replace placeholders with the actual paths and URLs specific to your environment.

### Go

1. Make sure you have Go installed on your system.
2. Copy the code to a file named, for example, "main.go".
3. Open a terminal and navigate to the directory containing the "main.go" file.
4. Run the program using the command: **go run main.go**.
5. Once the program is running, you can access the URL **http://localhost:6327/fetch-metar** in your browser or through an HTTP request tool to trigger the fetching and saving of metar data. The saved data will be in the "metar.txt" file.

- For implementing scheduled execution, you can refer to tutorials on how to achieve similar functionality using PHP. This will give you insights into creating scheduled tasks or cron jobs in a different programming language.

<p align="center">
  Looking forward to your ⭐️ Star!
</p>

