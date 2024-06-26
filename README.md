```
     ██╗██╗   ██╗       ██████╗  ██████╗  █████╗ 
     ██║██║   ██║      ██╔═████╗██╔════╝ ██╔══██╗
     ██║██║   ██║█████╗██║██╔██║███████╗ ███████║
██   ██║██║   ██║╚════╝████╔╝██║██╔═══██╗██╔══██║
╚█████╔╝╚██████╔╝      ╚██████╔╝╚██████╔╝██║  ██║
 ╚════╝  ╚═════╝        ╚═════╝  ╚═════╝ ╚═╝  ╚═╝                                                                                                             
 ```
 ```
  ██▀███   ▄▄▄       ███▄    █ ▓█████▄  ▒█████   ███▄ ▄███▓
▓██ ▒ ██▒▒████▄     ██ ▀█   █ ▒██▀ ██▌▒██▒  ██▒▓██▒▀█▀ ██▒
▓██ ░▄█ ▒▒██  ▀█▄  ▓██  ▀█ ██▒░██   █▌▒██░  ██▒▓██    ▓██░
▒██▀▀█▄  ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█▄   ▌▒██   ██░▒██    ▒██ 
░██▓ ▒██▒ ▓█   ▓██▒▒██░   ▓██░░▒████▓ ░ ████▓▒░▒██▒   ░██▒
░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒░   ░  ░
  ░▒ ░ ▒░  ▒   ▒▒ ░░ ░░   ░ ▒░ ░ ▒  ▒   ░ ▒ ▒░ ░  ░      ░
  ░░   ░   ░   ▒      ░   ░ ░  ░ ░  ░ ░ ░ ░ ▒  ░      ░   
   ░           ░  ░         ░    ░        ░ ░         ░   
                               ░                         
```

## About

This project creates random patches for the Roland JU-06a sound module.

## Usage
## Working with the tool
Build the app with:
`go build .`

Execuate the app with:
`./ju0ga-generator -patches <n>`
Where n is the number of patches to generate. The maximum is 128.

The application will output the directory the patches are created in, and the name of each patch as it is created.
```
Created directory farad...
Creating... darkle          
Creating... fowl            
Creating... Salk clout    
```

Once this is complete these patches can be transferred to the device.

## Working with the device
### Backup

* Connect your computer to the JU-06A’s USB port via USB cable.
* While holding down the BANK [2] button, turn on the power.
  *  It takes about 20 seconds to prepare the drive.
* Open the “JU-06A” drive on your computer.
  *  The backup files are located in the “BACKUP” folder of the “JU-06A” drive.
* Copy the backup files into your computer.
* After copying is completed, eject the USB drive.

###### Windows 8/7
* Right-click on the “JU-06A” icon in “My Computer” and execute “Eject.”
###### Mac OS
* Drag the “JU-06A” icon to the Trash icon in the Dock.

* Turn the JU-06A power off. 

### Restore
* As described in the procedure for “Backup” Step 1–3, open the “JU-06A” drive on your computer.
* Copy the JU-06A backup files into the “RESTORE” folder of the “JU-06A” drive.
* After copying is completed, eject the USB drive and then press the ARPEGGIO [ON/OFF] button.
* After the LEDs have completely stopped blinking, turn off the power. 
