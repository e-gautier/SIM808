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
```
RPI TX to RXD
RPI RX to TXD
RPI GROUND to GROUND
```

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
