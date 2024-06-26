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
This is a standard go app. To build this app simply use `go build .`

## Working with the device
### Backup

1. Connect your computer to the JU-06A’s USB port via USB cable.

2. While holding down the BANK [2] button, turn on the power.
  *  It takes about 20 seconds to prepare the drive.

3. Open the “JU-06A” drive on your computer.
  *  The backup files are located in the “BACKUP” folder of the “JU-06A” drive.

4. Copy the backup files into your computer.

5. After copying is completed, eject the USB drive.

#### Windows 8/7
* Right-click on the “JU-06A” icon in “My Computer” and execute “Eject.”
#### Mac OS
* Drag the “JU-06A” icon to the Trash icon in the Dock.

6. Turn the JU-06A power off. 
### Restore

1. As described in the procedure for “Backup” Step 1–3, open the “JU-06A” drive on your computer.

2. Copy the JU-06A backup files into the “RESTORE” folder of the “JU-06A” drive.

3. After copying is completed, eject the USB drive and then press the ARPEGGIO [ON/OFF] button.

4. After the LEDs have completely stopped blinking, turn off the power. 
