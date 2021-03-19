# WTFullstack-1-Site-TelegramBot-AdminConsole

Video Description:
Click for watching on YouTube (ru lang)
[![Watch the video](https://img.youtube.com/vi/mlLJmS0dBM4/maxresdefault.jpg)](https://www.youtube.com/watch?v=mlLJmS0dBM4&t=503s)

Text Description:<br>
My first try to implement the FullStack Web service.<br>
FrontEnd is based on html/css/js template from open sources:<br>
https://www.free-css.com/free-css-templates/page254/minimalista<br>
I restructured this into a landing page and replaced almost all of the content with a "go-template".<br>
A database based on json packages in the "data" folder.<br>
Inside data/config.json, the starting parameters of the web service are set.<br>
The service is managed through the Telegram Bot.<br><br>
# Admin Menu
<br>
1. Show current data on the site.<br>
There are two command options for this:<br>
   1. 'conf all' - request all<br>
   information.<br>
   2. Request one section<br>
   by number or<br>
   name from the list:<br>
      'conf 0' or 'title'<br>
      'conf 1' or 'head'<br>
      'conf 2' or 'special'<br>
      'conf 3' or 'gallery'<br>
      'conf 4' or 'company'<br>
      'conf 5' or 'choice'<br>
      'conf 6' or 'workers'<br>
      'conf 7' or 'contacts'<br>
<br>
<br>
2. Replace the data on the site.<br>
To do this, use the code command:<br>
   'XXX =' - section element code.<br>
   'Text' - data to be replaced.<br>
   Examples:<br>
      '307 = This is new data'<br>
      '102 = https: //i.site.com/T/1.jpg'<br>
<br>
<br>
3. Reset all site changes.<br>
   'set default conf' - reset command<br>
<br>
<br>
4. Notify about visiting the site.<br>
To control, use the commands:<br>
   'shut up' - disable<br>
   'talk all' - enable<br>
<br>
<br>
5. Write down a question / answer pair. Use the command:<br>
   'nscr' is the command to start<br>
Stages of pair recording:<br>
   1. Enter the identification number of the pair from which you want to continue the dialogue with the user.<br>
If you want this pair to work as the beginning of a new line of dialogue, then enter '0'.<br>
   2. Enter the message you expect from the user.<br>
   3. Enter the message that the bot will reply.<br>
   4. That's it! The couple is created. The bot will create a unique number for her, to which you can later link the development of the dialogue.<br>
<br>
<br>
6. Delete the question / answer pair.<br>
    'dscrX' - replace X with a unique number.<br>
<br>
<br>
7. View all question / answer pairs. Use the command:<br>
   'conf scr' is the command to view.<br>
<br>
<br>
8. Set a new password. Use the command:<br>
   'npasX' - replace X with a new password.<br>
<br>
<br>
9. Request statistics for the day. Use the command:<br>
   'uday' is the command to view the attendance.<br>
<br>
<br>
10. Exit administrator mode:<br>
   'close' is the exit command.<br>
<br>
