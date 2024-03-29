# SIM808

## Minicom configuration (communication between rpi and sim808)
```
A -    Serial Device      : /dev/ttyAMA0                              |
B - Lockfile Location     : /var/lock                                 |
C -   Callin Program      :                                           |
D -  Callout Program      :                                           |
E -    Bps/Par/Bits       : 9600 8N1                                  |
F - Hardware Flow Control : No                                        |
G - Software Flow Control : Yes                                       |
```

## hardware wiring
### RPI
https://www.raspberrypi.org/documentation/usage/gpio/
```
PIN 15: RX
PIN 14: TX
PIN GROUND: GROUND
```
### SIM808
sim808 evb-v3.2
```
RPI TX to RXD
RPI RX to TXD
RPI GROUND to GROUND
```
Notice: TX and RX ports are crossed.
## Software config
```
raspi-config

- enable serial interface
- disable serial login shell
```
Reboot

In case of issues have a look to [this SO post](https://raspberrypi.stackexchange.com/a/47851/71747)
## tests
```
> minicom

Welcome to minicom 2.7

OPTIONS: I18n
Compiled on Apr 22 2017, 09:14:19.
Port /dev/ttyAMA0, 11:19:28

Press CTRL-A Z for help on special keys

AT
OK
```
### sources
[SIM800_Series_AT_Command_Manual_V1.09.pdf](./SIM800_Series_AT_Command_Manual_V1.09.pdf)

http://acoptex.com/project/264/basics-project-053d-sim808-gsm-gprs-gps-bluetooth-evolution-board-evb-v32-at-acoptexcom/

https://wiki.dfrobot.com/SIM808_GPS_GPRS_GSM_Shield_SKU__TEL0097